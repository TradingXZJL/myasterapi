package myasterapi

import (
	"fmt"
	"strings"
)

func getTradeParam(symbol string) string {
	return fmt.Sprintf("%s@trade", strings.ToLower(symbol))
}

// 订阅归集交易流 如: "BTCUSDT"
func (ws *WsStreamClient) SubscribeTrade(symbol string) (*Subscription[WsTrade], error) {
	return ws.SubscribeTradeMultiple([]string{symbol})
}

// 批量订阅归集交易流 如: []string{"BTCUSDT","ETHUSDT"}
func (ws *WsStreamClient) SubscribeTradeMultiple(symbols []string) (*Subscription[WsTrade], error) {
	params := []string{}
	for _, symbol := range symbols {
		param := getTradeParam(symbol)
		params = append(params, param)
	}
	doSub, err := sendMsg[SubscribeResult](ws, 0, SUBSCRIBE, params)
	if err != nil {
		return nil, err
	}
	ws.commonSubMap.Store(doSub.ID, doSub)

	select {
	case err := <-doSub.ErrChan():
		log.Error("SubscribeTrade Error: ", err)
		return nil, err
	case subResult := <-doSub.ResultChan():
		if subResult.Error != "" {
			log.Error(subResult.Error)
			return nil, fmt.Errorf(subResult.Error)
		}
		log.Infof("SubscribeTrade Success: params:%v result:%v", doSub.Params, subResult)
		sub := &Subscription[WsTrade]{
			ID:         doSub.ID,
			Method:     SUBSCRIBE,
			Params:     params,
			resultChan: make(chan WsTrade),
			errChan:    make(chan error),
			closeChan:  make(chan struct{}),
			ws:         ws,
		}
		for _, param := range params {
			ws.tradeSubMap.Store(param, sub)
		}
		return sub, nil
	}
}
