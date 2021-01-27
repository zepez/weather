package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/robfig/cron/v3"
)

// home route logic
func home(w http.ResponseWriter, r *http.Request) {

	// getdata. pass in link
	response := GetData(r.FormValue("q"))
	// convert res to json
	resMarshalled, _ := json.Marshal(response)
	json := string(resMarshalled)

	// send response
	fmt.Fprintf(w, string(json))
	fmt.Println("Endpoint Hit: home")
}

// set up routes
func handleRequests() {
	http.HandleFunc("/", home)
	// log, listening on specified port
	fmt.Println("listening on port " + os.Getenv("port"))
	log.Fatal(http.ListenAndServe(":"+os.Getenv("port"), nil))
}

func main() {
	// set env variables for dev. uncomment below to set manually
	// os.Setenv("url", "https://forecast.weather.gov/MapClick.php?lat=35.722&lon=-77.9164")
	// os.Setenv("port", "3001")
	// os.Setenv("cron", "* * * * *")
	// os.Setenv("endpoint", "http://localhost:3000/v1/weather")

	if len(os.Getenv("cron")) > 0 {
		cj := cron.New()
		cj.AddFunc(os.Getenv("cron"), func() {
			// getdata. pass in link
			response := GetData(os.Getenv("url"))
			jsonMarsh, _ := json.Marshal(response)
			buf := bytes.NewBuffer(jsonMarsh)

			fmt.Println(response)

			// send response
			// create new reponse, get req and err
			req, err := http.NewRequest("POST", os.Getenv("endpoint"), buf)

			// set headers here
			req.Header.Set("Content-Type", "application/json")

			// error handling
			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				panic(err)
			}
			defer resp.Body.Close()

			// uncomment below for debug
			// fmt.Println("response Status:", resp.Status)
			// fmt.Println("response Headers:", resp.Header)
			// body, _ := ioutil.ReadAll(resp.Body)
			// fmt.Println("response Body:", string(body))
		})

		cj.Start()
	}

	handleRequests()
}
