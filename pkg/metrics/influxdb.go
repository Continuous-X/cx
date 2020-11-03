package metrics

import (
	"fmt"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

type InfluxDB struct {
	client influxdb2.Client
}

var influxDB *InfluxDB

func (influxDB InfluxDB) Connect() *InfluxDB {
	if influxDB.client == nil {
		influxDB.client = influxdb2.NewClientWithOptions(fmt.Sprintf("%s://%s:%s", influxDbProtocol, influxDbHostname, influxDbPort), "", influxdb2.DefaultOptions())
	}
	return &influxDB
}
