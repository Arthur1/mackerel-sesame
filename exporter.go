package mackerelsesame

import (
	"fmt"
	"time"

	"github.com/mackerelio/mackerel-client-go"
)

type MackerelExporter struct {
	client      *mackerel.Client
	serviceName string
}

func NewMackerelExporter(apiKey, serviceName string) *MackerelExporter {
	client := mackerel.NewClient(apiKey)
	return &MackerelExporter{client: client, serviceName: serviceName}
}

func (e *MackerelExporter) Export(metrics *SesameMetrics) error {
	time := metrics.Timestamp.Round(time.Minute).Unix()
	err := e.client.PostServiceMetricValues(e.serviceName, []*mackerel.MetricValue{
		{
			Name:  fmt.Sprintf("sesame.battery_percentage.%s", metrics.DeviceName),
			Time:  time,
			Value: metrics.BatteryPercentage,
		},
		{
			Name:  fmt.Sprintf("sesame.battery_voltage.%s", metrics.DeviceName),
			Time:  time,
			Value: metrics.BatteryVoltage,
		},
	})
	return err
}
