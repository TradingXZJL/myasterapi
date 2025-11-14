package myasterapi

// Ws账户推送相关

// aster SPOT ws账户推送  SpotUserDataStreamPost rest现货创建一个listenKey (USER_STREAM)
func (client *SpotRestClient) NewSpotUserDataStreamPost() *SpotUserDataStreamPostApi {
	return &SpotUserDataStreamPostApi{
		client: client,
		req:    &SpotUserDataStreamPostReq{},
	}
}
func (api *SpotUserDataStreamPostApi) Do() (*SpotUserDataStreamPostRes, error) {
	url := asterHandlerRequestApi(SPOT, api.req, SpotApiMap[SpotUserDataStreamPost])
	return asterCallApiWithSecret[SpotUserDataStreamPostRes](api.client.c, url, POST)
}

// aster SPOT ws账户推送  SpotUserDataStreamPut rest现货延长listenKey有效期 (USER_STREAM)
func (client *SpotRestClient) NewSpotUserDataStreamPut() *SpotUserDataStreamPutApi {
	return &SpotUserDataStreamPutApi{
		client: client,
		req:    &SpotUserDataStreamPutReq{},
	}
}
func (api *SpotUserDataStreamPutApi) Do() (*SpotUserDataStreamPutRes, error) {
	url := asterHandlerRequestApi(SPOT, api.req, SpotApiMap[SpotUserDataStreamPut])
	return asterCallApiWithSecret[SpotUserDataStreamPutRes](api.client.c, url, PUT)
}

// aster SPOT ws账户推送  SpotUserDataStreamDelete rest现货关闭listenKey (USER_STREAM)
func (client *SpotRestClient) NewSpotUserDataStreamDelete() *SpotUserDataStreamDeleteApi {
	return &SpotUserDataStreamDeleteApi{
		client: client,
		req:    &SpotUserDataStreamDeleteReq{},
	}
}
func (api *SpotUserDataStreamDeleteApi) Do() (*SpotUserDataStreamDeleteRes, error) {
	url := asterHandlerRequestApi(SPOT, api.req, SpotApiMap[SpotUserDataStreamDelete])
	return asterCallApiWithSecret[SpotUserDataStreamDeleteRes](api.client.c, url, DELETE)
}

// aster SPOT ws账户推送 SpotMarginUserDataStreamPost rest现货杠杆创建一个listenKey (USER_STREAM)
func (client *SpotRestClient) NewSpotMarginUserDataStreamPost() *SpotMarginUserDataStreamPostApi {
	return &SpotMarginUserDataStreamPostApi{
		client: client,
		req:    &SpotMarginUserDataStreamPostReq{},
	}
}
func (api *SpotMarginUserDataStreamPostApi) Do() (*SpotMarginUserDataStreamPostRes, error) {
	url := asterHandlerRequestApi(SPOT, api.req, SpotApiMap[SpotMarginUserDataStreamPost])
	return asterCallApiWithSecret[SpotMarginUserDataStreamPostRes](api.client.c, url, POST)
}

// aster SPOT ws账户推送  SpotMarginUserDataStreamPut rest现货杠杆延长listenKey有效期 (USER_STREAM)
func (client *SpotRestClient) NewSpotMarginUserDataStreamPut() *SpotMarginUserDataStreamPutApi {
	return &SpotMarginUserDataStreamPutApi{
		client: client,
		req:    &SpotMarginUserDataStreamPutReq{},
	}
}
func (api *SpotMarginUserDataStreamPutApi) Do() (*SpotMarginUserDataStreamPutRes, error) {
	url := asterHandlerRequestApi(SPOT, api.req, SpotApiMap[SpotMarginUserDataStreamPut])
	return asterCallApiWithSecret[SpotMarginUserDataStreamPutRes](api.client.c, url, PUT)
}

// aster SPOT ws账户推送  SpotMarginUserDataStreamDelete rest现货杠杆关闭listenKey (USER_STREAM)
func (client *SpotRestClient) NewSpotMarginUserDataStreamDelete() *SpotMarginUserDataStreamDeleteApi {
	return &SpotMarginUserDataStreamDeleteApi{
		client: client,
		req:    &SpotMarginUserDataStreamDeleteReq{},
	}
}
func (api *SpotMarginUserDataStreamDeleteApi) Do() (*SpotMarginUserDataStreamDeleteRes, error) {
	url := asterHandlerRequestApi(SPOT, api.req, SpotApiMap[SpotMarginUserDataStreamDelete])
	return asterCallApiWithSecret[SpotMarginUserDataStreamDeleteRes](api.client.c, url, DELETE)
}

// aster SPOT ws账户推送  SpotMarginIsolatedUserDataStreamPost rest现货杠杆逐仓创建一个listenKey (USER_STREAM)
func (client *SpotRestClient) NewSpotMarginIsolatedUserDataStreamPost() *SpotMarginIsolatedUserDataStreamPostApi {
	return &SpotMarginIsolatedUserDataStreamPostApi{
		client: client,
		req:    &SpotMarginIsolatedUserDataStreamPostReq{},
	}
}
func (api *SpotMarginIsolatedUserDataStreamPostApi) Do() (*SpotMarginIsolatedUserDataStreamPostRes, error) {
	url := asterHandlerRequestApi(SPOT, api.req, SpotApiMap[SpotMarginIsolatedUserDataStreamPost])
	return asterCallApiWithSecret[SpotMarginIsolatedUserDataStreamPostRes](api.client.c, url, POST)
}

// aster SPOT ws账户推送  SpotMarginIsolatedUserDataStreamPut rest现货杠杆逐仓延长listenKey有效期 (USER_STREAM)
func (client *SpotRestClient) NewSpotMarginIsolatedUserDataStreamPut() *SpotMarginIsolatedUserDataStreamPutApi {
	return &SpotMarginIsolatedUserDataStreamPutApi{
		client: client,
		req:    &SpotMarginIsolatedUserDataStreamPutReq{},
	}
}
func (api *SpotMarginIsolatedUserDataStreamPutApi) Do() (*SpotMarginIsolatedUserDataStreamPutRes, error) {
	url := asterHandlerRequestApi(SPOT, api.req, SpotApiMap[SpotMarginIsolatedUserDataStreamPut])
	return asterCallApiWithSecret[SpotMarginIsolatedUserDataStreamPutRes](api.client.c, url, PUT)
}

// aster SPOT ws账户推送  SpotMarginIsolatedUserDataStreamDelete rest现货杠杆逐仓关闭listenKey (USER_STREAM)
func (client *SpotRestClient) NewSpotMarginIsolatedUserDataStreamDelete() *SpotMarginIsolatedUserDataStreamDeleteApi {
	return &SpotMarginIsolatedUserDataStreamDeleteApi{
		client: client,
		req:    &SpotMarginIsolatedUserDataStreamDeleteReq{},
	}
}
func (api *SpotMarginIsolatedUserDataStreamDeleteApi) Do() (*SpotMarginIsolatedUserDataStreamDeleteRes, error) {
	url := asterHandlerRequestApi(SPOT, api.req, SpotApiMap[SpotMarginIsolatedUserDataStreamDelete])
	return asterCallApiWithSecret[SpotMarginIsolatedUserDataStreamDeleteRes](api.client.c, url, DELETE)
}
