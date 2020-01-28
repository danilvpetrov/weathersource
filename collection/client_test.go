package collection

import (
	"context"
	"testing"
)

func TestClient_Forecast(t *testing.T) {

	c := NewClient(
		"<place token here>",
		WithLatAndLong(-27.2049, 153.0175),
	)

	resp, err := c.Forecast(context.Background())

	if err != nil {
		t.Logf("Client.Forecast() error = %#+v", err)
	}

	t.Logf("Client.Forecast() response = %#+v", resp)

	t.Logf("Client.Forecast().Currently response = %#+v", resp.Currently)

	t.Logf("Client.Forecast().Flags() response = %#+v", resp.Flags)
}
