package http_accessors

import (
	"net/http"
	"encoding/json"
	"fmt"
)

type Struct_gsm_param struct {
	Enabled  string `json:"enabled"`
	Apn string `json:"apn"`
}

func GetGsmParam(w http.ResponseWriter, r *http.Request) {
	var p []string = getGSMParam()
	w.Header().Set("Content-Type", "application/json")

	var parameters = map[string]*Struct_gsm_param{
		"GSMParameters": {Enabled: p[0], Apn: p[1]},
	}

	param, error := json.Marshal(parameters)

	if error != nil || p[0] == "error"{
		http.Error(w, fmt.Sprintf("Error converting to json : "+ p[1] +"-> %v", error), 500)
		return
	}
	fmt.Fprint(w, string(param))
}
