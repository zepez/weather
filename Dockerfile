FROM golang:1.12 as builder

# Set Environment Variables
ENV port 3000
ENV url https://forecast.weather.gov/MapClick.php?lat=35.76148000000006&lon=-77.94274999999999
ENV cron * * * * *
ENV endpoint http://localhost:3001/rail/test

WORKDIR /app
COPY . .

RUN go mod download

# Build app
RUN go build -o main .

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /app

# Copy the pre-built binary file from the previous stage
COPY --from=builder /app/main .

EXPOSE 3000

CMD [ "main" ]
