FROM golang:alpine

# important!
ENV GO111MODULE=on
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
ENV GOFLAGS=-mod=vendor
ENV APP_USER app
ENV APP_HOME /go/src/microservices

# custom 
ENV port 3000
ENV url https://forecast.weather.gov/MapClick.php?lat=35.76148000000006&lon=-77.94274999999999
ENV cron * * * * *
ENV endpoint http://localhost:3000

RUN mkdir /app
ADD . /app
WORKDIR /app

# compile your project
RUN go mod vendor
RUN go build

# open the port 8000
EXPOSE 3000
CMD [ "/app/weather" ]










