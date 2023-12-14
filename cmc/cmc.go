package cmc

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

func getNextData(url string) (string, error) {
	// Make the HTTP request to the Next.js page
	response, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	// Load the HTML response into a GoQuery document
	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		return "", err
	}

	// Find the script tag containing the __NEXT_DATA__ variable
	script := doc.Find("script#__NEXT_DATA__").First()

	// Extract the value of the __NEXT_DATA__ variable
	nextData := strings.TrimSpace(script.Text())

	return nextData, nil
}

func CmcNew() {

	var rs map[int]*RecentlyAddedList = make(map[int]*RecentlyAddedList)

	for {

		// Next.js page URL
		url := "https://coinmarketcap.com/new/"
		// Crawl the Next.js page and get the __NEXT_DATA__ variable value
		nextData, err := getNextData(url)
		if err != nil {
			log.Fatal(err)
		}
		var v NewCoinsCmcResp
		json.Unmarshal([]byte(nextData), &v)

		for _, ral := range v.Props.PageProps.Data.Data.RecentlyAddedList {
			fmt.Printf("ral.PriceChange.Volume24H: %v\n", ral.PriceChange.Volume24H)
			if ral.PriceChange.Volume24H > 0.5e6 {
				if rs[ral.ID] == nil {
					rs[ral.ID] = &ral
					//send this new coin to tel
					println(ral.Name)
				}
			}

		}

		time.Sleep(12 * time.Hour)
	}
}

type NewCoinsCmcResp struct {
	Props Props `json:"props"`
}

type Empty struct {
}

type PriceChange struct {
	Price          float64   `json:"price"`
	PriceChange1H  float64   `json:"priceChange1h"`
	PriceChange24H float64   `json:"priceChange24h"`
	PriceChange7D  float64   `json:"priceChange7d"`
	PriceChange30D float64   `json:"priceChange30d"`
	Volume24H      float64   `json:"volume24h"`
	LastUpdate     time.Time `json:"lastUpdate"`
}
type Platforms struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Symbol string `json:"symbol"`
	Slug   string `json:"slug"`
}
type RecentlyAddedList struct {
	ID                    int         `json:"id"`
	Name                  string      `json:"name"`
	Symbol                string      `json:"symbol"`
	Slug                  string      `json:"slug"`
	Rank                  int         `json:"rank"`
	PriceChange           PriceChange `json:"priceChange"`
	Platforms             []Platforms `json:"platforms"`
	AddedDate             time.Time   `json:"addedDate"`
	MarketCap             float64     `json:"marketCap"`
	SelfReportedMarketCap float64     `json:"selfReportedMarketCap"`
	FullyDilutedMarketCap float64     `json:"fullyDilutedMarketCap"`
}
type Data00 struct {
	RecentlyAddedList []RecentlyAddedList `json:"recentlyAddedList"`
}
type Status struct {
	Timestamp    time.Time `json:"timestamp"`
	ErrorCode    string    `json:"error_code"`
	ErrorMessage string    `json:"error_message"`
	Elapsed      string    `json:"elapsed"`
	CreditCount  int       `json:"credit_count"`
}
type Data0 struct {
	Data   Data00 `json:"data"`
	Status Status `json:"status"`
}
type EtherscanGas struct {
	LastBlock                string `json:"lastBlock"`
	SlowPrice                string `json:"slowPrice"`
	SlowConfirmationTime     string `json:"slowConfirmationTime"`
	StandardPrice            string `json:"standardPrice"`
	StandardConfirmationTime string `json:"standardConfirmationTime"`
	FastPrice                string `json:"fastPrice"`
	FastConfirmationTime     string `json:"fastConfirmationTime"`
}
type GlobalMetrics struct {
	NumCryptocurrencies int          `json:"numCryptocurrencies"`
	NumMarkets          int          `json:"numMarkets"`
	ActiveExchanges     int          `json:"activeExchanges"`
	MarketCap           float64      `json:"marketCap"`
	MarketCapChange     float64      `json:"marketCapChange"`
	TotalVol            float64      `json:"totalVol"`
	StablecoinVol       float64      `json:"stablecoinVol"`
	StablecoinChange    float64      `json:"stablecoinChange"`
	TotalVolChange      float64      `json:"totalVolChange"`
	DefiVol             float64      `json:"defiVol"`
	DefiChange          float64      `json:"defiChange"`
	DefiMarketCap       float64      `json:"defiMarketCap"`
	DerivativesVol      float64      `json:"derivativesVol"`
	DerivativeChange    float64      `json:"derivativeChange"`
	BtcDominance        float64      `json:"btcDominance"`
	BtcDominanceChange  float64      `json:"btcDominanceChange"`
	EthDominance        float64      `json:"ethDominance"`
	EtherscanGas        EtherscanGas `json:"etherscanGas"`
}
type TopCategories struct {
	Title          string `json:"title"`
	RelatedTagSlug string `json:"relatedTagSlug"`
}
type CurrentIndex struct {
	Score      float64   `json:"score"`
	MaxScore   int       `json:"maxScore"`
	Name       string    `json:"name"`
	UpdateTime time.Time `json:"updateTime"`
}
type DialConfig struct {
	Start int    `json:"start"`
	End   int    `json:"end"`
	Name  string `json:"name"`
}
type FearGreedIndexData struct {
	CurrentIndex CurrentIndex `json:"currentIndex"`
	DialConfig   []DialConfig `json:"dialConfig"`
}
type FaqList struct {
	Question string `json:"question"`
	Answer   string `json:"answer"`
	Sequence int    `json:"sequence"`
}
type FaqData struct {
	Language string    `json:"language"`
	FaqList  []FaqList `json:"faqList"`
}
type HalvingInfo struct {
	CurrentBlockHeight    int   `json:"currentBlockHeight"`
	BlocksUntilHaving     int   `json:"blocksUntilHaving"`
	SecondsUntilHalving   int   `json:"secondsUntilHalving"`
	SecondsUntilNextBlock int   `json:"secondsUntilNextBlock"`
	ApproximateBlockTime  int   `json:"approximateBlockTime"`
	HalvingTimeStamp      int64 `json:"halvingTimeStamp"`
}
type Banners struct {
	ID               string   `json:"id"`
	Name             string   `json:"name"`
	Link             string   `json:"link"`
	Image            string   `json:"image"`
	IncludeCountries []string `json:"includeCountries"`
	Location         []string `json:"location"`
	Platform         []string `json:"platform"`
	ImageAlignment   string   `json:"imageAlignment"`
	CustomOptions    string   `json:"customOptions"`
}
type AdBanners struct {
	Location string    `json:"location"`
	Code     int       `json:"code"`
	Banners  []Banners `json:"banners"`
}
type AdsContent struct {
	AdBanners []AdBanners `json:"adBanners"`
}
type DeviceInfo struct {
	IsDesktop bool `json:"isDesktop"`
	IsTablet  bool `json:"isTablet"`
	IsMobile  bool `json:"isMobile"`
}
type PageSharedData struct {
	TopCategories      []TopCategories    `json:"topCategories"`
	FearGreedIndexData FearGreedIndexData `json:"fearGreedIndexData"`
	CountryCode        string             `json:"countryCode"`
	FaqData            FaqData            `json:"faqData"`
	HalvingInfo        HalvingInfo        `json:"halvingInfo"`
	AdsContent         AdsContent         `json:"adsContent"`
	DeviceInfo         DeviceInfo         `json:"deviceInfo"`
}
type PageProps struct {
	Data               Data0          `json:"data"`
	ReqLang            string         `json:"reqLang"`
	GlobalMetrics      GlobalMetrics  `json:"globalMetrics"`
	PageSharedData     PageSharedData `json:"pageSharedData"`
	PageSize           int            `json:"pageSize"`
	NamespacesRequired []string       `json:"namespacesRequired"`
	Noindex            bool           `json:"noindex"`
}

type Props struct {
	PageProps PageProps `json:"pageProps"`
}
