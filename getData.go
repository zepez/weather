package main

import (
	"strings"
	"time"

	"github.com/fatih/camelcase"
	"github.com/gocolly/colly"
)

type total struct {
	Image     string    `json:"image"`
	Condition string    `json:"condition"`
	F         string    `json:"f"`
	C         string    `json:"c"`
	CrawledAt time.Time `json:"crawled_at"`
	Days      []day     `json:"days"`
}

type day struct {
	Period    string `json:"period"`
	Condition string `json:"condition"`
	Image     string `json:"image"`
	Temp      string `json:"temp"`
}

// home route logic
func GetData(url string) interface{} {

	// define domain
	domain := "forecast.weather.gov"
	// create res
	res := total{}

	// start colly
	c := colly.NewCollector(
		// Visit only domain: forecast.weather.gov
		colly.AllowedDomains(domain),
	)

	// find current conditions by id
	c.OnHTML("#current_conditions-summary", func(e *colly.HTMLElement) {

		// get image url
		res.Image = (domain + "/" + e.ChildAttr("img", "src"))
		// short text
		res.Condition = e.ChildText(".myforecast-current")
		// temp in far
		res.F = e.ChildText(".myforecast-current-lrg")
		// temp in cel
		res.C = e.ChildText(".myforecast-current-sm")

	})

	// find current conditions by id
	c.OnHTML(".forecast-tombstone", func(e *colly.HTMLElement) {
		// create temp day var
		tempDay := day{}

		// period of time. cleaned up
		tempDay.Period = splitString(e.ChildText(".period-name"))

		// short description. cleaned up
		tempDay.Condition = splitString(e.ChildText(".short-desc"))

		// get image
		tempDay.Image = (domain + "/" + e.ChildAttr("img", "src"))

		// get temperature in F
		tempDay.Temp = e.ChildText(".temp")

		// append to days
		res.Days = append(res.Days, tempDay)
	})

	// visit page
	c.Visit(url)

	// wait. visit all pages first
	c.Wait()

	// return value
	return res
}

func splitString(s string) string {
	// split string via camelcase into slice
	arr := camelcase.Split(s)
	var cleanedArr []string

	// clean slice of empties
	for _, str := range arr {
		if str != " " {
			cleanedArr = append(cleanedArr, str)
		}
	}

	// join slice with space
	joined := strings.Join(cleanedArr, " ")
	return joined
}
