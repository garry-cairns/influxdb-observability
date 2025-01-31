package otel2influx

import (
	"context"
	"time"

	"github.com/influxdata/influxdb-observability/common"
)

type InfluxWriter interface {
	NewBatch() InfluxWriterBatch
}

type InfluxWriterBatch interface {
	WritePoint(ctx context.Context, measurement string, tags map[string]string, fields map[string]interface{}, ts time.Time, vType common.InfluxMetricValueType) error
	FlushBatch(ctx context.Context) error
}
