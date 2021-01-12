### Weather web scraper

This is a web scraper written in Go. Currently scrapes [forecast.weather.gov](https://www.weather.gov/) for Wilson, North Carolina only. Using [Colly](http://go-colly.org/).

This is a work in progress. To do: 

1. Currently stops scraping when it runs into a <br> tag. 
2. Automate via a cronjob. 
3. Docker container.
4. Configurable via environment variables. 
  - API endpoint
  - Cronjob times
  - By zip code? 

