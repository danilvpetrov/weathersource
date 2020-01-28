package mysqlstorage

import (
	"context"

	"github.com/danilvpetrov/weathersource/data"
)

// Save persists the forecast into a MySQL database.
func (s *Storage) Save(ctx context.Context, data []*data.Data) error {
	for _, d := range data {
		if err := s.saveData(ctx, d); err != nil {
			return err
		}
	}

	return nil
}

func (s *Storage) saveData(ctx context.Context, data *data.Data) error {
	_, err := s.DB.ExecContext(
		ctx,
		`REPLACE INTO weather_data (
			data_time,
			data_type,
			timezone,
			latitude,
			longitude,
			apparent_temperature_high,
			apparent_temperature_high_time,
			apparent_temperature_low,
			apparent_temperature_low_time,
			apparent_temperature_max,
			apparent_temperature_max_time,
			apparent_temperature_min,
			apparent_temperature_min_time,
			temperature,
			apparent_temperature,
			cloud_cover,
			dew_point,
			humidity,
			icon,
			moon_phase,
			ozone,
			precip_intensity,
			precip_intensity_max,
			precip_intensity_max_time,
			precip_probability,
			precip_type,
			pressure,
			summary,
			sunrise_time,
			sunset_time,
			temperature_high,
			temperature_high_time,
			temperature_low,
			temperature_low_time,
			temperature_max,
			temperature_max_time,
			temperature_min,
			temperature_min_time,
			uv_index,
			uv_index_time,
			visibility,
			wind_bearing,
			wind_gust,
			wind_gust_time,
			wind_speed
		) VALUES (
			?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?,
			?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?,
			?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?
		)
		`,
		data.Time,
		data.Type.String(),
		data.Timezone,
		data.Latitude,
		data.Longitude,
		data.ApparentTemperatureHigh,
		data.ApparentTemperatureHighTime,
		data.ApparentTemperatureLow,
		data.ApparentTemperatureLowTime,
		data.ApparentTemperatureMax,
		data.ApparentTemperatureMaxTime,
		data.ApparentTemperatureMin,
		data.ApparentTemperatureMinTime,
		data.Temperature,
		data.ApparentTemperature,
		data.CloudCover,
		data.DewPoint,
		data.Humidity,
		data.Icon,
		data.MoonPhase,
		data.Ozone,
		data.PrecipIntensity,
		data.PrecipIntensityMax,
		data.PrecipIntensityMaxTime,
		data.PrecipProbability,
		data.PrecipType,
		data.Pressure,
		data.Summary,
		data.SunriseTime,
		data.SunsetTime,
		data.TemperatureHigh,
		data.TemperatureHighTime,
		data.TemperatureLow,
		data.TemperatureLowTime,
		data.TemperatureMax,
		data.TemperatureMaxTime,
		data.TemperatureMin,
		data.TemperatureMinTime,
		data.UvIndex,
		data.UvIndexTime,
		data.Visibility,
		data.WindBearing,
		data.WindGust,
		data.WindGustTime,
		data.WindSpeed,
	)

	return err
}
