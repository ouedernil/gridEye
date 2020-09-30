package http_accessors

import (
"net/http"
"io/ioutil"
"encoding/json"
	"log"
	"fmt"
)


func UpdateGsmEnabled(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	jsonContent, err1 := ioutil.ReadAll(r.Body)
	if err1 != nil {
		http.Error(w, fmt.Sprintf("Error reading request", err1), 500)
		return
	}
	var replacementPairs map[string]string

	err2 := json.Unmarshal(jsonContent, &replacementPairs)
	if err2 != nil {
		http.Error(w, fmt.Sprintf("Error converting json to map: %v", err2), 500)
		return
	}
	defer r.Body.Close()
	if configureEnable(replacementPairs["enable"]) {
		w.WriteHeader(http.StatusOK)
		defer r.Body.Close()
		rebootSystem()

	}else{
		http.Error(w, fmt.Sprintf("Bad request"), 400)
		return
	}
}
