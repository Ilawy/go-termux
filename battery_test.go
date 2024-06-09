package termux

import (
	"testing"
)

func TestBattery(t *testing.T) {
	result, err := BatteryStatus()
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(result)
}
