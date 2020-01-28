package mysqlstorage

import (
	"context"
	"time"

	"github.com/danilvpetrov/weathersource/data"
)

// GetLatest retrieves the latest  from the storage along with the
// current meterological data. TO-DO: specify the type of the forecast and
// geolocation coordinates to retrieve.
func (s *Storage) GetLatest(ctx context.Context) ([]*data.Data, error) {
	var data []*data.Data
	current, err := s.getCurrent(ctx)
	if err != nil {
		return nil, err
	}

	data = append(data, current)

	forecasts, err := s.getForecasts(ctx)
	if err != nil {
		return nil, err
	}

	data = append(data, forecasts...)

	return data, nil
}

func (s *Storage) getCurrent(ctx context.Context) (*data.Data, error) {
	row := s.DB.QueryRowContext(
		ctx,
		`SELECT * FROM weather_data
		WHERE data_type = ?
		ORDER BY data_time DESC LIMIT 1`,
		data.DataType_CURRENT.String(),
	)

	return scan(row.Scan)
}

func (s *Storage) getForecasts(ctx context.Context) ([]*data.Data, error) {
	var dd []*data.Data

	rows, err := s.DB.QueryContext(
		ctx,
		`SELECT * FROM weather_data
		WHERE data_type != ?
		AND data_time > ?`,
		data.DataType_CURRENT.String(),
		time.Now().Unix(),
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		d, err := scan(rows.Scan)
		if err != nil {
			return nil, err
		}
		dd = append(dd, d)
	}

	if err := rows.Close(); err != nil {
		return nil, err
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return dd, nil
}

func scan(f func(dest ...interface{}) error) (*data.Data, error) {
	var (
		d   data.Data
		typ string
	)

	err := f(
		&d.Time,
		&typ,
		&d.Timezone,
		&d.Latitude,
		&d.Longitude,
		&d.ApparentTemperatureHigh,
		&d.ApparentTemperatureHighTime,
		&d.ApparentTemperatureLow,
		&d.ApparentTemperatureLowTime,
		&d.ApparentTemperatureMax,
		&d.ApparentTemperatureMaxTime,
		&d.ApparentTemperatureMin,
		&d.ApparentTemperatureMinTime,
		&d.Temperature,
		&d.ApparentTemperature,
		&d.CloudCover,
		&d.DewPoint,
		&d.Humidity,
		&d.Icon,
		&d.MoonPhase,
		&d.Ozone,
		&d.PrecipIntensity,
		&d.PrecipIntensityMax,
		&d.PrecipIntensityMaxTime,
		&d.PrecipProbability,
		&d.PrecipType,
		&d.Pressure,
		&d.Summary,
		&d.SunriseTime,
		&d.SunsetTime,
		&d.TemperatureHigh,
		&d.TemperatureHighTime,
		&d.TemperatureLow,
		&d.TemperatureLowTime,
		&d.TemperatureMax,
		&d.TemperatureMaxTime,
		&d.TemperatureMin,
		&d.TemperatureMinTime,
		&d.UvIndex,
		&d.UvIndexTime,
		&d.Visibility,
		&d.WindBearing,
		&d.WindGust,
		&d.WindGustTime,
		&d.WindSpeed,
	)

	d.Type = data.DataType(data.DataType_value[typ])

	return &d, err
}
