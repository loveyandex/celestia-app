package bingx

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const HOST = "open-api-vst.bingx.com"

// const HOST = "open-api-vst.bingx.com"
const API_KEY = "RyRJT93zjlT2f2NjelJZWfHjgKQZycyL4KtSqu3y83ncYOucPp0UeDGIZwunUAfI9AxqdSVYTUnIGg55RWCw"
const API_SECRET = "vHADQEvRkRfyV42p3mQ9SInLcrWBtRB4as8UxZzZzhSizoZ37NwKaVZmGOiXqo7pdZ8S4AFLbuixNfPzYt2Q"

type APIResp struct {
	Code int         `json:"code,omitempty"`
	Msg  string      `json:"msg,omitempty"`
	Data []CtrctData `json:"data,omitempty"`
}
type CtrctData struct {
	ContractID        string  `json:"contractId,omitempty"`
	Symbol            string  `json:"symbol,omitempty"`
	Size              string  `json:"size,omitempty"`
	QuantityPrecision int     `json:"quantityPrecision,omitempty"`
	PricePrecision    int     `json:"pricePrecision,omitempty"`
	FeeRate           float64 `json:"feeRate,omitempty"`
	TradeMinLimit     int     `json:"tradeMinLimit,omitempty"`
	MaxLongLeverage   int     `json:"maxLongLeverage,omitempty"`
	MaxShortLeverage  int     `json:"maxShortLeverage,omitempty"`
	Currency          string  `json:"currency,omitempty"`
	Asset             string  `json:"asset,omitempty"`
	Status            int     `json:"status,omitempty"`
}

type ApiResp[X interface{}] struct {
	Code int    `json:"code,omitempty"`
	Msg  string `json:"msg,omitempty"`
	Data X      `json:"data,omitempty"`
}

type ApiRespBal struct {
	Code int         `json:"code,omitempty"`
	Msg  string      `json:"msg,omitempty"`
	Data DataBalance `json:"data,omitempty"`
}
type Balance struct {
	UserID           string `json:"userId,omitempty"`
	Asset            string `json:"asset,omitempty"`
	Balance          string `json:"balance,omitempty"`
	Equity           string `json:"equity,omitempty"`
	UnrealizedProfit string `json:"unrealizedProfit,omitempty"`
	RealisedProfit   string `json:"realisedProfit,omitempty"`
	AvailableMargin  string `json:"availableMargin,omitempty"`
	UsedMargin       string `json:"usedMargin,omitempty"`
	FreezedMargin    string `json:"freezedMargin,omitempty"`
}
type DataBalance struct {
	Balance Balance `json:"balance,omitempty"`
}

func Cntrcts() (*[]CtrctData, error) {

	dataStr := `{
    "uri": "/openApi/swap/v2/quote/contracts",
    "method": "GET",
    "protocol": "https"
}`
	payload := `{}`
	TIMESTAMP := time.Now().UnixNano() / 1e6
	apiMap := getParameters(dataStr, payload, false, TIMESTAMP)
	sign := computeHmac256(fmt.Sprintf("%v", apiMap["parameters"]), API_SECRET)
	fmt.Println("parameters:", fmt.Sprintf("%v", apiMap["parameters"]))
	fmt.Println("sign:", sign)
	parameters := ""
	contains := strings.ContainsAny(fmt.Sprintf("%v", apiMap["parameters"]), "[{")
	if contains {
		apiMap2 := getParameters(dataStr, payload, true, TIMESTAMP)
		parameters = fmt.Sprintf("%v&signature=%s", apiMap2["parameters"], sign)
	} else {
		parameters = fmt.Sprintf("%v&signature=%s", apiMap["parameters"], sign)
	}
	url := fmt.Sprintf("%v://%s%v?%s", apiMap["protocol"], HOST, apiMap["uri"], parameters)
	method := fmt.Sprintf("%v", apiMap["method"])
	client := &http.Client{}
	fmt.Println("url:", url)
	fmt.Println("method:", method)
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return nil, err

	}
	req.Header.Add("X-BX-APIKEY", API_KEY)
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil, err

	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	fmt.Println(string(body))

	var c ApiResp[[]CtrctData]
	json.Unmarshal(body, &c)

	for _, cd := range c.Data {

		fmt.Printf("cd.Asset: %v\n", cd.Asset)
	}

	if c.Code == 0 {
		return &c.Data, err

	} else {
		return nil, errors.New(c.Msg)
	}

}

func GtBalance() (*Balance, error) {

	dataStr := `{
	  "uri": "/openApi/swap/v2/user/balance",
	  "method": "GET",
	  "protocol": "https"
  }`
	payload := `{
	  "recvWindow": 0
  }`
	TIMESTAMP := time.Now().UnixNano() / 1e6
	apiMap := getParameters(dataStr, payload, false, TIMESTAMP)
	sign := computeHmac256(fmt.Sprintf("%v", apiMap["parameters"]), API_SECRET)
	fmt.Println("parameters:", fmt.Sprintf("%v", apiMap["parameters"]))
	fmt.Println("sign:", sign)
	parameters := ""
	contains := strings.ContainsAny(fmt.Sprintf("%v", apiMap["parameters"]), "[{")
	if contains {
		apiMap2 := getParameters(dataStr, payload, true, TIMESTAMP)
		parameters = fmt.Sprintf("%v&signature=%s", apiMap2["parameters"], sign)
	} else {
		parameters = fmt.Sprintf("%v&signature=%s", apiMap["parameters"], sign)
	}
	url := fmt.Sprintf("%v://%s%v?%s", apiMap["protocol"], HOST, apiMap["uri"], parameters)
	method := fmt.Sprintf("%v", apiMap["method"])
	client := &http.Client{}
	fmt.Println("url:", url)
	fmt.Println("method:", method)
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	req.Header.Add("X-BX-APIKEY", API_KEY)
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil, err

	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return nil, err

	}
	fmt.Println(string(body))

	var b ApiResp[DataBalance]

	err = json.Unmarshal(body, &b)

	if err != nil {

		fmt.Printf("err: %v\n", err)

	}

	fmt.Printf("b: %v\n", b)

	return &b.Data.Balance, nil

}

func SetLeverage(symbol, side string, l int) {

	dataStr := `{
		"uri": "/openApi/swap/v2/trade/leverage",
		"method": "POST",
		"protocol": "https"
	}`
	payload := fmt.Sprintf(`{
		"symbol": "%s",
		"side": "%s",
		"leverage": %d,
		"recvWindow": 0
	}`, symbol, side, l)

	TIMESTAMP := time.Now().UnixNano() / 1e6
	apiMap := getParameters(dataStr, payload, false, TIMESTAMP)
	sign := computeHmac256(fmt.Sprintf("%v", apiMap["parameters"]), API_SECRET)
	fmt.Println("parameters:", fmt.Sprintf("%v", apiMap["parameters"]))
	fmt.Println("sign:", sign)
	parameters := ""
	contains := strings.ContainsAny(fmt.Sprintf("%v", apiMap["parameters"]), "[{")
	if contains {
		apiMap2 := getParameters(dataStr, payload, true, TIMESTAMP)
		parameters = fmt.Sprintf("%v&signature=%s", apiMap2["parameters"], sign)
	} else {
		parameters = fmt.Sprintf("%v&signature=%s", apiMap["parameters"], sign)
	}
	url := fmt.Sprintf("%v://%s%v?%s", apiMap["protocol"], HOST, apiMap["uri"], parameters)
	method := fmt.Sprintf("%v", apiMap["method"])
	client := &http.Client{}
	fmt.Println("url:", url)
	fmt.Println("method:", method)
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("X-BX-APIKEY", API_KEY)
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))

}

type Order struct {
	Symbol        string `json:"symbol,omitempty"`
	OrderID       int64  `json:"orderId,omitempty"`
	Side          string `json:"side,omitempty"`
	PositionSide  string `json:"positionSide,omitempty"`
	Type          string `json:"type,omitempty"`
	ClientOrderID string `json:"clientOrderID,omitempty"`
	WorkingType   string `json:"workingType,omitempty"`
}
type DataOrder struct {
	Order Order `json:"order,omitempty"`
}

func TakeMrktOrder(symbol string, amount int64) {

	dataStr := `{
		"uri": "/openApi/swap/v2/trade/order",
		"method": "POST",
		"protocol": "https"
	}`
	payload := fmt.Sprintf(`{
		"symbol": "%s",
		"side": "BUY",
		"positionSide": "LONG",
		"type": "MARKET",
		"quantity": %d
	}`, symbol, amount)

	TIMESTAMP := time.Now().UnixNano() / 1e6
	apiMap := getParameters2(dataStr, payload, false, TIMESTAMP)
	sign := computeHmac2562(fmt.Sprintf("%v", apiMap["parameters"]), API_SECRET)
	fmt.Println("parameters:", fmt.Sprintf("%v", apiMap["parameters"]))
	fmt.Println("sign:", sign)
	parameters := ""
	contains := strings.ContainsAny(fmt.Sprintf("%v", apiMap["parameters"]), "[{")
	if contains {
		apiMap2 := getParameters2(dataStr, payload, true, TIMESTAMP)
		parameters = fmt.Sprintf("%v&signature=%s", apiMap2["parameters"], sign)
	} else {
		parameters = fmt.Sprintf("%v&signature=%s", apiMap["parameters"], sign)
	}
	url := fmt.Sprintf("%v://%s%v?%s", apiMap["protocol"], HOST, apiMap["uri"], parameters)
	method := fmt.Sprintf("%v", apiMap["method"])
	client := &http.Client{}
	fmt.Println("url:", url)
	fmt.Println("method:", method)
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("X-BX-APIKEY", API_KEY)
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(body))
	var v ApiResp[DataOrder]

	json.Unmarshal(body, &v)

	if v.Code == 0 {

		fmt.Printf("v.Data.Order: %v\n", v.Data.Order)

	}
}

func StateOfOrder(orderid int64) error {

	//4486585016063211929

	return nil
}

func computeHmac256(strMessage string, strSecret string) string {
	key := []byte(strSecret)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(strMessage))
	return hex.EncodeToString(h.Sum(nil))
}
func getParameters(dataStr string, payload string, urlEncode bool, timestemp int64) map[string]interface{} {

	var apiMap map[string]interface{}
	var payloadMap map[string]interface{}
	err := json.Unmarshal([]byte(dataStr), &apiMap)
	if err != nil {
		fmt.Printf("json to map error,err:%s", err)
		return apiMap
	}
	err = json.Unmarshal([]byte(payload), &payloadMap)
	if err != nil {
		fmt.Printf("json to map error,err:%s", err)
		return apiMap
	}
	parameters := ""
	for key, value := range payloadMap {
		if urlEncode {
			encodedStr := url.QueryEscape(fmt.Sprintf("%v", value))
			encodedStr = strings.ReplaceAll(encodedStr, "+", "%20")
			parameters = parameters + key + "=" + encodedStr + "&"
		} else {
			parameters = parameters + key + "=" + fmt.Sprintf("%v", value) + "&"
		}
	}
	parameters += "timestamp=" + fmt.Sprintf("%d", timestemp)
	apiMap["parameters"] = fmt.Sprintf("%v", parameters)
	return apiMap
}
