package myasterapi

// 行情接口
// aster FUTURE FutureKlines restK线数据 (MARKET_DATA)
func (client *FutureRestClient) NewFutureKlines() *FutureKlinesApi {
	return &FutureKlinesApi{
		client: client,
		req:    &FutureKlinesReq{},
	}
}
func (api *FutureKlinesApi) Do() (*KlinesRes, error) {
	url := asterHandlerRequestApi(FUTURE, api.req, FutureApiMap[FutureKlines])
	res, err := asterCallApiWithSecret[KlinesMiddle](api.client.c, url, GET)
	if err != nil {
		return nil, err
	}
	res2 := res.ConvertToRes()
	return &res2, nil
}

// aster FUTURE FutureDepth rest深度信息 (MARKET_DATA)
func (client *FutureRestClient) NewFutureDepth() *FutureDepthApi {
	return &FutureDepthApi{
		client: client,
		req:    &FutureDepthReq{},
	}
}
func (api *FutureDepthApi) Do() (*FutureDepthRes, error) {
	url := asterHandlerRequestApi(FUTURE, api.req, FutureApiMap[FutureDepth])
	res, err := asterCallApiWithSecret[FutureDepthResMiddle](api.client.c, url, GET)
	if err != nil {
		return nil, err
	}
	return res.ConvertToRes(), nil
}

// aster FUTURE FutureTrades rest最新成交 (MARKET_DATA)
func (client *FutureRestClient) NewFutureTrades() *FutureTradesApi {
	return &FutureTradesApi{
		client: client,
		req:    &FutureTradesReq{},
	}
}
func (api *FutureTradesApi) Do() (*FutureTradesRes, error) {
	url := asterHandlerRequestApi(FUTURE, api.req, FutureApiMap[FutureTrades])
	return asterCallApiWithSecret[FutureTradesRes](api.client.c, url, GET)
}

// aster FUTURE FutureHistoricalTrades rest历史成交 (MARKET_DATA)
func (client *FutureRestClient) NewFutureHistoricalTrades() *FutureHistoricalTradesApi {
	return &FutureHistoricalTradesApi{
		client: client,
		req:    &FutureHistoricalTradesReq{},
	}
}
func (api *FutureHistoricalTradesApi) Do() (*FutureHistoricalTradesRes, error) {
	url := asterHandlerRequestApi(FUTURE, api.req, FutureApiMap[FutureHistoricalTrades])
	return asterCallApiWithSecret[FutureHistoricalTradesRes](api.client.c, url, GET)
}

// aster FUTURE FutureAggTrades rest近期成交(归集) (MARKET_DATA)
func (client *FutureRestClient) NewFutureAggTrades() *FutureAggTradesApi {
	return &FutureAggTradesApi{
		client: client,
		req:    &FutureAggTradesReq{},
	}
}
func (api *FutureAggTradesApi) Do() (*FutureAggTradesRes, error) {
	url := asterHandlerRequestApi(FUTURE, api.req, FutureApiMap[FutureAggTrades])
	return asterCallApiWithSecret[FutureAggTradesRes](api.client.c, url, GET)
}

// aster FUTURE FuturePremiumIndex rest最新标记价格和资金费率 (MARKET_DATA)
func (client *FutureRestClient) NewFuturePremiumIndex() *FuturePremiumIndexApi {
	return &FuturePremiumIndexApi{
		client: client,
		req:    &FuturePremiumIndexReq{},
	}
}
func (api *FuturePremiumIndexApi) Do() (*FuturePremiumIndexRes, error) {
	url := asterHandlerRequestApi(FUTURE, api.req, FutureApiMap[FuturePremiumIndex])
	if api.req.Symbol != nil && *api.req.Symbol != "" {
		res, err := asterCallApiWithSecret[FuturePremiumIndexResRow](api.client.c, url, GET)
		if err != nil {
			return nil, err
		}
		return &FuturePremiumIndexRes{*res}, nil
	} else {
		return asterCallApiWithSecret[FuturePremiumIndexRes](api.client.c, url, GET)
	}

}

// aster FUTURE FutureFundingRate rest查询资金费率历史 (MARKET_DATA)
func (client *FutureRestClient) NewFutureFundingRate() *FutureFundingRateApi {
	return &FutureFundingRateApi{
		client: client,
		req:    &FutureFundingRateReq{},
	}
}
func (api *FutureFundingRateApi) Do() (*FutureFundingRateRes, error) {
	url := asterHandlerRequestApi(FUTURE, api.req, FutureApiMap[FutureFundingRate])
	return asterCallApiWithSecret[FutureFundingRateRes](api.client.c, url, GET)
}

// aster FUTURE FutureFundingInfo rest查询资金费率信息 (MARKET_DATA)
func (client *FutureRestClient) NewFutureFundingInfo() *FutureFundingInfoApi {
	return &FutureFundingInfoApi{
		client: client,
		req:    &FutureFundingInfoReq{},
	}
}
func (api *FutureFundingInfoApi) Do() (*FutureFundingInfoRes, error) {
	url := asterHandlerRequestApi(FUTURE, api.req, FutureApiMap[FutureFundingInfo])
	return asterCallApiWithSecret[FutureFundingInfoRes](api.client.c, url, GET)
}

// aster FUTURE FutureTicker24hr rest24hr价格变动情况 (MARKET_DATA)
func (client *FutureRestClient) NewFutureTicker24hr() *FutureTicker24hrApi {
	return &FutureTicker24hrApi{
		client: client,
		req:    &FutureTicker24hrReq{},
	}
}
func (api *FutureTicker24hrApi) Do() (*FutureTicker24hrRes, error) {
	url := asterHandlerRequestApi(FUTURE, api.req, FutureApiMap[FutureTicker24hr])
	if api.req.Symbol != nil && *api.req.Symbol != "" {
		res, err := asterCallApiWithSecret[FutureTicker24hrResRow](api.client.c, url, GET)
		if err != nil {
			return nil, err
		}
		return &FutureTicker24hrRes{*res}, nil
	} else {
		return asterCallApiWithSecret[FutureTicker24hrRes](api.client.c, url, GET)
	}
}

// aster FUTURE FutureTickerPrice rest最新价格 (MARKET_DATA)
func (client *FutureRestClient) NewFutureTickerPrice() *FutureTickerPriceApi {
	return &FutureTickerPriceApi{
		client: client,
		req:    &FutureTickerPriceReq{},
	}
}
func (api *FutureTickerPriceApi) Do() (*FutureTickerPriceRes, error) {
	url := asterHandlerRequestApi(FUTURE, api.req, FutureApiMap[FutureTickerPrice])
	if api.req.Symbol != nil && *api.req.Symbol != "" {
		res, err := asterCallApiWithSecret[FutureTickerPriceResRow](api.client.c, url, GET)
		if err != nil {
			return nil, err
		}
		return &FutureTickerPriceRes{*res}, nil
	} else {
		return asterCallApiWithSecret[FutureTickerPriceRes](api.client.c, url, GET)
	}
}

// aster FUTURE FutureTickerBookTicker rest当前最优挂单 (MARKET_DATA)
func (client *FutureRestClient) NewFutureTickerBookTicker() *FutureTickerBookTickerApi {
	return &FutureTickerBookTickerApi{
		client: client,
		req:    &FutureTickerBookTickerReq{},
	}
}
func (api *FutureTickerBookTickerApi) Do() (*FutureTickerBookTickerRes, error) {
	url := asterHandlerRequestApi(FUTURE, api.req, FutureApiMap[FutureTickerBookTicker])
	if api.req.Symbol != nil && *api.req.Symbol != "" {
		res, err := asterCallApiWithSecret[FutureTickerBookTickerResRow](api.client.c, url, GET)
		if err != nil {
			return nil, err
		}
		return &FutureTickerBookTickerRes{*res}, nil
	} else {
		return asterCallApiWithSecret[FutureTickerBookTickerRes](api.client.c, url, GET)
	}
}

// aster FUTURE FutureDataBasis rest基差数据 (MARKET_DATA)
func (client *FutureRestClient) NewFutureDataBasis() *FutureDataBasisApi {
	return &FutureDataBasisApi{
		client: client,
		req:    &FutureDataBasisReq{},
	}
}
func (api *FutureDataBasisApi) Do() (*FutureDataBasisRes, error) {
	url := asterHandlerRequestApi(FUTURE, api.req, FutureApiMap[FutureDataBasis])
	return asterCallApiWithSecret[FutureDataBasisRes](api.client.c, url, GET)
}
