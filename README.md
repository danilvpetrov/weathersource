# weathersource

Weathersource is a weather forecast data collection service. It uses [Dark Sky
API](https://darksky.net/dev/docs) to upload the weather forecast data into the cache storage and exposes web
UI to read and analyze the uploaded data.

One of the main purposes of this service is to test the efficiency of the web
assets (files with extensions .html, .js, .css, and so on) compiled as binary
data into a Golang execution file.

The service consumes the following environment variables:

- **WEATHER_DARKSKY_TOKEN** - a security token required to download forecast data from [Dark Sky
API](https://darksky.net/dev/docs).

- **WEATHER_LOC_LATITUDE** - The latitude of the geographical location of the weather forecast. Must be in the float point format, e.g. `-153.5689`.

- **WEATHER_LOC_LONGITUDE** - The longitude of the geographical location of the weather forecast. Must be in the floating point format, e.g. `24.84348`.



