package myasterapi

import "time"

// 交易接口
// aster FUTURE FutureOpenOrders rest查询当前挂单 (USER_DATA)
func (client *FutureRestClient) NewOpenOrders() *FutureOpenOrdersApi {
	return &FutureOpenOrdersApi{
		client: client,
		req:    &FutureOpenOrdersReq{},
	}
}
func (api *FutureOpenOrdersApi) Do() (*FutureOpenOrdersRes, error) {
	api.Timestamp(time.Now().UnixMilli() + serverTimeDelta)
	url := asterHandlerRequestApiWithSecret(FUTURE, api.req, FutureApiMap[FutureOpenOrders], api.client.c.ApiSecret)
	return asterCallApiWithSecret[FutureOpenOrdersRes](api.client.c, url, GET)
}

// aster FUTURE FutureAllOrders rest查询所有订单 (USER_DATA)
func (client *FutureRestClient) NewAllOrders() *FutureAllOrdersApi {
	return &FutureAllOrdersApi{
		client: client,
		req:    &FutureAllOrdersReq{},
	}
}
func (api *FutureAllOrdersApi) Do() (*FutureAllOrdersRes, error) {
	api.Timestamp(time.Now().UnixMilli() + serverTimeDelta)
	url := asterHandlerRequestApiWithSecret(FUTURE, api.req, FutureApiMap[FutureAllOrders], api.client.c.ApiSecret)
	return asterCallApiWithSecret[FutureAllOrdersRes](api.client.c, url, GET)
}

// aster FUTURE FutureOrderPost rest下单 (TRADE)
func (client *FutureRestClient) NewFutureOrderPost() *FutureOrderPostApi {
	return &FutureOrderPostApi{
		client: client,
		req:    &FutureOrderPostReq{},
	}
}
func (api *FutureOrderPostApi) Do() (*FutureOrderPostRes, error) {
	api.Timestamp(time.Now().UnixMilli() + serverTimeDelta)
	url := asterHandlerRequestApiWithSecret(FUTURE, api.req, FutureApiMap[FutureOrderPost], api.client.c.ApiSecret)
	return asterCallApiWithSecret[FutureOrderPostRes](api.client.c, url, POST)
}

// aster FUTURE FutureOrderPut rest修改订单 (TRADE)
func (client *FutureRestClient) NewFutureOrderPut() *FutureOrderPutApi {
	return &FutureOrderPutApi{
		client: client,
		req:    &FutureOrderPutReq{},
	}
}
func (api *FutureOrderPutApi) Do() (*FutureOrderPutRes, error) {
	api.Timestamp(time.Now().UnixMilli() + serverTimeDelta)
	url := asterHandlerRequestApiWithSecret(FUTURE, api.req, FutureApiMap[FutureOrderPut], api.client.c.ApiSecret)
	return asterCallApiWithSecret[FutureOrderPutRes](api.client.c, url, PUT)
}

// aster FUTURE FutureOrderGet  rest查询订单 (USER_DATA)
func (client *FutureRestClient) NewFutureOrderGet() *FutureOrderGetApi {
	return &FutureOrderGetApi{
		client: client,
		req:    &FutureOrderGetReq{},
	}
}
func (api *FutureOrderGetApi) Do() (*FutureOrderGetRes, error) {
	api.Timestamp(time.Now().UnixMilli() + serverTimeDelta)
	url := asterHandlerRequestApiWithSecret(FUTURE, api.req, FutureApiMap[FutureOrderGet], api.client.c.ApiSecret)
	return asterCallApiWithSecret[FutureOrderGetRes](api.client.c, url, GET)
}

// aster FUTURE FutureOrderDelete rest撤销订单 (TRADE)
func (client *FutureRestClient) NewFutureOrderDelete() *FutureOrderDeleteApi {
	return &FutureOrderDeleteApi{
		client: client,
		req:    &FutureOrderDeleteReq{},
	}
}
func (api *FutureOrderDeleteApi) Do() (*FutureOrderDeleteRes, error) {
	api.Timestamp(time.Now().UnixMilli() + serverTimeDelta)
	url := asterHandlerRequestApiWithSecret(FUTURE, api.req, FutureApiMap[FutureOrderDelete], api.client.c.ApiSecret)
	return asterCallApiWithSecret[FutureOrderDeleteRes](api.client.c, url, DELETE)
}

// aster FUTURE FutureBatchOrdersPost rest批量下单 (TRADE)
func (client *FutureRestClient) NewFutureBatchOrdersPost() *FutureBatchOrdersPostApi {
	return &FutureBatchOrdersPostApi{
		client: client,
		req:    &FutureBatchOrdersPostReq{},
	}
}
func (api *FutureBatchOrdersPostApi) Do() (*FutureBatchOrdersPostRes, error) {
	api.Timestamp(time.Now().UnixMilli() + serverTimeDelta)
	url := asterHandlerRequestApiWithSecret(FUTURE, api.req, FutureApiMap[FutureBatchOrdersPost], api.client.c.ApiSecret)
	return asterCallApiWithSecret[FutureBatchOrdersPostRes](api.client.c, url, POST)
}

// aster FUTURE FutureBatchOrdersPut rest批量修改订单 (TRADE)
func (client *FutureRestClient) NewFutureBatchOrdersPut() *FutureBatchOrdersPutApi {
	return &FutureBatchOrdersPutApi{
		client: client,
		req:    &FutureBatchOrdersPutReq{},
	}
}
func (api *FutureBatchOrdersPutApi) Do() (*FutureBatchOrdersPutRes, error) {
	api.Timestamp(time.Now().UnixMilli() + serverTimeDelta)
	url := asterHandlerRequestApiWithSecret(FUTURE, api.req, FutureApiMap[FutureBatchOrdersPut], api.client.c.ApiSecret)
	return asterCallApiWithSecret[FutureBatchOrdersPutRes](api.client.c, url, PUT)
}

// aster FUTURE FutureBatchOrdersDelete rest批量撤销订单 (TRADE)
func (client *FutureRestClient) NewFutureBatchOrdersDelete() *FutureBatchOrdersDeleteApi {
	return &FutureBatchOrdersDeleteApi{
		client: client,
		req:    &FutureBatchOrdersDeleteReq{},
	}
}
func (api *FutureBatchOrdersDeleteApi) Do() (*FutureBatchOrdersDeleteRes, error) {
	api.Timestamp(time.Now().UnixMilli() + serverTimeDelta)
	url := asterHandlerRequestApiWithSecret(FUTURE, api.req, FutureApiMap[FutureBatchOrdersDelete], api.client.c.ApiSecret)
	return asterCallApiWithSecret[FutureBatchOrdersDeleteRes](api.client.c, url, DELETE)
}

// aster FUTURE FutureUserTrades rest账户成交历史 (USER_DATA)
func (client *FutureRestClient) NewFutureUserTrades() *FutureUserTradesApi {
	return &FutureUserTradesApi{
		client: client,
		req:    &FutureUserTradesReq{},
	}
}
func (api *FutureUserTradesApi) Do() (*FutureUserTradesRes, error) {
	api.Timestamp(time.Now().UnixMilli() + serverTimeDelta)
	url := asterHandlerRequestApiWithSecret(FUTURE, api.req, FutureApiMap[FutureUserTrades], api.client.c.ApiSecret)
	return asterCallApiWithSecret[FutureUserTradesRes](api.client.c, url, GET)
}

// aster FUTURE  FutureCommissionRate rest查询用户当前的手续费率 (USER_DATA)
func (client *FutureRestClient) NewFutureCommissionRate() *FutureCommissionRateApi {
	return &FutureCommissionRateApi{
		client: client,
		req:    &FutureCommissionRateReq{},
	}
}
func (api *FutureCommissionRateApi) Do() (*FutureCommissionRateRes, error) {
	api.Timestamp(time.Now().UnixMilli() + serverTimeDelta)
	url := asterHandlerRequestApiWithSecret(FUTURE, api.req, FutureApiMap[FutureCommissionRate], api.client.c.ApiSecret)
	return asterCallApiWithSecret[FutureCommissionRateRes](api.client.c, url, GET)
}
