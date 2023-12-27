package bingx

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

func XorderSts(oid int64) {

	dataStr := `{
		"uri": "/openApi/swap/v2/trade/order",
		"method": "GET",
		"protocol": "https"
	}`
	payload := fmt.Sprintf(`{
		"symbol": "1000BONK-USDT",
		"orderId": "%d"
	}`,oid)
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



func XallOrder(symbol string,lim uint )  {

	
	dataStr := `{
		"uri": "/openApi/swap/v2/trade/allOrders",
		"method": "GET",
		"protocol": "https"
	}`
	  payload := fmt.Sprintf(`{
		"symbol": "%s",
		"limit": %d
	}`,symbol,lim)
	  TIMESTAMP := time.Now().UnixNano()/1e6
	  apiMap := getParameters(dataStr, payload,false,TIMESTAMP)
	  sign := computeHmac256(fmt.Sprintf("%v",apiMap["parameters"]), API_SECRET)
	  fmt.Println("parameters:",fmt.Sprintf("%v",apiMap["parameters"]))
	  fmt.Println("sign:",sign)
	  parameters := ""
	  contains := strings.ContainsAny( fmt.Sprintf("%v",apiMap["parameters"]), "[{")
	  if contains {
		apiMap2 := getParameters(dataStr, payload,true,TIMESTAMP)
		parameters = fmt.Sprintf("%v&signature=%s",apiMap2["parameters"], sign)
	  } else {
		parameters = fmt.Sprintf("%v&signature=%s",apiMap["parameters"], sign)
	  }
	  url := fmt.Sprintf("%v://%s%v?%s",apiMap["protocol"],HOST,apiMap["uri"],parameters)
	  method := fmt.Sprintf("%v",apiMap["method"])
	  client := &http.Client {}
	  fmt.Println("url:",url)
	  fmt.Println("method:",method)
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

func XcloseAllPositions()  {
	dataStr := `{
		"uri": "/openApi/swap/v2/trade/closeAllPositions",
		"method": "POST",
		"protocol": "https"
	}`
	  payload := `{
		"recvWindow": 0
	}`
	  TIMESTAMP := time.Now().UnixNano()/1e6
	  apiMap := getParameters(dataStr, payload,false,TIMESTAMP)
	  sign := computeHmac256(fmt.Sprintf("%v",apiMap["parameters"]), API_SECRET)
	  fmt.Println("parameters:",fmt.Sprintf("%v",apiMap["parameters"]))
	  fmt.Println("sign:",sign)
	  parameters := ""
	  contains := strings.ContainsAny( fmt.Sprintf("%v",apiMap["parameters"]), "[{")
	  if contains {
		apiMap2 := getParameters(dataStr, payload,true,TIMESTAMP)
		parameters = fmt.Sprintf("%v&signature=%s",apiMap2["parameters"], sign)
	  } else {
		parameters = fmt.Sprintf("%v&signature=%s",apiMap["parameters"], sign)
	  }
	  url := fmt.Sprintf("%v://%s%v?%s",apiMap["protocol"],HOST,apiMap["uri"],parameters)
	  method := fmt.Sprintf("%v",apiMap["method"])
	  client := &http.Client {}
	  fmt.Println("url:",url)
	  fmt.Println("method:",method)
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