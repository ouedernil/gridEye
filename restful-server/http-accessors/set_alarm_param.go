package http_accessors

import (
	"net/http"
	"io/ioutil"
	"encoding/json"
	"strconv"
	"fmt"
)


func UpdateAlarmParameters(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	jsonContent, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error reading request: %v", err), 500)
		return
	}
	var replacementPairs map[string]string

	err2 := json.Unmarshal(jsonContent, &replacementPairs)
	if err2 != nil {
		http.Error(w, fmt.Sprintf("Error converting json to map: %v", err2), 500)
		return
	}

	th, err := strconv.ParseInt(replacementPairs["threshold"], 10, 32)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error converting to int: %v", err), 500)
		return
	}
	hy, err := strconv.ParseInt(replacementPairs["hysteresis"], 10, 32)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error converting to int: %v", err), 500)
		return
	}
	vt, err := strconv.ParseInt(replacementPairs["time"], 10, 32)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error converting to int: %v", err), 500)
		return
	}
	switch replacementPairs["parameter"] {
		case "over current":
			sendOverCurrent(uint32(th), uint32(hy), uint32(vt))
		case "under current":
			sendUndeCurrent(uint32(th), uint32(hy), uint32(vt))
		case "over voltage":
			sendOverVoltage(uint32(th), uint32(hy), uint32(vt))
		case "over current N":
				sendOverCurrentN(uint32(th), uint32(hy), uint32(vt))
		}
	defer r.Body.Close()
	w.WriteHeader(http.StatusOK)

}

func sendOverCurrent(threshold uint32, hysteresis uint32, viol_time uint32){
	var t uint32 = 3
	var vT uint32 = 4
	var h uint32 = 5
	protoMaker(t, threshold)
	protoMaker(h, hysteresis)
	protoMaker(vT, viol_time)
}

func sendUndeCurrent(threshold uint32, hysteresis uint32, viol_time uint32){
	var t uint32 = 6
	var vT uint32 = 7
	var h uint32 = 8
	protoMaker(t, threshold)
	protoMaker(h, hysteresis)
	protoMaker(vT, viol_time)
}

func sendOverVoltage(threshold uint32, hysteresis uint32, viol_time uint32){
	var t uint32 = 10
	var vT uint32 = 11
	var h uint32 = 12
	protoMaker(t, threshold)
	protoMaker(h, hysteresis)
	protoMaker(vT, viol_time)
}

func sendOverCurrentN(threshold uint32, hysteresis uint32, viol_time uint32){
	var t uint32 = 13
	var vT uint32 = 14
	var h uint32 = 15
	protoMaker(t, threshold)
	protoMaker(h, hysteresis)
	protoMaker(vT, viol_time)
}