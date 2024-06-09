package termux

import (
	"encoding/json"
)

func BatteryStatus() (TermuxBatteryStatusResul, error) {
	raw, err := CallCommand("termux-battery-status")
	output := &TermuxBatteryStatusResul{}
	json.Unmarshal(raw, output)
	return *output, err
}

type TermuxBatteryStatusResul struct {
	Health      string  `json:"health"`
	Percentage  int     `json:"percentage"`
	Plugged     string  `json:"plugged"`
	Status      string  `json:"status"`
	Temperature float32 `json:"temperature"`
	Current     int     `json:"current"`
}
