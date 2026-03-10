package bybit_connector

import (
	"net/http"
	"testing"

	"github.com/prodex/bybit.go.api/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type marketTestSuite struct {
	baseTestSuite
}

func TestMarketService(t *testing.T) {
	suite.Run(t, new(marketTestSuite))
}

// --- Request tests ---

func (s *marketTestSuite) TestGetMarketKline() {
	data := []byte(`{
    "retCode": 0,
    "retMsg": "OK",
    "result": {
        "symbol": "BTCUSD",
        "category": "inverse",
        "list": [
            ["1670608800000","17071","17073","17027","17055.5","268611","15.74462667"],
            ["1670605200000","17071.5","17071.5","17061","17071","4177","0.24469757"],
            ["1670601600000","17086.5","17088","16978","17071.5","6356","0.37288112"]
        ]
    },
    "retExtInfo": {},
    "time": 1672025956592
}`)
	s.mockDo(data, nil)
	defer s.assertDo()

	s.assertReq(func(r *request) {
		e := newRequest()
		e.method = http.MethodGet
		e.setParams(params{
			"category": "inverse",
			"symbol":   "BTCUSD",
			"interval": "1",
			"start":    uint64(1499040000000),
			"end":      uint64(1499040000001),
			"limit":    10,
		})
		s.assertRequestEqual(e, r)
	})

	svc := s.client.NewUtaBybitServiceWithParams(map[string]interface{}{
		"category": "inverse",
		"symbol":   "BTCUSD",
		"interval": "1",
		"start":    uint64(1499040000000),
		"end":      uint64(1499040000001),
		"limit":    10,
	})
	_, err := svc.GetMarketKline(newContext())
	s.r().NoError(err)
}

func (s *marketTestSuite) TestGetMarkPriceKline() {
	data := []byte(`{
    "retCode": 0,
    "retMsg": "OK",
    "result": {
        "symbol": "BTCUSDT",
        "category": "linear",
        "list": [["1670608800000","17164.16","17164.16","17121.5","17131.64"]]
    },
    "retExtInfo": {},
    "time": 1672026361839
}`)
	s.mockDo(data, nil)
	defer s.assertDo()

	s.assertReq(func(r *request) {
		e := newRequest()
		e.method = http.MethodGet
		e.setParams(params{
			"category": "inverse",
			"symbol":   "BTCUSDT",
			"interval": "1",
			"start":    uint64(1499040000000),
			"end":      uint64(1499040000001),
			"limit":    10,
		})
		s.assertRequestEqual(e, r)
	})

	svc := s.client.NewUtaBybitServiceWithParams(map[string]interface{}{
		"category": "inverse",
		"symbol":   "BTCUSDT",
		"interval": "1",
		"start":    uint64(1499040000000),
		"end":      uint64(1499040000001),
		"limit":    10,
	})
	_, err := svc.GetMarkPriceKline(newContext())
	s.r().NoError(err)
}

func (s *marketTestSuite) TestGetIndexPriceKline() {
	data := []byte(`{
    "retCode": 0,
    "retMsg": "OK",
    "result": {
        "symbol": "BTCUSDZ22",
        "category": "inverse",
        "list": [
            ["1670608800000","17167.00","17167.00","17161.90","17163.07"],
            ["1670608740000","17166.54","17167.69","17165.42","17167.00"]
        ]
    },
    "retExtInfo": {},
    "time": 1672026471128
}`)
	s.mockDo(data, nil)
	defer s.assertDo()

	s.assertReq(func(r *request) {
		e := newRequest()
		e.method = http.MethodGet
		e.setParams(params{
			"category": "inverse",
			"symbol":   "BTCUSDZ22",
			"interval": "1",
			"start":    uint64(1499040000000),
			"end":      uint64(1499040000001),
			"limit":    10,
		})
		s.assertRequestEqual(e, r)
	})

	svc := s.client.NewUtaBybitServiceWithParams(map[string]interface{}{
		"category": "inverse",
		"symbol":   "BTCUSDZ22",
		"interval": "1",
		"start":    uint64(1499040000000),
		"end":      uint64(1499040000001),
		"limit":    10,
	})
	_, err := svc.GetIndexPriceKline(newContext())
	s.r().NoError(err)
}

func (s *marketTestSuite) TestGetPremiumIndexPriceKline() {
	data := []byte(`{
    "retCode": 0,
    "retMsg": "OK",
    "result": {
        "symbol": "BTCPERP",
        "category": "linear",
        "list": [
            ["1672026540000","0.000000","0.000000","0.000000","0.000000"],
            ["1672026480000","0.000000","0.000000","0.000000","0.000000"]
        ]
    },
    "retExtInfo": {},
    "time": 1672026605042
}`)
	s.mockDo(data, nil)
	defer s.assertDo()

	s.assertReq(func(r *request) {
		e := newRequest()
		e.method = http.MethodGet
		e.setParams(params{
			"category": "inverse",
			"symbol":   "BTCPERP",
			"interval": "1",
			"start":    uint64(1499040000000),
			"end":      uint64(1499040000001),
			"limit":    10,
		})
		s.assertRequestEqual(e, r)
	})

	svc := s.client.NewUtaBybitServiceWithParams(map[string]interface{}{
		"category": "inverse",
		"symbol":   "BTCPERP",
		"interval": "1",
		"start":    uint64(1499040000000),
		"end":      uint64(1499040000001),
		"limit":    10,
	})
	_, err := svc.GetPremiumIndexPriceKline(newContext())
	s.r().NoError(err)
}

func (s *marketTestSuite) TestGetInstrumentInfo() {
	data := []byte(`{
    "retCode": 0,
    "retMsg": "OK",
    "result": {
        "category": "spot",
        "list": [
            {
                "symbol": "BTCUSDT",
                "baseCoin": "BTC",
                "quoteCoin": "USDT",
                "innovation": "0",
                "status": "Trading",
                "marginTrading": "both",
                "lotSizeFilter": {
                    "basePrecision": "0.000001",
                    "quotePrecision": "0.00000001",
                    "minOrderQty": "0.000048",
                    "maxOrderQty": "71.73956243",
                    "minOrderAmt": "1",
                    "maxOrderAmt": "2000000"
                },
                "priceFilter": {
                    "tickSize": "0.01"
                }
            }
        ]
    },
    "retExtInfo": {},
    "time": 1672712468011
}`)
	s.mockDo(data, nil)
	defer s.assertDo()

	s.assertReq(func(r *request) {
		e := newRequest()
		e.method = http.MethodGet
		e.setParams(params{
			"category": "inverse",
			"symbol":   "BTCUSD",
			"status":   models.SymbolStatusTrading,
			"baseCoin": "BTC",
			"limit":    10,
			"cursor":   "cursor",
		})
		s.assertRequestEqual(e, r)
	})

	svc := s.client.NewUtaBybitServiceWithParams(map[string]interface{}{
		"category": "inverse",
		"symbol":   "BTCUSD",
		"status":   models.SymbolStatusTrading,
		"baseCoin": "BTC",
		"limit":    10,
		"cursor":   "cursor",
	})
	_, err := svc.GetInstrumentInfo(newContext())
	s.r().NoError(err)
}

func (s *marketTestSuite) TestGetOrderBookInfo() {
	data := []byte(`{
    "retCode": 0,
    "retMsg": "SUCCESS",
    "result": {
        "s": "BTC-30DEC22-18000-C",
        "b": [["5","3.12"]],
        "a": [["175","4.88"]],
        "u": 1203433656,
        "ts": 1672043188375
    },
    "retExtInfo": {},
    "time": 1672043199230
}`)
	s.mockDo(data, nil)
	defer s.assertDo()

	s.assertReq(func(r *request) {
		e := newRequest()
		e.method = http.MethodGet
		e.setParams(params{
			"category": "inverse",
			"symbol":   "BTCUSD",
			"limit":    10,
		})
		s.assertRequestEqual(e, r)
	})

	svc := s.client.NewUtaBybitServiceWithParams(map[string]interface{}{
		"category": "inverse",
		"symbol":   "BTCUSD",
		"limit":    10,
	})
	_, err := svc.GetOrderBookInfo(newContext())
	s.r().NoError(err)
}

func (s *marketTestSuite) TestGetMarketTickers() {
	data := []byte(`{
    "retCode": 0,
    "retMsg": "OK",
    "result": {
        "category": "inverse",
        "list": [
            {
                "symbol": "BTCUSD",
                "lastPrice": "16597.00",
                "indexPrice": "16598.54",
                "markPrice": "16596.00",
                "prevPrice24h": "16464.50",
                "price24hPcnt": "0.008047",
                "highPrice24h": "30912.50",
                "lowPrice24h": "15700.00",
                "prevPrice1h": "16595.50",
                "openInterest": "373504107",
                "openInterestValue": "22505.67",
                "turnover24h": "2352.94950046",
                "volume24h": "49337318",
                "fundingRate": "-0.001034",
                "nextFundingTime": "1672387200000",
                "predictedDeliveryPrice": "",
                "basisRate": "",
                "deliveryFeeRate": "",
                "deliveryTime": "0",
                "ask1Size": "1",
                "bid1Price": "16596.00",
                "ask1Price": "16597.50",
                "bid1Size": "1",
                "basis": ""
            }
        ]
    },
    "retExtInfo": {},
    "time": 1672376496682
}`)
	s.mockDo(data, nil)
	defer s.assertDo()

	s.assertReq(func(r *request) {
		e := newRequest()
		e.method = http.MethodGet
		e.setParams(params{
			"category": "inverse",
			"symbol":   "BTCUSD",
			"baseCoin": "BTC",
			"expDate":  "2022-12-30",
		})
		s.assertRequestEqual(e, r)
	})

	svc := s.client.NewUtaBybitServiceWithParams(map[string]interface{}{
		"category": "inverse",
		"symbol":   "BTCUSD",
		"baseCoin": "BTC",
		"expDate":  "2022-12-30",
	})
	_, err := svc.GetMarketTickers(newContext())
	s.r().NoError(err)
}

func (s *marketTestSuite) TestGetFundingRateHistory() {
	data := []byte(`{
    "retCode": 0,
    "retMsg": "OK",
    "result": {
        "category": "linear",
        "list": [
            {
                "symbol": "ETHPERP",
                "fundingRate": "0.0001",
                "fundingRateTimestamp": "1672041600000"
            }
        ]
    },
    "retExtInfo": {},
    "time": 1672051897447
}`)
	s.mockDo(data, nil)
	defer s.assertDo()

	s.assertReq(func(r *request) {
		e := newRequest()
		e.method = http.MethodGet
		e.setParams(params{
			"category":  "inverse",
			"symbol":    "BTCUSD",
			"startTime": uint64(1499040000000),
			"endTime":   uint64(1499040000001),
			"limit":     10,
		})
		s.assertRequestEqual(e, r)
	})

	svc := s.client.NewUtaBybitServiceWithParams(map[string]interface{}{
		"category":  "inverse",
		"symbol":    "BTCUSD",
		"startTime": uint64(1499040000000),
		"endTime":   uint64(1499040000001),
		"limit":     10,
	})
	_, err := svc.GetFundingRateHistory(newContext())
	s.r().NoError(err)
}

func (s *marketTestSuite) TestGetPublicRecentTrades() {
	data := []byte(`{
    "retCode": 0,
    "retMsg": "OK",
    "result": {
        "category": "spot",
        "list": [
            {
                "execId": "2100000000007764263",
                "symbol": "BTCUSDT",
                "price": "16618.49",
                "size": "0.00012",
                "side": "Buy",
                "time": "1672052955758",
                "isBlockTrade": false
            }
        ]
    },
    "retExtInfo": {},
    "time": 1672053054358
}`)
	s.mockDo(data, nil)
	defer s.assertDo()

	s.assertReq(func(r *request) {
		e := newRequest()
		e.method = http.MethodGet
		e.setParams(params{
			"category":   "inverse",
			"symbol":     "BTCUSD",
			"baseCoin":   "BTC",
			"optionType": "optionType",
			"limit":      10,
		})
		s.assertRequestEqual(e, r)
	})

	svc := s.client.NewUtaBybitServiceWithParams(map[string]interface{}{
		"category":   "inverse",
		"symbol":     "BTCUSD",
		"baseCoin":   "BTC",
		"optionType": "optionType",
		"limit":      10,
	})
	_, err := svc.GetPublicRecentTrades(newContext())
	s.r().NoError(err)
}

func (s *marketTestSuite) TestGetOpenInterests() {
	data := []byte(`{
    "retCode": 0,
    "retMsg": "OK",
    "result": {
        "symbol": "BTCUSD",
        "category": "inverse",
        "list": [
            {"openInterest": "461134384.00000000", "timestamp": "1669571400000"},
            {"openInterest": "461134292.00000000", "timestamp": "1669571100000"}
        ],
        "nextPageCursor": ""
    },
    "retExtInfo": {},
    "time": 1672053548579
}`)
	s.mockDo(data, nil)
	defer s.assertDo()

	s.assertReq(func(r *request) {
		e := newRequest()
		e.method = http.MethodGet
		e.setParams(params{
			"category":     "inverse",
			"symbol":       "BTCUSD",
			"intervalTime": "intervalTime",
			"startTime":    uint64(1499040000000),
			"endTime":      uint64(1499040000001),
			"limit":        10,
			"cursor":       "cursor",
		})
		s.assertRequestEqual(e, r)
	})

	svc := s.client.NewUtaBybitServiceWithParams(map[string]interface{}{
		"category":     "inverse",
		"symbol":       "BTCUSD",
		"intervalTime": "intervalTime",
		"startTime":    uint64(1499040000000),
		"endTime":      uint64(1499040000001),
		"limit":        10,
		"cursor":       "cursor",
	})
	_, err := svc.GetOpenInterests(newContext())
	s.r().NoError(err)
}

func (s *marketTestSuite) TestGetHistoryVolatility() {
	data := []byte(`{
    "retCode": 0,
    "retMsg": "SUCCESS",
    "category": "option",
    "result": [
        {"period": 7, "value": "0.27545620", "time": "1672232400000"}
    ]
}`)
	s.mockDo(data, nil)
	defer s.assertDo()

	s.assertReq(func(r *request) {
		e := newRequest()
		e.method = http.MethodGet
		e.setParams(params{
			"category":  "inverse",
			"baseCoin":  "BTC",
			"period":    "period",
			"startTime": uint64(1499040000000),
			"endTime":   uint64(1499040000001),
		})
		s.assertRequestEqual(e, r)
	})

	svc := s.client.NewUtaBybitServiceWithParams(map[string]interface{}{
		"category":  "inverse",
		"baseCoin":  "BTC",
		"period":    "period",
		"startTime": uint64(1499040000000),
		"endTime":   uint64(1499040000001),
	})
	_, err := svc.GetHistoryVolatility(newContext())
	s.r().NoError(err)
}

func (s *marketTestSuite) TestGetMarketInsurance() {
	data := []byte(`{
    "retCode": 0,
    "retMsg": "OK",
    "result": {
        "updatedTime": "1672012800000",
        "list": [{"coin": "ETH", "balance": "0.00187332", "value": "0"}]
    },
    "retExtInfo": {},
    "time": 1672053931991
}`)
	s.mockDo(data, nil)
	defer s.assertDo()

	s.assertReq(func(r *request) {
		e := newRequest()
		e.method = http.MethodGet
		e.setParams(params{
			"coin": "BTC",
		})
		s.assertRequestEqual(e, r)
	})

	svc := s.client.NewUtaBybitServiceWithParams(map[string]interface{}{
		"coin": "BTC",
	})
	_, err := svc.GetMarketInsurance(newContext())
	s.r().NoError(err)
}

func (s *marketTestSuite) TestGetMarketRiskLimits() {
	data := []byte(`{
    "retCode": 0,
    "retMsg": "OK",
    "result": {
        "category": "inverse",
        "list": [
            {
                "id": 1,
                "symbol": "BTCUSD",
                "riskLimitValue": "150",
                "maintenanceMargin": "0.5",
                "initialMargin": "1",
                "isLowestRisk": 1,
                "maxLeverage": "100.00"
            }
        ]
    },
    "retExtInfo": {},
    "time": 1672054488010
}`)
	s.mockDo(data, nil)
	defer s.assertDo()

	s.assertReq(func(r *request) {
		e := newRequest()
		e.method = http.MethodGet
		e.setParams(params{
			"category": "inverse",
			"symbol":   "BTCUSDT",
		})
		s.assertRequestEqual(e, r)
	})

	svc := s.client.NewUtaBybitServiceWithParams(map[string]interface{}{
		"category": "inverse",
		"symbol":   "BTCUSDT",
	})
	_, err := svc.GetMarketRiskLimits(newContext())
	s.r().NoError(err)
}

func (s *marketTestSuite) TestGetDeliveryPrice() {
	data := []byte(`{
    "retCode": 0,
    "retMsg": "success",
    "result": {
        "category": "option",
        "nextPageCursor": "",
        "list": [
            {
                "symbol": "ETH-26DEC22-1400-C",
                "deliveryPrice": "1220.728594450",
                "deliveryTime": "1672041600000"
            }
        ]
    },
    "retExtInfo": {},
    "time": 1672055336993
}`)
	s.mockDo(data, nil)
	defer s.assertDo()

	s.assertReq(func(r *request) {
		e := newRequest()
		e.method = http.MethodGet
		e.setParams(params{
			"category": "inverse",
			"symbol":   "BTCUSDT",
			"baseCoin": "BTC",
			"limit":    10,
			"cursor":   "cursor",
		})
		s.assertRequestEqual(e, r)
	})

	svc := s.client.NewUtaBybitServiceWithParams(map[string]interface{}{
		"category": "inverse",
		"symbol":   "BTCUSDT",
		"baseCoin": "BTC",
		"limit":    10,
		"cursor":   "cursor",
	})
	_, err := svc.GetDeliveryPrice(newContext())
	s.r().NoError(err)
}

func (s *marketTestSuite) TestGetLongShortRatio() {
	data := []byte(`{
    "retCode": 0,
    "retMsg": "OK",
    "result": {
        "list": [
            {
                "symbol": "BTCUSDT",
                "buyRatio": "0.5777",
                "sellRatio": "0.4223",
                "timestamp": "1695772800000"
            }
        ]
    },
    "retExtInfo": {},
    "time": 1695785131028
}`)
	s.mockDo(data, nil)
	defer s.assertDo()

	s.assertReq(func(r *request) {
		e := newRequest()
		e.method = http.MethodGet
		e.setParams(params{
			"category": "inverse",
			"baseCoin": "BTC",
			"period":   "period",
			"limit":    10,
		})
		s.assertRequestEqual(e, r)
	})

	svc := s.client.NewUtaBybitServiceWithParams(map[string]interface{}{
		"category": "inverse",
		"baseCoin": "BTC",
		"period":   "period",
		"limit":    10,
	})
	_, err := svc.GetLongShortRatio(newContext())
	s.r().NoError(err)
}

func (s *marketTestSuite) TestGetServerTime() {
	data := []byte(`{
    "retCode": 0,
    "retMsg": "OK",
    "result": {
        "timeSecond": "1688639403",
        "timeNano": "1688639403423213947"
    },
    "retExtInfo": {},
    "time": 1688639403423
}`)
	s.mockDo(data, nil)
	defer s.assertDo()

	s.assertReq(func(r *request) {
		e := newRequest()
		e.method = http.MethodGet
		s.assertRequestEqual(e, r)
	})

	svc := s.client.NewUtaBybitServiceNoParams()
	_, err := svc.GetServerTime(newContext())
	s.r().NoError(err)
}

// --- Response parsing tests ---

func TestParseMarketKlineResponse(t *testing.T) {
	data := []byte(`{
    "retCode": 0,
    "retMsg": "OK",
    "result": {
        "symbol": "BTCUSD",
        "category": "inverse",
        "list": [
            ["1670608800000","17071","17073","17027","17055.5","268611","15.74462667"],
            ["1670605200000","17071.5","17071.5","17061","17071","4177","0.24469757"],
            ["1670601600000","17086.5","17088","16978","17071.5","6356","0.37288112"]
        ]
    },
    "retExtInfo": {},
    "time": 1672025956592
}`)
	res, _, err := GetMarketKlineResponse(nil, data, nil)
	assert.NoError(t, err)
	assert.Equal(t, "BTCUSD", res.Symbol)
	assert.Equal(t, models.Category("inverse"), res.Category)
	assert.Len(t, res.List, 3)
	assert.Equal(t, "1670608800000", res.List[0].StartTime)
	assert.Equal(t, "17071", res.List[0].OpenPrice)
	assert.Equal(t, "17073", res.List[0].HighPrice)
	assert.Equal(t, "17027", res.List[0].LowPrice)
	assert.Equal(t, "17055.5", res.List[0].ClosePrice)
	assert.Equal(t, "268611", res.List[0].Volume)
	assert.Equal(t, "15.74462667", res.List[0].Turnover)
}

func TestParseMarkPriceKlineResponse(t *testing.T) {
	data := []byte(`{
    "retCode": 0,
    "retMsg": "OK",
    "result": {
        "symbol": "BTCUSDT",
        "category": "linear",
        "list": [["1670608800000","17164.16","17164.16","17121.5","17131.64"]]
    },
    "retExtInfo": {},
    "time": 1672026361839
}`)
	res, err := GetMarkPriceKline(nil, data, nil)
	assert.NoError(t, err)
	assert.Equal(t, "BTCUSDT", res.Symbol)
	assert.Equal(t, models.Category("linear"), res.Category)
	assert.Len(t, res.List, 1)
	assert.Equal(t, "17164.16", res.List[0].OpenPrice)
}

func TestParseIndexPriceKlineResponse(t *testing.T) {
	data := []byte(`{
    "retCode": 0,
    "retMsg": "OK",
    "result": {
        "symbol": "BTCUSDZ22",
        "category": "inverse",
        "list": [
            ["1670608800000","17167.00","17167.00","17161.90","17163.07"],
            ["1670608740000","17166.54","17167.69","17165.42","17167.00"]
        ]
    },
    "retExtInfo": {},
    "time": 1672026471128
}`)
	res, err := GetIndexPriceKline(nil, data, nil)
	assert.NoError(t, err)
	assert.Equal(t, "BTCUSDZ22", res.Symbol)
	assert.Len(t, res.List, 2)
}

func TestParsePremiumIndexKlineResponse(t *testing.T) {
	data := []byte(`{
    "retCode": 0,
    "retMsg": "OK",
    "result": {
        "symbol": "BTCPERP",
        "category": "linear",
        "list": [
            ["1672026540000","0.000000","0.000000","0.000000","0.000000"],
            ["1672026480000","0.000000","0.000000","0.000000","0.000000"]
        ]
    },
    "retExtInfo": {},
    "time": 1672026605042
}`)
	res, err := GetPremiumIndexKline(nil, data, nil)
	assert.NoError(t, err)
	assert.Equal(t, "BTCPERP", res.Symbol)
	assert.Len(t, res.List, 2)
}

func TestParseServerTimeResponse(t *testing.T) {
	data := []byte(`{
    "retCode": 0,
    "retMsg": "OK",
    "result": {
        "timeSecond": "1688639403",
        "timeNano": "1688639403423213947"
    },
    "retExtInfo": {},
    "time": 1688639403423
}`)
	var res models.GetServerTimeResponse
	err := json.Unmarshal(data, &res)
	assert.NoError(t, err)
	assert.Equal(t, 0, res.RetCode)
	assert.Equal(t, "1688639403", res.Result.TimeSecond)
	assert.Equal(t, "1688639403423213947", res.Result.TimeNano)
}

func TestParseOrderBookResponse(t *testing.T) {
	data := []byte(`{
    "retCode": 0,
    "retMsg": "SUCCESS",
    "result": {
        "s": "BTC-30DEC22-18000-C",
        "b": [["5","3.12"]],
        "a": [["175","4.88"]],
        "u": 1203433656,
        "ts": 1672043188375
    },
    "retExtInfo": {},
    "time": 1672043199230
}`)
	var res models.MarketOrderBookResponse
	err := json.Unmarshal(data, &res)
	assert.NoError(t, err)
	assert.Equal(t, "BTC-30DEC22-18000-C", res.Result.Symbol)
	assert.Equal(t, int64(1672043188375), res.Result.Timestamp)
	assert.Equal(t, int64(1203433656), res.Result.UpdateID)
	assert.Len(t, res.Result.Bids, 1)
	assert.Equal(t, models.OrderBookEntry{"5", "3.12"}, res.Result.Bids[0])
	assert.Len(t, res.Result.Asks, 1)
	assert.Equal(t, models.OrderBookEntry{"175", "4.88"}, res.Result.Asks[0])
}

func TestParseTickersResponse(t *testing.T) {
	data := []byte(`{
    "retCode": 0,
    "retMsg": "OK",
    "result": {
        "category": "inverse",
        "list": [
            {
                "symbol": "BTCUSD",
                "lastPrice": "16597.00",
                "indexPrice": "16598.54",
                "markPrice": "16596.00",
                "prevPrice24h": "16464.50",
                "price24hPcnt": "0.008047",
                "highPrice24h": "30912.50",
                "lowPrice24h": "15700.00",
                "prevPrice1h": "16595.50",
                "openInterest": "373504107",
                "openInterestValue": "22505.67",
                "turnover24h": "2352.94950046",
                "volume24h": "49337318",
                "fundingRate": "-0.001034",
                "nextFundingTime": "1672387200000",
                "predictedDeliveryPrice": "",
                "basisRate": "",
                "deliveryFeeRate": "",
                "deliveryTime": "0",
                "ask1Size": "1",
                "bid1Price": "16596.00",
                "ask1Price": "16597.50",
                "bid1Size": "1",
                "basis": ""
            }
        ]
    },
    "retExtInfo": {},
    "time": 1672376496682
}`)
	var res models.MarketTickersResponse
	err := json.Unmarshal(data, &res)
	assert.NoError(t, err)
	assert.Equal(t, "inverse", res.Result.Category)
	assert.Len(t, res.Result.List, 1)
	assert.Equal(t, "BTCUSD", res.Result.List[0].Symbol)
	assert.Equal(t, "16597.00", res.Result.List[0].LastPrice)
	assert.Equal(t, "-0.001034", res.Result.List[0].FundingRate)
}

func TestParseFundingRatesResponse(t *testing.T) {
	data := []byte(`{
    "retCode": 0,
    "retMsg": "OK",
    "result": {
        "category": "linear",
        "list": [
            {
                "symbol": "ETHPERP",
                "fundingRate": "0.0001",
                "fundingRateTimestamp": "1672041600000"
            }
        ]
    },
    "retExtInfo": {},
    "time": 1672051897447
}`)
	var res models.MarketFundingRatesResponse
	err := json.Unmarshal(data, &res)
	assert.NoError(t, err)
	assert.Equal(t, "linear", res.Result.Category)
	assert.Len(t, res.Result.List, 1)
	assert.Equal(t, "ETHPERP", res.Result.List[0].Symbol)
	assert.Equal(t, "0.0001", res.Result.List[0].FundingRate)
}

func TestParseInstrumentInfoResponse(t *testing.T) {
	data := []byte(`{
    "retCode": 0,
    "retMsg": "OK",
    "result": {
        "category": "spot",
        "list": [
            {
                "symbol": "BTCUSDT",
                "baseCoin": "BTC",
                "quoteCoin": "USDT",
                "innovation": "0",
                "status": "Trading",
                "marginTrading": "both",
                "lotSizeFilter": {
                    "basePrecision": "0.000001",
                    "quotePrecision": "0.00000001",
                    "minOrderQty": "0.000048",
                    "maxOrderQty": "71.73956243",
                    "minOrderAmt": "1",
                    "maxOrderAmt": "2000000"
                },
                "priceFilter": {
                    "tickSize": "0.01"
                }
            }
        ],
        "nextPageCursor": ""
    },
    "retExtInfo": {},
    "time": 1672712468011
}`)
	var wrapper struct {
		Result models.InstrumentInfoResponse `json:"result"`
	}
	err := json.Unmarshal(data, &wrapper)
	assert.NoError(t, err)
	res := wrapper.Result
	assert.Equal(t, models.Category("spot"), res.Category)
	assert.Len(t, res.List, 1)
	assert.Equal(t, "BTCUSDT", res.List[0].Symbol)
	assert.Equal(t, "BTC", res.List[0].BaseCoin)
	assert.Equal(t, "USDT", res.List[0].QuoteCoin)
	assert.Equal(t, models.SymbolStatusTrading, res.List[0].Status)
	assert.Equal(t, "both", res.List[0].MarginTrading)
	assert.Equal(t, "0.000001", res.List[0].LotSizeFilter.BasePrecision)
	assert.Equal(t, "0.01", res.List[0].PriceFilter.TickSize)
}

func TestParseRecentTradesResponse(t *testing.T) {
	data := []byte(`{
    "retCode": 0,
    "retMsg": "OK",
    "result": {
        "category": "spot",
        "list": [
            {
                "execId": "2100000000007764263",
                "symbol": "BTCUSDT",
                "price": "16618.49",
                "size": "0.00012",
                "side": "Buy",
                "time": "1672052955758",
                "isBlockTrade": false
            }
        ]
    },
    "retExtInfo": {},
    "time": 1672053054358
}`)
	var res models.GetPublicRecentTradesResponse
	err := json.Unmarshal(data, &res)
	assert.NoError(t, err)
	assert.Equal(t, "spot", res.Result.Category)
	assert.Len(t, res.Result.List, 1)
	assert.Equal(t, "2100000000007764263", res.Result.List[0].ExecId)
	assert.Equal(t, "BTCUSDT", res.Result.List[0].Symbol)
	assert.Equal(t, false, res.Result.List[0].IsBlockTrade)
}

func TestParseOpenInterestsResponse(t *testing.T) {
	data := []byte(`{
    "retCode": 0,
    "retMsg": "OK",
    "result": {
        "symbol": "BTCUSD",
        "category": "inverse",
        "list": [
            {"openInterest": "461134384.00000000", "timestamp": "1669571400000"},
            {"openInterest": "461134292.00000000", "timestamp": "1669571100000"}
        ],
        "nextPageCursor": ""
    },
    "retExtInfo": {},
    "time": 1672053548579
}`)
	var res models.GetOpenInterestsResponse
	err := json.Unmarshal(data, &res)
	assert.NoError(t, err)
	assert.Equal(t, "BTCUSD", res.Result.Symbol)
	assert.Equal(t, "inverse", res.Result.Category)
	assert.Len(t, res.Result.List, 2)
	assert.Equal(t, "461134384.00000000", res.Result.List[0].OpenInterest)
}
