# Weather web scraper

This is a web scraper written in Go. Currently scrapes [forecast.weather.gov](https://www.weather.gov/) for Wilson, North Carolina only. Using [Colly](http://go-colly.org/).

### Features

- Endpoint available at localhost:3000. Returns scraped data for today and the next 7 days.
- Configurable via environment variables
  - url = https://forecast.weather.gov/MapClick.php?lat=35.76148000000006&lon=-77.94274999999999
  - port = 3000

To customize the endpoint go [here](https://www.weather.gov/) and search for your zip code. 

