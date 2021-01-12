package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/gocolly/colly/v2"
)

type total struct {
	Image     string
	Condition string
	F         string
	C         string
	CrawledAt time.Time
	Days      []day
}

type day struct {
	Period    string
	Condition string
	Image     string
	Temp      string
}

func main() {
	// define domain
	url := "forecast.weather.gov"
	// create res
	res := total{}

	// start colly
	c := colly.NewCollector(
		// Visit only domains: forecast.weather.gov
		colly.AllowedDomains(url),
	)

	// find current conditions by id
	c.OnHTML("#current_conditions-summary", func(e *colly.HTMLElement) {

		// get image url
		res.Image = (url + "/" + e.ChildAttr("img", "src"))
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

		// period of time
		tempDay.Period = e.ChildText(".period-name")

		// short description
		tempDay.Condition = e.ChildText(".short-desc")

		// get image
		tempDay.Image = (url + "/" + e.ChildAttr("img", "src"))

		// get temperature in F
		tempDay.Temp = e.ChildText(".temp")

		// append to days
		res.Days = append(res.Days, tempDay)
	})

	// visit page
	c.Visit("https://forecast.weather.gov/MapClick.php?lat=35.76148000000006&lon=-77.94274999999999")

	c.Wait()

	json, _ := json.Marshal(res)
	fmt.Println(string(json))
}
