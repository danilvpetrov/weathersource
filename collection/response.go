package collection

import (
	"github.com/danilvpetrov/weathersource/data"
)

// Datum is a unit of weather information queried from https://api.darksky.net.
type Datum struct {
	ApparentTemperatureHigh     float64 `json:"apparentTemperatureHigh,omitempty"`
	ApparentTemperatureHighTime int64   `json:"apparentTemperatureHighTime,omitempty"`
	ApparentTemperatureLow      float64 `json:"apparentTemperatureLow,omitempty"`
	ApparentTemperatureLowTime  int64   `json:"apparentTemperatureLowTime,omitempty"`
	ApparentTemperatureMax      float64 `json:"apparentTemperatureMax,omitempty"`
	ApparentTemperatureMaxTime  int64   `json:"apparentTemperatureMaxTime,omitempty"`
	ApparentTemperatureMin      float64 `json:"apparentTemperatureMin,omitempty"`
	ApparentTemperatureMinTime  int64   `json:"apparentTemperatureMinTime,omitempty"`
	Temperature                 float64 `json:"temperature,omitempty"`
	ApparentTemperature         float64 `json:"apparentTemperature,omitempty"`
	CloudCover                  float64 `json:"cloudCover,omitempty"`
	DewPoint                    float64 `json:"dewPoint,omitempty"`
	Humidity                    float64 `json:"humidity,omitempty"`
	Icon                        string  `json:"icon,omitempty"`
	MoonPhase                   float64 `json:"moonPhase,omitempty"`
	Ozone                       float64 `json:"ozone,omitempty"`
	PrecipIntensity             float64 `json:"precipIntensity,omitempty"`
	PrecipIntensityMax          float64 `json:"precipIntensityMax,omitempty"`
	PrecipIntensityMaxTime      int64   `json:"precipIntensityMaxTime,omitempty"`
	PrecipProbability           float64 `json:"precipProbability,omitempty"`
	PrecipType                  string  `json:"precipType,omitempty"`
	Pressure                    float64 `json:"pressure,omitempty"`
	Summary                     string  `json:"summary,omitempty"`
	SunriseTime                 int64   `json:"sunriseTime,omitempty"`
	SunsetTime                  int64   `json:"sunsetTime,omitempty"`
	TemperatureHigh             float64 `json:"temperatureHigh,omitempty"`
	TemperatureHighTime         int64   `json:"temperatureHighTime,omitempty"`
	TemperatureLow              float64 `json:"temperatureLow,omitempty"`
	TemperatureLowTime          int64   `json:"temperatureLowTime,omitempty"`
	TemperatureMax              float64 `json:"temperatureMax,omitempty"`
	TemperatureMaxTime          int64   `json:"temperatureMaxTime,omitempty"`
	TemperatureMin              float64 `json:"temperatureMin,omitempty"`
	TemperatureMinTime          int64   `json:"temperatureMinTime,omitempty"`
	Time                        int64   `json:"time,omitempty"`
	UvIndex                     int64   `json:"uvIndex,omitempty"`
	UvIndexTime                 int64   `json:"uvIndexTime,omitempty"`
	Visibility                  float64 `json:"visibility,omitempty"`
	WindBearing                 int64   `json:"windBearing,omitempty"`
	WindGust                    float64 `json:"windGust,omitempty"`
	WindGustTime                int64   `json:"windGustTime,omitempty"`
	WindSpeed                   float64 `json:"windSpeed,omitempty"`
}

// ToData Datum converts to *data.Data.
func (d *Datum) ToData(
	typ data.DataType,
	lat, lon float64,
	tz string,
) *data.Data {
	if d == nil {
		return nil
	}

	return &data.Data{
		Type:                        typ,
		Latitude:                    lat,
		Longitude:                   lon,
		Timezone:                    tz,
		ApparentTemperatureHigh:     d.ApparentTemperatureHigh,
		ApparentTemperatureHighTime: d.ApparentTemperatureHighTime,
		ApparentTemperatureLow:      d.ApparentTemperatureLow,
		ApparentTemperatureLowTime:  d.ApparentTemperatureLowTime,
		ApparentTemperatureMax:      d.ApparentTemperatureMax,
		ApparentTemperatureMaxTime:  d.ApparentTemperatureMaxTime,
		ApparentTemperatureMin:      d.ApparentTemperatureMin,
		ApparentTemperatureMinTime:  d.ApparentTemperatureMinTime,
		Temperature:                 d.Temperature,
		ApparentTemperature:         d.ApparentTemperature,
		CloudCover:                  d.CloudCover,
		DewPoint:                    d.DewPoint,
		Humidity:                    d.Humidity,
		Icon:                        d.Icon,
		MoonPhase:                   d.MoonPhase,
		Ozone:                       d.Ozone,
		PrecipIntensity:             d.PrecipIntensity,
		PrecipIntensityMax:          d.PrecipIntensityMax,
		PrecipIntensityMaxTime:      d.PrecipIntensityMaxTime,
		PrecipProbability:           d.PrecipProbability,
		PrecipType:                  d.PrecipType,
		Pressure:                    d.Pressure,
		Summary:                     d.Summary,
		SunriseTime:                 d.SunriseTime,
		SunsetTime:                  d.SunsetTime,
		TemperatureHigh:             d.TemperatureHigh,
		TemperatureHighTime:         d.TemperatureHighTime,
		TemperatureLow:              d.TemperatureLow,
		TemperatureLowTime:          d.TemperatureLowTime,
		TemperatureMax:              d.TemperatureMax,
		TemperatureMaxTime:          d.TemperatureMaxTime,
		TemperatureMin:              d.TemperatureMin,
		TemperatureMinTime:          d.TemperatureMinTime,
		Time:                        d.Time,
		UvIndex:                     d.UvIndex,
		UvIndexTime:                 d.UvIndexTime,
		Visibility:                  d.Visibility,
		WindBearing:                 d.WindBearing,
		WindGust:                    d.WindGust,
		WindGustTime:                d.WindGustTime,
		WindSpeed:                   d.WindSpeed,
	}
}

// DatumBlock contains a collection of Datum supplemented with general summary
// and icon.
type DatumBlock struct {
	Data    []*Datum `json:"data,omitempty"`
	Summary string   `json:"summary,omitempty"`
	Icon    string   `json:"icon,omitempty"`
}

// GetData retrieves a slice of Datum from DatumBlock.
func (d *DatumBlock) GetData() []*Datum {
	if d == nil {
		return nil
	}
	return d.Data
}

// Alert contains objects representing the severe weather warnings issued for
// the requested location by a governmental authority.
type Alert struct {
	Description string   `json:"description,omitempty"`
	Expires     int64    `json:"expires,omitempty"`
	Regions     []string `json:"regions,omitempty"`
	Severity    string   `json:"severity,omitempty"`
	Time        int64    `json:"time,omitempty"`
	Title       string   `json:"title,omitempty"`
	URI         string   `json:"uri,omitempty"`
}

// Flags contains various metadata information related to the request.
type Flags struct {
	DarkSkyUnavailable string   `json:"darksky-unavailable,omitempty"`
	NearestStation     float64  `json:"nearest-station"`
	Sources            []string `json:"sources,omitempty"`
	Units              string   `json:"units,omitempty"`
}

// Response is the response containing weather forecast data requested
// from https://api.darksky.net.
type Response struct {
	Latitude  float64     `json:"latitude,omitempty"`
	Longitude float64     `json:"longitude,omitempty"`
	Timezone  string      `json:"timezone,omitempty"`
	Currently *Datum      `json:"currently,omitempty"`
	Minutely  *DatumBlock `json:"minutely,omitempty"`
	Hourly    *DatumBlock `json:"hourly,omitempty"`
	Daily     *DatumBlock `json:"daily,omitempty"`
	Alerts    []*Alert    `json:"alerts,omitempty"`
	Flags     *Flags      `json:"flags,omitempty"`
}

// ToData converts response to the slice of *data.Data.
func (r *Response) ToData() []*data.Data {
	dd := make(
		[]*data.Data,
		0,
		1+len(r.Minutely.GetData())+
			len(r.Hourly.GetData())+
			len(r.Daily.GetData()),
	)

	dd = append(dd, r.Currently.ToData(
		data.DataType_CURRENT,
		r.Latitude,
		r.Longitude,
		r.Timezone,
	))

	for _, d := range r.Minutely.GetData() {
		dd = append(dd, d.ToData(
			data.DataType_MINUTELY,
			r.Latitude,
			r.Longitude,
			r.Timezone,
		))
	}

	for _, d := range r.Hourly.GetData() {
		dd = append(dd, d.ToData(
			data.DataType_HOURLY,
			r.Latitude,
			r.Longitude,
			r.Timezone,
		))
	}

	for _, d := range r.Daily.GetData() {
		dd = append(dd, d.ToData(
			data.DataType_DAILY,
			r.Latitude,
			r.Longitude,
			r.Timezone,
		))
	}

	return dd
}
