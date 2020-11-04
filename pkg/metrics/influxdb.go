package metrics

import (
	"context"
	"fmt"
	"github.com/influxdata/influxdb-client-go/v2/api"
	"github.com/influxdata/influxdb-client-go/v2/api/write"
	"log"
	"time"

	influxClient "github.com/influxdata/influxdb-client-go/v2"
)

type InfluxDB struct {
	MetricName  string
	Tags        map[string]string
	MetricsData map[string]interface{}
}

var (
	metricPoint *write.Point
	writeApi    api.WriteAPIBlocking
	client      influxClient.Client
)

func (influx InfluxDB) WriteMetric() {
	influx.setMetricPoint()
	influx.setWriteApi()
	influx.writeMetricPoint()
}

func (influx InfluxDB) connect() {
	influxUrl := fmt.Sprintf("%s://%s:%s", influxDbProtocol, influxDbHostname, influxDbPort)
	client = influxClient.NewClientWithOptions(
		influxUrl,
		"",
		influxClient.DefaultOptions().SetBatchSize(20))
	defer client.Close()
}

func (influx InfluxDB) setWriteApi() {
	influx.connect()
	writeApi = client.WriteAPIBlocking("", metricDbName+"/autogen")
}

func (influx InfluxDB) setMetricPoint() {
	metricPoint = influxClient.NewPoint(influx.MetricName, influx.Tags, influx.MetricsData, time.Now())
}

func (influx InfluxDB) writeMetricPoint() {
	writeError := writeApi.WritePoint(context.Background(), metricPoint)
	if writeError != nil {
		log.Printf("Write error: %s\n", writeError.Error())
	}
}
