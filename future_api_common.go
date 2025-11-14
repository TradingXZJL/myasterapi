package myasterapi

// 通用接口
// aster FUTURE FuturePing rest测试服务器连通性 (NONE)
func (client *FutureRestClient) NewPing() *FuturePingApi {
	return &FuturePingApi{
		client: client,
		req:    &FuturePingReq{},
	}
}
func (api *FuturePingApi) Do() (*FuturePingRes, error) {
	url := asterHandlerRequestApi(FUTURE, api.req, FutureApiMap[FuturePing])
	return asterCallApiWithSecret[FuturePingRes](api.client.c, url, GET)
}

// aster FUTURE FutureTime rest获取服务器时间 (NONE)
func (client *FutureRestClient) NewServerTime() *FutureServerTimeApi {
	return &FutureServerTimeApi{
		client: client,
		req:    &FutureServerTimeReq{},
	}
}
func (api *FutureServerTimeApi) Do() (*FutureTimeRes, error) {
	url := asterHandlerRequestApi(FUTURE, api.req, FutureApiMap[FutureServerTime])
	return asterCallApiWithSecret[FutureTimeRes](api.client.c, url, GET)
}

// aster FUTURE FutureExchangeInfo rest交易规范信息和交易对信息 (NONE)
func (client *FutureRestClient) NewExchangeInfo() *FutureExchangeInfoApi {
	return &FutureExchangeInfoApi{
		client: client,
		req:    &FutureExchangeInfoReq{},
	}
}
func (api *FutureExchangeInfoApi) Do() (*FutureExchangeInfoRes, error) {
	url := asterHandlerRequestApi(FUTURE, api.req, FutureApiMap[FutureExchangeInfo])
	return asterCallApiWithSecret[FutureExchangeInfoRes](api.client.c, url, GET)
}
