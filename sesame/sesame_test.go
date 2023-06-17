package sesame

import (
	"net/http"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func TestGetDeviceStatus(t *testing.T) {
	t.Run("fetch the device status from Sesame Web API", func(t *testing.T) {
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()
		httpmock.RegisterResponder(
			http.MethodGet, "https://app.candyhouse.co/api/sesame2/id",
			httpmock.NewJsonResponderOrPanic(http.StatusOK, map[string]any{
				"batteryPercentage": 94,
				"batteryVoltage":    5.869794721407625,
				"position":          11,
				"CHSesame2Status":   "locked",
				"timestamp":         1598523693,
			}),
		)
		client := NewClient("")
		want := &DeviceStatus{
			BatteryPercentage: 94,
			BatteryVoltage:    5.869794721407625,
			Position:          11,
			CHSesame2Status:   "locked",
			Timestamp:         1598523693,
			Wm2State:          false,
		}
		got, err := client.GetDeviceStatus("id")
		assert.NoError(t, err)
		assert.Empty(t, cmp.Diff(want, got))
	})
}
