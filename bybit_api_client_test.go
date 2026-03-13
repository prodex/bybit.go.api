package bybit_connector

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type baseTestSuite struct {
	suite.Suite
	client    *mockedClient
	apiKey    string
	apiSecret string
	baseURL   string
}

func (s *baseTestSuite) r() *require.Assertions {
	return s.Require()
}

func (s *baseTestSuite) SetupTest() {
	s.apiKey = "dummyAPIKey"
	s.apiSecret = "dummyApiSecret"
	s.baseURL = "https://dummyapi.com"
	s.client = newMockedClient(s.apiKey, s.apiSecret, s.baseURL)
}

func (s *baseTestSuite) mockDo(data []byte, err error, statusCode ...int) {
	s.client.Client.do = s.client.do
	code := http.StatusOK
	if len(statusCode) > 0 {
		code = statusCode[0]
	}
	s.client.On("do", anyHTTPRequest()).Return(newHTTPResponse(data, code), err)
}

func (s *baseTestSuite) assertDo() {
	s.client.AssertCalled(s.T(), "do", anyHTTPRequest())
}

func (s *baseTestSuite) assertReq(f func(r *request)) {
	s.client.assertReq = f
}

func (s *baseTestSuite) assertRequestEqual(e, a *request) {
	s.assertURLValuesEqual(e.query, a.query)
}

func (s *baseTestSuite) assertURLValuesEqual(e, a url.Values) {
	var eKeys, aKeys []string
	for k := range e {
		eKeys = append(eKeys, k)
	}
	for k := range a {
		aKeys = append(aKeys, k)
	}
	r := s.r()
	r.Len(aKeys, len(eKeys))
	for k := range a {
		switch k {
		case timestampKey, signatureKey:
			r.NotEmpty(a.Get(k))
			continue
		}
		r.Equal(e[k], a[k], k)
	}
}

func anythingOfType(t string) mock.AnythingOfTypeArgument {
	return mock.AnythingOfType(t)
}

func newContext() context.Context {
	return context.Background()
}

func anyHTTPRequest() mock.AnythingOfTypeArgument {
	return anythingOfType("*http.Request")
}

func newHTTPResponse(data []byte, statusCode int) *http.Response {
	return &http.Response{
		Body:       io.NopCloser(bytes.NewBuffer(data)),
		StatusCode: statusCode,
	}
}

func newRequest() *request {
	r := &request{
		query: url.Values{},
	}
	return r
}

func newSignedRequest() *request {
	r, _ := newRequest().setParams(params{
		timestampKey:  "",
		signatureKey:  "",
		apiRequestKey: "",
		recvWindowKey: "5000",
		signTypeKey:   "2",
	})
	return r
}

type assertReqFunc func(r *request)

type mockedClient struct {
	mock.Mock
	*Client
	assertReq assertReqFunc
}

func newMockedClient(apiKey, apiSecret, baseURL string) *mockedClient {
	m := new(mockedClient)
	m.Client = NewBybitHttpClient(apiKey, apiSecret, WithBaseURL(baseURL))
	return m
}

func (m *mockedClient) do(req *http.Request) (*http.Response, error) {
	if m.assertReq != nil {
		r := newRequest()
		r.query = req.URL.Query()
		if req.Body != nil && req.ContentLength > 0 {
			bs := make([]byte, req.ContentLength)
			_, err := req.Body.Read(bs)
			if err != nil && err != io.EOF {
				return nil, err // Handle read error
			}
			_ = req.Body.Close() // Close the body if we have read from it
			r.body = bytes.NewBuffer(bs)
		}
		m.assertReq(r)
	}
	args := m.Called(req)
	return args.Get(0).(*http.Response), args.Error(1)
}

func TestFormatTimestamp(t *testing.T) {
	tm, _ := time.Parse("2006-01-02 15:04:05", "2018-06-01 01:01:01")
	assert.Equal(t, int64(1527814861000), FormatTimestamp(tm))
}

// trackableBody tracks whether Close was called
type trackableBody struct {
	io.Reader
	closed bool
}

func (tb *trackableBody) Close() error {
	tb.closed = true
	return nil
}

func TestCallAPI_BodyClosed(t *testing.T) {
	body := &trackableBody{Reader: bytes.NewBufferString(`{"retCode":0,"retMsg":"OK"}`)}
	c := NewBybitHttpClient("key", "secret", WithBaseURL("https://test.com"))
	c.do = func(req *http.Request) (*http.Response, error) {
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       body,
		}, nil
	}

	r := &request{method: http.MethodGet, endpoint: "/test"}
	_, err := c.callAPI(context.Background(), r)
	assert.NoError(t, err)
	assert.True(t, body.closed, "response body should be closed after callAPI")
}

func TestCallAPI_HTTPError(t *testing.T) {
	data := `{"retCode":10001,"retMsg":"invalid request"}`
	c := NewBybitHttpClient("key", "secret", WithBaseURL("https://test.com"))
	c.do = func(req *http.Request) (*http.Response, error) {
		return &http.Response{
			StatusCode: http.StatusBadRequest,
			Body:       io.NopCloser(bytes.NewBufferString(data)),
		}, nil
	}

	r := &request{method: http.MethodGet, endpoint: "/test"}
	result, err := c.callAPI(context.Background(), r)
	assert.Error(t, err)
	assert.Nil(t, result)
}

func TestCallAPI_SignedPOST(t *testing.T) {
	c := NewBybitHttpClient("testAPIKey", "testAPISecret", WithBaseURL("https://test.com"))
	var capturedReq *http.Request
	c.do = func(req *http.Request) (*http.Response, error) {
		capturedReq = req
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(bytes.NewBufferString(`{"retCode":0,"retMsg":"OK"}`)),
		}, nil
	}

	r := &request{
		method:   http.MethodPost,
		endpoint: "/v5/order/create",
		secType:  secTypeSigned,
	}
	r.setParams(params{"symbol": "BTCUSDT", "side": "Buy"})
	_, err := c.callAPI(context.Background(), r)

	assert.NoError(t, err)
	assert.NotNil(t, capturedReq)
	assert.NotEmpty(t, capturedReq.Header.Get(signatureKey))
	assert.NotEmpty(t, capturedReq.Header.Get(timestampKey))
	assert.Equal(t, "testAPIKey", capturedReq.Header.Get(apiRequestKey))
	assert.Equal(t, "application/json", capturedReq.Header.Get("Content-Type"))
}

func TestNewBybitHttpClient_WithProxy(t *testing.T) {
	// Save original transport
	originalTransport := http.DefaultClient.Transport

	_ = NewBybitHttpClient("key", "secret", WithProxyURL("http://proxy.example.com:8080"))

	// http.DefaultClient.Transport should NOT have been mutated
	assert.Equal(t, originalTransport, http.DefaultClient.Transport,
		"http.DefaultClient.Transport should not be mutated by proxy configuration")
}

func TestNewBybitHttpClient_DefaultTimeout(t *testing.T) {
	c := NewBybitHttpClient("key", "secret")
	assert.Equal(t, 30*time.Second, c.HTTPClient.Timeout,
		"default HTTP client should have 30s timeout")
}
