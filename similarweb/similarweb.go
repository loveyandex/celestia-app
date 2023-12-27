package similarweb

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func Tw(dom string ) {


	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://data.similarweb.com/api/v1/data?domain="+dom, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("authority", "data.similarweb.com")
	req.Header.Set("accept", "*/*")
	req.Header.Set("accept-language", "en-US,en;q=0.9,fa;q=0.8")
	req.Header.Set("content-type", "application/json")
	// req.Header.Set("cookie", "sgID=5a737922-f364-6f27-648c-787b4fc66c24; sw_extension_installed=1698551648015; _gcl_au=1.1.1498424364.1698551651; _ga_3YZBGLLSHF=GS1.2.1698551651.1.0.1698551651.0.0.0; _ga=GA1.1.2027769837.1698551651; _ga_XGRRHHKH0P=GS1.1.1698551652.1.0.1698551652.60.0.0; _ga_5M34ZH84YZ=GS1.1.1698551652.1.0.1698551652.60.0.0; _pk_id.1.fd33=ebb0d3e374f99739.1698551653.; _abck=D08780C4DF1271134F4EA004A268933D~-1~YAAQPezEF8qOYUmLAQAAH2qReQp+ZW7qn8s38OYjuUBI6N4jdMamrrwo3U29GikygGvEWm7C0eJIZPGE2V6bta5G2drVoCVf5h7m81nzVbJ8cmt3zthk9AY1vVm9GsLelOIAUkwhlWYp9ShK74OUyf/uyvOO0rWVq4Jztn+qBvTV4GISlYdr0uYFDURmRYvh0Foi/AtT/vYdpvd8ogiUaH3ucFWfpL9WjpQk2bA7Qz5e7gnwfMTyoEJqcF0ZWCXuwJ+L/o94JmtE3+ZoLKsmTvsg0pnkBJ+IO2Bv1mqayAAy8gyWSmqHlZHQxUr4uCli7V/dVQ9t43jjjKj+FJ/rhzZyZUPLDDTgDcG7WHvyRMj9YKsZalIcI15CcP2+Zj7B~-1~-1~-1; cb_user_id=null; cb_group_id=null; cb_anonymous_id=%22f33624cb-b32b-45c8-862f-72151eb73417%22; __q_state_9u7uiM39FyWVMWQF=eyJ1dWlkIjoiNzVjYzc1OWEtOGFmNi00ZjU3LWJkMTgtMzdlZTYzYjZlYjkwIiwiY29va2llRG9tYWluIjoic2ltaWxhcndlYi5jb20iLCJtZXNzZW5nZXJFeHBhbmRlZCI6ZmFsc2UsInByb21wdERpc21pc3NlZCI6ZmFsc2UsImNvbnZlcnNhdGlvbklkIjoiMTI1NTczMjUzNDc0NzA4ODAwMyJ9")
	req.Header.Set("sec-fetch-dest", "empty")
	req.Header.Set("sec-fetch-mode", "cors")
	req.Header.Set("sec-fetch-site", "none")
	req.Header.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/119.0.0.0 Safari/537.36")
	req.Header.Set("x-extension-version", "6.8.2")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", bodyText)

	var v SimilarwebDmnResp

	err2 := json.Unmarshal(bodyText, &v)

	if err2 != nil {

		
		
	}

	fmt.Printf("v: %v\n", v)



}


type SimilarwebDmnResp struct {

	Version                int                    `json:"Version,omitempty"`
	SiteName               string                 `json:"SiteName,omitempty"`
	Description            string                 `json:"Description,omitempty"`
	TopCountryShares       []TopCountryShares     `json:"TopCountryShares,omitempty"`
	Title                  string                 `json:"Title,omitempty"`
	Engagments             Engagments             `json:"Engagments,omitempty"`
	EstimatedMonthlyVisits EstimatedMonthlyVisits `json:"EstimatedMonthlyVisits,omitempty"`
	GlobalRank             GlobalRank             `json:"GlobalRank,omitempty"`
	CountryRank            CountryRank            `json:"CountryRank,omitempty"`
	CategoryRank           CategoryRank           `json:"CategoryRank,omitempty"`
	GlobalCategoryRank     any                    `json:"GlobalCategoryRank,omitempty"`
	IsSmall                bool                   `json:"IsSmall,omitempty"`
	Policy                 int                    `json:"Policy,omitempty"`
	TrafficSources         TrafficSources         `json:"TrafficSources,omitempty"`
	Category               string                 `json:"Category,omitempty"`
	LargeScreenshot        string                 `json:"LargeScreenshot,omitempty"`
	IsDataFromGa           bool                   `json:"IsDataFromGa,omitempty"`
	Countries              []Countries            `json:"Countries,omitempty"`
	Competitors            Competitors            `json:"Competitors,omitempty"`
	Notification           Notification           `json:"Notification,omitempty"`
}
type TopCountryShares struct {
	Country     int     `json:"Country,omitempty"`
	CountryCode string  `json:"CountryCode,omitempty"`
	Value       float64 `json:"Value,omitempty"`
}
type Engagments struct {
	BounceRate   string `json:"BounceRate,omitempty"`
	Month        string `json:"Month,omitempty"`
	Year         string `json:"Year,omitempty"`
	PagePerVisit string `json:"PagePerVisit,omitempty"`
	Visits       string `json:"Visits,omitempty"`
	TimeOnSite   string `json:"TimeOnSite,omitempty"`
}
type EstimatedMonthlyVisits struct {
	Two0230801 int64 `json:"2023-08-01,omitempty"`
	Two0230901 int64 `json:"2023-09-01,omitempty"`
	Two0231001 int64 `json:"2023-10-01,omitempty"`
}
type GlobalRank struct {
	Rank int `json:"Rank,omitempty"`
}
type CountryRank struct {
	Country     int    `json:"Country,omitempty"`
	CountryCode string `json:"CountryCode,omitempty"`
	Rank        int    `json:"Rank,omitempty"`
}
type CategoryRank struct {
	Rank     string `json:"Rank,omitempty"`
	Category string `json:"Category,omitempty"`
}
type TrafficSources struct {
	Social        float64 `json:"Social,omitempty"`
	PaidReferrals float64 `json:"Paid Referrals,omitempty"`
	Mail          float64 `json:"Mail,omitempty"`
	Referrals     float64 `json:"Referrals,omitempty"`
	Search        float64 `json:"Search,omitempty"`
	Direct        float64 `json:"Direct,omitempty"`
}
type Countries struct {
	Code    string `json:"Code,omitempty"`
	URLCode string `json:"UrlCode,omitempty"`
	Name    string `json:"Name,omitempty"`
}
type Competitors struct {
	TopSimilarityCompetitors []any `json:"TopSimilarityCompetitors,omitempty"`
}
type Notification struct {
	Content any `json:"Content,omitempty"`
}