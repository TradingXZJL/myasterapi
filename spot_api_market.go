package myasterapi

// 行情接口
// aster SPOT spotTickerPrice rest价格 (NONE)
func (client *SpotRestClient) NewSpotTickerPrice() *SpotTickerPriceApi {
	return &SpotTickerPriceApi{
		client: client,
		req:    &SpotTickerPriceReq{},
	}
}
func (api *SpotTickerPriceApi) Do() (*SpotTickerPriceRes, error) {
	url := asterHandlerRequestApi(SPOT, api.req, SpotApiMap[SpotTickerPrice])
	if api.req.Symbol != nil && *api.req.Symbol != "" {
		res, err := asterCallApiWithSecretForJson[SpotTickerPriceResRow](api.client.c, url, GET)
		if err != nil {
			return nil, err
		}
		return &SpotTickerPriceRes{*res}, nil
	} else {
		return asterCallApiWithSecretForJson[SpotTickerPriceRes](api.client.c, url, GET)
	}

}

// aster SPOT spotKlines restK线数据 (NONE)
func (client *SpotRestClient) NewSpotKlines() *SpotKlinesApi {
	return &SpotKlinesApi{
		client: client,
		req:    &SpotKlinesReq{},
	}
}
func (api *SpotKlinesApi) Do() (*KlinesRes, error) {
	url := asterHandlerRequestApi(SPOT, api.req, SpotApiMap[SpotKlines])
	res, err := asterCallApiWithSecretForJson[KlinesMiddle](api.client.c, url, GET)
	if err != nil {
		return nil, err
	}
	res2 := res.ConvertToRes()
	return &res2, nil
}

// aster SPOT spotDepth rest深度信息 (NONE)
func (client *SpotRestClient) NewSpotDepth() *SpotDepthApi {
	return &SpotDepthApi{
		client: client,
		req:    &SpotDepthReq{},
	}
}
func (api *SpotDepthApi) Do() (*SpotDepthRes, error) {
	url := asterHandlerRequestApi(SPOT, api.req, SpotApiMap[SpotDepth])
	res, err := asterCallApiWithSecretForJson[SpotDepthResMiddle](api.client.c, url, GET)
	if err != nil {
		return nil, err
	}
	return res.ConvertToRes(), nil
}

// aster SPOT SpotTrades rest最近成交 (NONE)
func (client *SpotRestClient) NewSpotTrades() *SpotTradesApi {
	return &SpotTradesApi{
		client: client,
		req:    &SpotTradesReq{},
	}
}
func (api *SpotTradesApi) Do() (*SpotTradesRes, error) {
	url := asterHandlerRequestApi(SPOT, api.req, SpotApiMap[SpotTrades])
	return asterCallApiWithSecretForJson[SpotTradesRes](api.client.c, url, GET)
}

// aster SPOT spotHistoricalTrades rest历史成交 (NONE)
func (client *SpotRestClient) NewSpotHistoricalTrades() *SpotHistoricalTradesApi {
	return &SpotHistoricalTradesApi{
		client: client,
		req:    &SpotHistoricalTradesReq{},
	}
}
func (api *SpotHistoricalTradesApi) Do() (*SpotHistoricalTradesRes, error) {
	url := asterHandlerRequestApi(SPOT, api.req, SpotApiMap[SpotHistoricalTrades])
	return asterCallApiWithSecretForJson[SpotHistoricalTradesRes](api.client.c, url, GET)
}

// aster SPOT spotAggTrades rest近期成交(归集)(NONE)
func (client *SpotRestClient) NewSpotAggTrades() *SpotAggTradesApi {
	return &SpotAggTradesApi{
		client: client,
		req:    &SpotAggTradesReq{},
	}
}
func (api *SpotAggTradesApi) Do() (*SpotAggTradesRes, error) {
	url := asterHandlerRequestApi(SPOT, api.req, SpotApiMap[SpotAggTrades])
	return asterCallApiWithSecretForJson[SpotAggTradesRes](api.client.c, url, GET)
}

// aster SPOT spotAvgPrice rest当前平均价格 (NONE)
func (client *SpotRestClient) NewSpotAvgPrice() *SpotAvgPriceApi {
	return &SpotAvgPriceApi{
		client: client,
		req:    &SpotAvgPriceReq{},
	}
}
func (api *SpotAvgPriceApi) Do() (*SpotAvgPriceRes, error) {
	url := asterHandlerRequestApi(SPOT, api.req, SpotApiMap[SpotAvgPrice])
	return asterCallApiWithSecretForJson[SpotAvgPriceRes](api.client.c, url, GET)
}

// aster SPOT spotUiKlines restUI K线数据 (NONE)
func (client *SpotRestClient) NewSpotUiKlines() *SpotUiKlinesApi {
	return &SpotUiKlinesApi{
		client: client,
		req:    &SpotKlinesReq{},
	}
}
func (api *SpotUiKlinesApi) Do() (*KlinesRes, error) {
	url := asterHandlerRequestApi(SPOT, api.req, SpotApiMap[SpotKlines])
	res, err := asterCallApiWithSecretForJson[KlinesMiddle](api.client.c, url, GET)
	if err != nil {
		return nil, err
	}
	res2 := res.ConvertToRes()
	return &res2, nil
}

// aster SPOT spotTicker24hr rest24hr价格变动情况 (NONE)
func (client *SpotRestClient) NewSpotTicker24hr() *SpotTicker24hrApi {
	return &SpotTicker24hrApi{
		client: client,
		req:    &SpotTicker24hrReq{},
	}
}
func (api *SpotTicker24hrApi) Do() (*SpotTicker24hrRes, error) {
	url := asterHandlerRequestApi(SPOT, api.req, SpotApiMap[SpotTicker24hr])
	if api.req.Symbol != nil && *api.req.Symbol != "" {
		res, err := asterCallApiWithSecretForJson[SpotTicker24hrResRow](api.client.c, url, GET)
		if err != nil {
			return nil, err
		}
		return &SpotTicker24hrRes{*res}, nil
	} else {
		return asterCallApiWithSecretForJson[SpotTicker24hrRes](api.client.c, url, GET)
	}
}

// aster SPOT spotTickerBookTicker rest当前最优挂单 (NONE)
func (client *SpotRestClient) NewSpotTickerBookTicker() *SpotTickerBookTickerApi {
	return &SpotTickerBookTickerApi{
		client: client,
		req:    &SpotTickerBookTickerReq{},
	}
}
func (api *SpotTickerBookTickerApi) Do() (*SpotTickerBookTickerRes, error) {
	url := asterHandlerRequestApi(SPOT, api.req, SpotApiMap[SpotTickerBookTicker])
	if api.req.Symbol != nil && *api.req.Symbol != "" {
		res, err := asterCallApiWithSecretForJson[SpotTickerBookTickerResRow](api.client.c, url, GET)
		if err != nil {
			return nil, err
		}
		return &SpotTickerBookTickerRes{*res}, nil
	} else {
		return asterCallApiWithSecretForJson[SpotTickerBookTickerRes](api.client.c, url, GET)
	}
}

// aster SPOT spotTicker rest滚动窗口价格变动统计
func (client *SpotRestClient) NewSpotTicker() *SpotTickerApi {
	return &SpotTickerApi{
		client: client,
		req:    &SpotTickerReq{},
	}
}
func (api *SpotTickerApi) Do() (*SpotTickerRes, error) {
	url := asterHandlerRequestApi(SPOT, api.req, SpotApiMap[SpotTicker])
	if api.req.Symbol != nil && *api.req.Symbol != "" {
		res, err := asterCallApiWithSecretForJson[SpotTickerResRow](api.client.c, url, GET)
		if err != nil {
			return nil, err
		}
		return &SpotTickerRes{*res}, nil
	} else {
		return asterCallApiWithSecretForJson[SpotTickerRes](api.client.c, url, GET)
	}
}
