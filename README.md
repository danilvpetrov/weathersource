![CI](https://github.com/danilvpetrov/weathersource/workflows/CI/badge.svg)

# Weathersource

Weathersource is a weather forecast data collection service. It uses [Dark Sky API](https://darksky.net/dev/docs) to upload the weather forecast data into the
cache storage and exposes web UI to read and analyze the uploaded data.

One of the main purposes of this service is to test the efficiency of the web
assets (files with extensions .html, .js, .css, and so on) compiled as binary
data into a Golang execution file.

The service consumes the following environment variables:

- **WEATHER_DARKSKY_TOKEN** - a security token required to download forecast
  data from [Dark Sky API](https://darksky.net/dev/docs). Must be specified.

- **WEATHER_LOC_LATITUDE** - The latitude of the geographical location of the
  weather forecast. This variable is mandatory and must be specified as
  decimal fraction, e.g. `-153.5689`.

- **WEATHER_LOC_LONGITUDE** - The longitude of the geographical location of the
  weather forecast. This variable is mandatory and must be specified as
  decimal fraction, e.g. `24.84348`.

- **WEATHER_UPDATE_INTERVAL** - The time interval to download forecast data from
  [Dark Sky API](https://darksky.net/dev/docs). Must be in the format that
  [`time.ParseDuration()`](https://golang.org/pkg/time/#ParseDuration) function
  understands. If not specified, it defaults to 5 minutes.

- **WEATHER_DB_DSN** - The DSN of the MySQL database to persist (cache) all
  weather forecast data. Must be specified.

- **WEATHER_GRPC_PORT** - The port number to expose gRPC service on. If not
  specified, it defaults to `8081`.

- **WEATHER_HTTP_PORT** - The port number to expose HTTP server on (Web UI/API).
  If not specified, it defaults to `8080`.




