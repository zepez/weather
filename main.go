package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

// home route logic
func home(w http.ResponseWriter, r *http.Request) {

	// getdata. pass in link
	json := GetData(r.FormValue("q"))

	// send response
	fmt.Fprintf(w, string(json))
	fmt.Println("Endpoint Hit: home")
}

// set up routes
func handleRequests() {
	http.HandleFunc("/", home)
	// log, listening on port 10000
	fmt.Println("listening on port " + os.Getenv("port"))
	// listen on port 10000
	log.Fatal(http.ListenAndServe(":"+os.Getenv("port"), nil))
}

func main() {
	// set env variables for dev
	os.Setenv("url", "https://forecast.weather.gov/MapClick.php?lat=35.76148000000006&lon=-77.94274999999999")
	os.Setenv("port", "3000")
	handleRequests()
}
