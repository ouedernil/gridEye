package http_accessors

import (
	"net/http"
	"encoding/json"
	"io/ioutil"
	"fmt"
)

func UpdateNetwork(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		//Read the body request
		jsonContent, err1 := ioutil.ReadAll(r.Body)
		if err1 != nil {
			http.Error(w, fmt.Sprintf("Error reading request: %v", err1), 500)
			return
		}
		var replacementPairs map[string] string
		//De-serialize the json body request
		err2 := json.Unmarshal(jsonContent, &replacementPairs)
		if err2 != nil {
			http.Error(w, fmt.Sprintf("Error converting json to map: %v", err2), 500)
			return
		}
		defer r.Body.Close()
		//Update newtork configurations
		if configureNetwork(replacementPairs) {
			w.WriteHeader(http.StatusOK)
			defer r.Body.Close()
			rebootSystem()
		}else{
			http.Error(w, fmt.Sprintf("Bad request"), 400)
			return
		}

}


