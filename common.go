package myasterapi

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"sort"
	"sync"
	"time"

	"net/url"
	"reflect"
	"strconv"
	"strings"

	jsoniter "github.com/json-iterator/go"
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
)

const (
	BIT_BASE_10 = 10
	BIT_SIZE_64 = 64
	//BIT_SIZE_32 = 32
)

type RequestType string

const (
	GET    = "GET"
	POST   = "POST"
	DELETE = "DELETE"
	PUT    = "PUT"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

var log = logrus.New()

func SetLogger(logger *logrus.Logger) {
	log = logger
}

var httpTimeout = 100 * time.Second

func SetHttpTimeout(timeout time.Duration) {
	httpTimeout = timeout
}

func GetPointer[T any](v T) *T {
	return &v
}

func HmacSha256(secret, data string) string {
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(data))
	sha := hex.EncodeToString(h.Sum(nil))
	return sha
}

type MyAster struct {
}

const (
	ASTER_API_SPOT_HTTP   = "sapi.asterdex.com"
	ASTER_API_FUTURE_HTTP = "fapi.asterdex.com"

	IS_GZIP = true
)

type NetType int

const (
	MAIN_NET NetType = iota
)

var NowNetType = MAIN_NET

func SetNetType(netType NetType) {
	NowNetType = netType
}

type ApiType int

const (
	SPOT ApiType = iota
	FUTURE
)

func (apiType *ApiType) String() string {
	switch *apiType {
	case SPOT:
		return "SPOT"
	case FUTURE:
		return "FUTURE"
	}
	return ""
}

type Client struct {
	ApiKey    string
	ApiSecret string
}
type RestClient struct {
	c *Client
}

type SpotRestClient RestClient
type FutureRestClient RestClient
type SwapRestClient RestClient
type PortfolioMarginRestClient RestClient

func (*MyAster) NewSpotRestClient(apiKey string, apiSecret string) *SpotRestClient {
	client := &SpotRestClient{
		&Client{
			ApiKey:    apiKey,
			ApiSecret: apiSecret,
		},
	}
	return client
}
func (*MyAster) NewFutureRestClient(apiKey string, apiSecret string) *FutureRestClient {
	client := &FutureRestClient{
		&Client{
			ApiKey:    apiKey,
			ApiSecret: apiSecret,
		},
	}
	return client
}
func (*MyAster) NewSwapRestClient(apiKey string, apiSecret string) *SwapRestClient {
	client := &SwapRestClient{
		&Client{
			ApiKey:    apiKey,
			ApiSecret: apiSecret,
		},
	}
	return client
}
func (*MyAster) NewPortfolioMarginClient(apiKey string, apiSecret string) *PortfolioMarginRestClient {
	client := &PortfolioMarginRestClient{
		&Client{
			ApiKey:    apiKey,
			ApiSecret: apiSecret,
		},
	}
	return client
}

var serverTimeDelta int64 = 0

func SetServerTimeDelta(delta int64) {
	serverTimeDelta = delta
}

// // 通用鉴权接口调用
// func asterCallApiWithSecret[T any](client *Client, url, method string) (*T, error) {
// 	// log.Warn(method, ": ", url)
// 	body, err := RequestWithHeader(url, method, map[string]string{
// 		"X-MBX-APIKEY": client.ApiKey,
// 		"Content-Type": "application/x-www-form-urlencoded",
// 	}, IS_GZIP)
// 	if err != nil {
// 		return nil, err
// 	}
// 	res, err := handlerCommonRest[T](body)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &res.Result, res.handlerError()
// }

// // 通用body鉴权接口调用
// func asterCallApiWithSecretForBody[T any](client *Client, url, method string, reqBody []byte) (*T, error) {
// 	// log.Warn(method, ": ", url)
// 	body, err := RequestWithHeaderAndBody(url, method, map[string]string{
// 		"X-MBX-APIKEY": client.ApiKey,
// 		"Content-Type": "application/x-www-form-urlencoded",
// 	}, reqBody, IS_GZIP)
// 	if err != nil {
// 		return nil, err
// 	}
// 	res, err := handlerCommonRest[T](body)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &res.Result, res.handlerError()
// }

func asterCallApiWithSecret[T any](client *Client, url, method string) (*T, error) {
	return asterCallApiWithSecretForJsonBody[T](client, url, method, []byte{})
}

// 通用json 无body接口调用
func asterCallApiWithSecretForJson[T any](client *Client, url, method string) (*T, error) {
	return asterCallApiWithSecretForJsonBody[T](client, url, method, []byte{})
}

// 通用json body鉴权接口调用
func asterCallApiWithSecretForJsonBody[T any](client *Client, url, method string, reqBody []byte) (*T, error) {
	// log.Warn(method, ": ", url)
	headerMap := map[string]string{
		"Content-Type": "application/json",
	}
	if client.ApiKey != "" {
		headerMap["X-MBX-APIKEY"] = client.ApiKey
	}
	body, err := RequestWithHeaderAndBody(url, method, headerMap, reqBody, IS_GZIP)
	if err != nil {
		return nil, err
	}
	res, err := handlerCommonRest[T](body)
	if err != nil {
		return nil, err
	}
	return &res.Result, res.handlerError()
}

// 通用鉴权接口封装
func asterHandlerRequestApiWithSecret[T any](apiType ApiType, request *T, name, secret string) string {
	query := asterHandlerReq(request)
	sign := HmacSha256(secret, query)

	u := url.URL{
		Scheme:   "https",
		Host:     AsterGetRestHostByApiType(apiType),
		Path:     name,
		RawQuery: query + "&signature=" + sign,
	}

	return u.String()
}

// 通用query body鉴权接口封装
func asterHandlerRequestApiWithSecretForBody[T any](apiType ApiType, request *T, name, secret string) ([]byte, string) {
	query := asterHandlerReq(request)

	sign := HmacSha256(secret, query)
	requestBody := []byte(query + "&signature=" + sign)
	u := url.URL{
		Scheme: "https",
		Host:   AsterGetRestHostByApiType(apiType),
		Path:   name,
	}

	log.Warn("reqUrl: ", u.String())
	log.Warn("reqBody: ", string(requestBody))
	return requestBody, u.String()
}

// 通用url param + json Body鉴权接口封装
func asterHandlerRequestApiWithSecretForUrlRequestAndJsonBody[T any](apiType ApiType, urlRequest *T, jsonBody []byte, name, secret string) ([]byte, string) {
	query := asterHandlerReq(urlRequest)
	signSource := query //+ string(jsonBody)
	sign := HmacSha256(secret, signSource)
	_ = sign

	u := url.URL{
		Scheme:   "https",
		Host:     AsterGetRestHostByApiType(apiType),
		Path:     name,
		RawQuery: query + "&signature=" + sign,
	}

	return jsonBody, u.String()
}

// 通用非鉴权接口封装
func asterHandlerRequestApi[T any](apiType ApiType, request *T, name string) string {
	query := asterHandlerReq(request)
	u := url.URL{
		Scheme:   "https",
		Host:     AsterGetRestHostByApiType(apiType),
		Path:     name,
		RawQuery: query,
	}
	// log.Debug(u.String())
	return u.String()
}

func asterHandlerReq[T any](req *T) string {
	var paramBuffer bytes.Buffer
	t := reflect.TypeOf(req)
	v := reflect.ValueOf(req)
	if v.IsNil() {
		return ""
	}
	t = t.Elem()
	v = v.Elem()
	count := v.NumField()
	for i := 0; i < count; i++ {
		paramName := t.Field(i).Tag.Get("json")
		paramName = strings.ReplaceAll(paramName, ",omitempty", "")
		switch v.Field(i).Elem().Kind() {
		case reflect.String:
			paramBuffer.WriteString(paramName + "=" + v.Field(i).Elem().String() + "&")
		case reflect.Int, reflect.Int32, reflect.Int64:
			paramBuffer.WriteString(paramName + "=" + strconv.FormatInt(v.Field(i).Elem().Int(), BIT_BASE_10) + "&")
		case reflect.Float32, reflect.Float64:
			paramBuffer.WriteString(paramName + "=" + decimal.NewFromFloat(v.Field(i).Elem().Float()).String() + "&")
		case reflect.Bool:
			paramBuffer.WriteString(paramName + "=" + strconv.FormatBool(v.Field(i).Elem().Bool()) + "&")
		case reflect.Struct:
			sv := reflect.ValueOf(v.Field(i).Interface())
			ToStringMethod := sv.MethodByName("String")
			params := make([]reflect.Value, 0)
			result := ToStringMethod.Call(params)
			paramBuffer.WriteString(paramName + "=" + result[0].String() + "&")
		case reflect.Slice:
			s := v.Field(i).Interface()
			d, _ := json.Marshal(s)
			paramBuffer.WriteString(paramName + "=" + url.QueryEscape(string(d)) + "&")
		case reflect.Invalid:
		default:
			log.Errorf("req type error %s:%s", paramName, v.Field(i).Elem().Kind())
		}
	}
	return strings.TrimRight(paramBuffer.String(), "&")
}

func asterHandlerWsApiReq[T any](req *T) string {
	var paramBuffer bytes.Buffer
	t := reflect.TypeOf(req)
	v := reflect.ValueOf(req)
	if v.IsNil() {
		return ""
	}
	t = t.Elem()
	v = v.Elem()
	count := v.NumField()

	sortMap := map[string]int{}
	sortName := make([]string, 0)
	for i := 0; i < count; i++ {
		paramName := t.Field(i).Tag.Get("json")
		paramName = strings.ReplaceAll(paramName, ",omitempty", "")
		sortName = append(sortName, paramName)
		sortMap[paramName] = i
	}

	sort.Strings(sortName)
	for _, name := range sortName {
		if i, ok := sortMap[name]; ok {
			paramName := t.Field(i).Tag.Get("json")
			paramName = strings.ReplaceAll(paramName, ",omitempty", "")
			switch v.Field(i).Elem().Kind() {
			case reflect.String:
				paramBuffer.WriteString(paramName + "=" + url.QueryEscape(v.Field(i).Elem().String()) + "&")
			case reflect.Int, reflect.Int32, reflect.Int64:
				paramBuffer.WriteString(paramName + "=" + strconv.FormatInt(v.Field(i).Elem().Int(), BIT_BASE_10) + "&")
			case reflect.Float32, reflect.Float64:
				paramBuffer.WriteString(paramName + "=" + decimal.NewFromFloat(v.Field(i).Elem().Float()).String() + "&")
			case reflect.Bool:
				paramBuffer.WriteString(paramName + "=" + strconv.FormatBool(v.Field(i).Elem().Bool()) + "&")
			case reflect.Struct:
				sv := reflect.ValueOf(v.Field(i).Interface())
				ToStringMethod := sv.MethodByName("String")
				params := make([]reflect.Value, 0)
				result := ToStringMethod.Call(params)
				paramBuffer.WriteString(paramName + "=" + url.QueryEscape(result[0].String()) + "&")
			case reflect.Slice:
				s := v.Field(i).Interface()
				d, _ := json.Marshal(s)
				paramBuffer.WriteString(paramName + "=" + url.QueryEscape(string(d)) + "&")
			case reflect.Invalid:
			default:
				log.Errorf("req type error %s:%s", paramName, v.Field(i).Elem().Kind())
			}
		}
	}

	return strings.TrimRight(paramBuffer.String(), "&")
}

func AsterGetRestHostByApiType(apiType ApiType) string {
	switch apiType {
	case SPOT:
		switch NowNetType {
		case MAIN_NET:
			return ASTER_API_SPOT_HTTP
		}
	case FUTURE:
		switch NowNetType {
		case MAIN_NET:
			return ASTER_API_FUTURE_HTTP
		}
	}
	return ""
}

func interfaceStringToFloat64(inter interface{}) float64 {
	return stringToFloat64(inter.(string))
}

func interfaceStringToInt64(inter interface{}) int64 {
	return int64(inter.(float64))
}

func stringToFloat64(str string) float64 {
	f, _ := strconv.ParseFloat(str, BIT_SIZE_64)
	return f
}

type MySyncMap[K any, V any] struct {
	smap sync.Map
}

func NewMySyncMap[K any, V any]() MySyncMap[K, V] {
	return MySyncMap[K, V]{
		smap: sync.Map{},
	}
}
func (m *MySyncMap[K, V]) Load(k K) (V, bool) {
	v, ok := m.smap.Load(k)

	if ok {
		return v.(V), true
	}
	var resv V
	return resv, false
}
func (m *MySyncMap[K, V]) Store(k K, v V) {
	m.smap.Store(k, v)
}

func (m *MySyncMap[K, V]) Delete(k K) {
	m.smap.Delete(k)
}
func (m *MySyncMap[K, V]) Range(f func(k K, v V) bool) {
	m.smap.Range(func(k, v any) bool {
		return f(k.(K), v.(V))
	})
}

func (m *MySyncMap[K, V]) Length() int {
	length := 0
	m.Range(func(k K, v V) bool {
		length += 1
		return true
	})
	return length
}

func (m *MySyncMap[K, V]) MapValues(f func(k K, v V) V) *MySyncMap[K, V] {
	var res = NewMySyncMap[K, V]()
	m.Range(func(k K, v V) bool {
		res.Store(k, f(k, v))
		return true
	})
	return &res
}
