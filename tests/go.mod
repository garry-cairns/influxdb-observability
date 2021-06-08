module tests

go 1.15

require (
	github.com/containerd/containerd v1.5.2 // indirect
	github.com/influxdata/influxdb-observability/common v0.0.0-20210503044220-4051d4b8738f
	github.com/influxdata/influxdb-observability/influx2otel v0.0.0-20210605003714-a868e4b21ba8 // indirect
	github.com/influxdata/influxdb-observability/otlp v0.0.0-20210605003714-a868e4b21ba8
	github.com/open-telemetry/opentelemetry-collector-contrib/exporter/influxdbexporter v0.0.0-20210607140028-439ff5f266ab
	github.com/open-telemetry/opentelemetry-collector-contrib/receiver/influxdbreceiver v0.0.0-20210607140028-439ff5f266ab
	github.com/sirupsen/logrus v1.8.1 // indirect
	github.com/stretchr/testify v1.7.0
	go.opentelemetry.io/collector v0.27.1-0.20210603182316-5369d7e9e83e
	go.uber.org/zap v1.17.0
	google.golang.org/protobuf v1.26.0
)

replace (
	github.com/influxdata/influxdb-observability/common => ../common
	github.com/influxdata/influxdb-observability/influx2otel => ../influx2otel
	github.com/influxdata/influxdb-observability/otlp => ../otlp
)