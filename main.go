package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/fatih/camelcase"
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

// home route logic
func home(w http.ResponseWriter, r *http.Request) {

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

		// short description. cleaned up
		// get text
		desc := e.ChildText(".short-desc")
		// split on camel case (didn't bring over linebreaks)
		// and delete any empty slices
		descSplit := delete_empty(camelcase.Split(desc))
		// join slice with space between
		descJoin := strings.Join(descSplit, " ")
		// add to tempday
		tempDay.Condition = descJoin

		// get image
		tempDay.Image = (url + "/" + e.ChildAttr("img", "src"))

		// get temperature in F
		tempDay.Temp = e.ChildText(".temp")

		// append to days
		res.Days = append(res.Days, tempDay)
	})

	// visit page
	c.Visit("https://forecast.weather.gov/MapClick.php?lat=35.76148000000006&lon=-77.94274999999999")

	// wait. visit all pages first
	c.Wait()

	// convert res to json
	json, _ := json.Marshal(res)

	// send response
	fmt.Fprintf(w, string(json))
	fmt.Println("Endpoint Hit: home")
}

func handleRequests() {
	// set up routes
	http.HandleFunc("/", home)
	// log, listening on port 10000
	fmt.Println("listening on port 10000")
	// listen on port 10000
	log.Fatal(http.ListenAndServe(":10000", nil))

}

// run main
func main() {
	handleRequests()
}

func delete_empty(s []string) []string {
	var r []string
	for _, str := range s {
		if str != " " {
			r = append(r, str)
		}
	}
	return r
}
