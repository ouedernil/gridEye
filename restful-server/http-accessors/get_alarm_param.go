package http_accessors

import (
	"net/http"
	"strings"
	"encoding/json"
	"fmt"
)

type Struct_alarm_param struct {
	Threshold string `json:"threshold"`
	Hysteresis  string `json:"hysteresis"`
	Date_time string `json:"date_time"`
}



func GetAlarmParam(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	path := strings.Split(r.URL.Path, "/alarm_param/")
	var p [3]string = getAlarmParam(path[1])
	defer r.Body.Close()
	if p[0] != "error" {
			var parameters = map[string]*Struct_alarm_param{
				"AlarmParam": {Threshold: p[0], Hysteresis: p[1], Date_time: p[2]},
			}
			param, error := json.Marshal(parameters)
			if error != nil || p[0] == "error"{
				http.Error(w, fmt.Sprintf("Error converting to json %v", error), 500)
				return
			}
			fmt.Fprint(w, string(param))
	}else{
		http.Error(w, fmt.Sprintf("Bad request"), 400)
		return
	}
}