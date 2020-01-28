package forecast

import (
	"context"

	"github.com/danilvpetrov/weathersource/data"
	"github.com/danilvpetrov/weathersource/storage"
)

// CopyFrom copies forecast data from the source that implements
// storage.DataAccessor.
func (f *Forecast) CopyFrom(ctx context.Context, da storage.DataAccessor) error {

	dd, err := da.GetLatest(ctx)
	if err != nil {
		return err
	}

	for _, d := range dd {
		d := d // capture the variable
		switch d.Type {
		case data.DataType_CURRENT:
			f.Current = d
		case data.DataType_MINUTELY:
			f.Minutely = append(f.Minutely, d)
		case data.DataType_HOURLY:
			f.Hourly = append(f.Hourly, d)
		case data.DataType_DAILY:
			f.Daily = append(f.Daily, d)
		}
	}

	return nil
}
