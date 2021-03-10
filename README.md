# Weather web scraper

This is a [forecast.weather.gov](https://www.weather.gov/) web scraper written in Go. Using [Colly](http://go-colly.org/).

### Features

- Endpoint available at localhost:3000. Returns scraped data for today and the next 7 days.
- Configurable environment variables
  - url = https://forecast.weather.gov/MapClick.php?lat=35.76148000000006&lon=-77.94274999999999 <br/>
  To set your own location, go [here](https://www.weather.gov/) and search for your zip code. 

  - port = 3000
  Changes the internal API port. For example, if you pass 3000 then the API will be available at
  ```http://localhost:3000?q=https://forecast.weather.gov/MapClick.php?lat=35.76148000000006&lon=-77.94274999999999```

  - endpoint = http://localhost:8080 <br/>
  If you pass any value here, it will attempt to scrape and send the data to this location at the interval specified below. API endpoint at the specified port will still be available. Not required. 

  - cron = * * * * * <br/>
  Any valid cron interval. Only needed if using the above endpoint var. Not required. <br/><br/>

Defaults are provided in the Dockerfile. 


