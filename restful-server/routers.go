package main

import (
	"net/http"
	"github.com/gorilla/mux"
	ha "restful-server/http-accessors"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}
	return router
}



var routes = Routes{
	Route{
		"GetGsmParam",
		"GET",
		"/GridEye/1.0/gsm_param/",
		ha.GetGsmParam,
	},

	Route{
		"GetNetworkParam",
		"GET",
		"/GridEye/1.0/network_param/{iface}",
		ha.GetNetworkParam,
	},

	Route{
		"GetAlarmParam",
		"GET",
		"/GridEye/1.0/alarm_param/{para}",
		ha.GetAlarmParam,
	},


	Route{
		"UpdateGsmApn",
		"PUT",
		"/GridEye/1.0/gsm_param/apn",
		ha.UpdateGsmApn,
	},

	Route{
		"UpdateNetwork",
		"PUT",
		"/GridEye/1.0/network_param",
		ha.UpdateNetwork,
	},

	Route{
		"UpdateGsmPinCode",
		"PUT",
		"/GridEye/1.0/gsm_param/pin_code",
		ha.UpdateGsmPinCode,
	},

	Route{
		"UpdateGsmEnabled",
		"PUT",
		"/GridEye/1.0/gsm_param/enabled",
		ha.UpdateGsmEnabled,
	},

	Route{
		"GetMonitorFrequency",
		"GET",
		"/GridEye/1.0/monitor_frequency",
		ha.GetMonitorFrequency,
	},

	Route{
		"UpdateMonitorFrequency",
		"PUT",
		"/GridEye/1.0/monitor_frequency",
		ha.UpdateMonitorFrequency,
	},

	Route{
		"UpdateAlarmParameters",
		"PUT",
		"/GridEye/1.0/alarm_param",
		ha.UpdateAlarmParameters,
	},
}