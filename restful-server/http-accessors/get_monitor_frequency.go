package http_accessors

import (
	"net/http"
	"encoding/json"
	"fmt"
)

type Struct_monitor_frequency struct {
	Frequency  string `json:"frequency"`
}
func GetMonitorFrequency(w http.ResponseWriter, r *http.Request) {
	frequency := getMonitorFrequency()
	var p []string
	p = append(p, frequency)
	w.Header().Set("Content-Type", "application/json")

	var parameters = map[string]*Struct_monitor_frequency{
		"MonitorFrequency": {Frequency: p[0]},
	}

	param, error := json.Marshal(parameters)

	if error != nil || p[0] == "error"{
		http.Error(w, fmt.Sprintf("Error converting to json %v", error), 500)
		return
	}

	fmt.Fprint(w, string(param))
}