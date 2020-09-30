package http_accessors

import (
	"net/http"
	"io/ioutil"
	"encoding/json"
	"fmt"
)


func UpdateGsmPinCode(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=UTF-8")
	jsonContent, err1 := ioutil.ReadAll(r.Body)
	if err1 != nil {
		http.Error(w, fmt.Sprintf("Error reading request: %v", err1), 500)
		return
	}
	var replacementPairs map[string]string
	err2 := json.Unmarshal(jsonContent, &replacementPairs)
	if err2 != nil {
		http.Error(w, fmt.Sprintf("Error converting json to map: %v", err2), 500)
		return
	}
	defer r.Body.Close()
	if checkPinCode(replacementPairs["currentPin"]) {
		if configurePinCode(replacementPairs["pin"]) {
			w.WriteHeader(http.StatusOK)
			defer r.Body.Close()
			rebootSystem()
		}
	}else{
		fmt.Fprintf(w, "The current PIN is wrong. Please enter a valid PIN")
		return
	}

}

