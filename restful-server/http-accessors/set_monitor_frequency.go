package http_accessors

import (
	"net/http"
	"io/ioutil"
	"encoding/json"
	"strconv"
	"fmt"
)

func UpdateMonitorFrequency(w http.ResponseWriter, r *http.Request) {
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
	var tMRC uint32 = 2
	tV, err := strconv.ParseInt(replacementPairs["tagValue"], 10, 32)
	protoMaker(tMRC, uint32(tV))
	defer r.Body.Close()
	w.WriteHeader(http.StatusOK)
}