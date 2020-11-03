package metrics

import "errors"

type MetricGHProtection struct {
	BranchName string
}

func (metric MetricGHProtection) Write() error {
	influxClient := InfluxDB{}.Connect()
	if influxClient != nil {
		return nil
	}
	return errors.New("mal was ausgeben")
}
