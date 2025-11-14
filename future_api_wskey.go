package myasterapi

// Ws账户推送相关
// aster FUTURE FutureListenKeyPost rest生成listenKey (USER_STREAM)
func (client *FutureRestClient) NewFutureListenKeyPost() *FutureListenKeyPostApi {
	return &FutureListenKeyPostApi{
		client: client,
		req:    &FutureListenKeyPostReq{},
	}
}
func (api *FutureListenKeyPostApi) Do() (*FutureListenKeyPostRes, error) {
	url := asterHandlerRequestApiWithSecret(FUTURE, api.req, FutureApiMap[FutureListenKeyPost], api.client.c.ApiSecret)
	return asterCallApiWithSecret[FutureListenKeyPostRes](api.client.c, url, POST)
}

// aster FUTURE FutureListenKeyPut rest延长listenKey有效期 (USER_STREAM)
func (client *FutureRestClient) NewFutureListenKeyPut() *FutureListenKeyPutApi {
	return &FutureListenKeyPutApi{
		client: client,
		req:    &FutureListenKeyPutReq{},
	}
}
func (api *FutureListenKeyPutApi) Do() (*FutureListenKeyPutRes, error) {
	url := asterHandlerRequestApiWithSecret(FUTURE, api.req, FutureApiMap[FutureListenKeyPut], api.client.c.ApiSecret)
	return asterCallApiWithSecret[FutureListenKeyPutRes](api.client.c, url, PUT)
}

// aster FUTURE FutureListenKeyDelete rest关闭listenKey (USER_STREAM)
func (client *FutureRestClient) NewFutureListenKeyDelete() *FutureListenKeyDeleteApi {
	return &FutureListenKeyDeleteApi{
		client: client,
		req:    &FutureListenKeyDeleteReq{},
	}
}
func (api *FutureListenKeyDeleteApi) Do() (*FutureListenKeyDeleteRes, error) {
	url := asterHandlerRequestApiWithSecret(FUTURE, api.req, FutureApiMap[FutureListenKeyDelete], api.client.c.ApiSecret)
	return asterCallApiWithSecret[FutureListenKeyDeleteRes](api.client.c, url, DELETE)
}
