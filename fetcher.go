package mackerelsesame

import (
	"time"

	"github.com/Arthur1/mackerel-sesame/sesame"
)

type SesameMetrics struct {
	DeviceName        string
	Timestamp         time.Time
	BatteryPercentage int64
	BatteryVoltage    float64
}

type SesameFetcher struct {
	client *sesame.Client
}

func NewSesameFetcher(apiKey string) *SesameFetcher {
	client := sesame.NewClient(apiKey)
	return &SesameFetcher{client: client}
}

func (s *SesameFetcher) Fetch(deviceUUID, deviceName string) (*SesameMetrics, error) {
	status, err := s.client.GetDeviceStatus(deviceUUID)
	if err != nil {
		return nil, err
	}
	metrics := &SesameMetrics{
		DeviceName:        deviceName,
		Timestamp:         time.Now(),
		BatteryPercentage: status.BatteryPercentage,
		BatteryVoltage:    status.BatteryVoltage,
	}
	return metrics, nil
}
