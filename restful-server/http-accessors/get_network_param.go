package http_accessors

import (
	"net/http"
	"encoding/json"
	"fmt"
	"strings"
)

type Struct_static_network_param struct {
	Config string `json:"config"`
	Ip  string `json:"ip"`
	Netmask string `json:"netmask"`
	Gateway string `json:"gateway"`
}

type Struct_static_network_param_dhcp struct {
	Config string `json:"config"`
	Ip  string `json:"ip"`
	Netmask string `json:"netmask"`
}

func GetNetworkParam(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	path := strings.Split(r.URL.Path, "/network_param/")
	var p []string = getNetworkParam(path[1])
	defer r.Body.Close()
	if p[0] != "error" {
		if p[0] == "manual" {
			var parameters = map[string]*Struct_static_network_param{

				"NetworkParam": {Config: p[0], Ip: p[1], Netmask: p[2], Gateway: p[3]},

			}
			param, error := json.Marshal(parameters)

			if error != nil || p[0] == "error"{

				http.Error(w, fmt.Sprintf("Error converting to json : %v", error), 500)
				return
			}

			fmt.Fprint(w, string(param))

		}else if p[0] == "dhcp" {

			var parameters = map[string]*Struct_static_network_param_dhcp{
				"NetworkParam": {Config: p[0], Ip: p[1], Netmask: p[2]},

			}

			param, error := json.Marshal(parameters)

			if error != nil || p[0] == "error"{

				http.Error(w, fmt.Sprintf("Error converting to json : %v", error), 500)
				return
			}

			fmt.Fprint(w, string(param))

		}else if p[0] == "down" {
			var parameters = map[string]*Struct_static_network_param{

				"NetworkParam": {Config: "Interface " + path[1] + " is down"},

			}

			param, error := json.Marshal(parameters)

			if error != nil || p[0] == "error"{

				http.Error(w, fmt.Sprintf("Error converting to json : %v", error), 500)
				return

			}

			fmt.Fprintf(w,string(param))
		}
	}else{

		fmt.Fprintf(w, p[0]+" : "+p[1])

	}
}
