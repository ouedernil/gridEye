package http_accessors

import (
	"net/http"
	"io/ioutil"
	"encoding/json"
	"fmt"
)


func UpdateGsmApn(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		jsonContent, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		var replacementPairs map[string]string

		err2 := json.Unmarshal(jsonContent, &replacementPairs)
		if err2 != nil {
			http.Error(w, fmt.Sprintf("Error converting json to map: %v", err2), 500)
			return
		}
		defer r.Body.Close()
		if configureApn(replacementPairs["apn"]) {
			w.WriteHeader(http.StatusOK)
			defer r.Body.Close()
			rebootSystem()
		}
}

