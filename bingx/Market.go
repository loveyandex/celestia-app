package bingx

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func XPrice(symbol string) (float64, error) {

	dataStr := `{
		"uri": "/openApi/swap/v2/quote/price",
		"method": "GET",
		"protocol": "https"
	}`
	payload := fmt.Sprintf(`{
		"symbol": "%s"
	}`, symbol)
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
		return 0, err
	}
	req.Header.Add("X-BX-APIKEY", API_KEY)
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return 0, err

	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	fmt.Println(string(body))

	var v ApiResp[SymbolPrice]

	err2 := json.Unmarshal(body, &v)

	if err2 != nil {
		return -55, err
	}

	f, err2 := strconv.ParseFloat(v.Data.Price, 64)

	if err2 != nil {
		return -55, err
	}

	return f, nil

}

type SymbolPrice struct {
	Symbol string `json:"symbol,omitempty"`
	Price  string `json:"price,omitempty"`
	Time   int64  `json:"time,omitempty"`
}
