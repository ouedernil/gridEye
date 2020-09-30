package main

import (
	"log"
	"net/http"
	"gopkg.in/jcelliott/turnpike.v2"
	"encoding/json"
	"fmt"
	"time"
	"strconv"
)

var client *turnpike.Client


type Struct_alarm_message struct {
	U1_overbase  bool `json:"u1_overbase"`
	U2_overbase  bool `json:"u2_overbase"`
	U3_overbase  bool `json:"u3_overbase"`
	U1_underbase  bool `json:"u1_underbase"`
	U2_underbase  bool `json:"u2_underbase"`
	U3_underbase  bool `json:"u3_underbase"`
	I1_overbase  bool `json:"i1_overbase"`
	I2_overbase  bool `json:"i2_overbase"`
	I3_overbase  bool `json:"i3_overbase"`
	In_over  bool `json:"in_over"`
	Date_time string `json:"date_time"`
}

func wampInit() {
	turnpike.Debug()
	s := turnpike.NewBasicWebsocketServer("grideye.ws")
	allowAllOrigin := func(r *http.Request) bool { return true }
	s.Upgrader.CheckOrigin = allowAllOrigin
	server := &http.Server{
		Handler: s,
		Addr:    ":9000",
	}
	client, _ = s.GetLocalClient("grideye.ws", nil)
	log.Print("["+time.Now().String()+"] - Log message : ")
	log.Println("turnpike server starting on port 9000")
	log.Fatal(server.ListenAndServe())
}

// publishMsg publish the recieved event on its topic.
// It checks the type of the message, and send it on its websocket
func publishMsg(msg *MonitoringEvt){
	if msg.MsgId.String() == "measureNotify"{
		//Serialize in JSON format
		jsonMeasures, err := json.Marshal(msg)
		if err != nil {
			fmt.Println(err)
			return
		}
		//Send the JSON structure on "mon_event_measure.top"
		client.Publish("mon_event_measure.top", nil, []interface{}{string(jsonMeasures)}, nil)
	}else if msg.MsgId.String() == "alarmNotify"{
		//Construct the alarm structure
		aEnd, aStart := getAlarmMessages(msg)
		//Serialize in JSON format
		aEndJson, err := json.Marshal(aEnd)
		if err != nil {
			log.Print("["+time.Now().String()+"] - Log message : ")
			log.Println(err)
			return
		}
		aStartJson, err := json.Marshal(aStart)
		if err != nil {
			log.Print("["+time.Now().String()+"] - Log message : ")
			log.Println(err)
			return
		}
		//Send the JSON structure on "mon_event_alarm.top"
		client.Publish("mon_event_alarm.top", nil, []interface{}{string(aEndJson)}, nil)
		client.Publish("mon_event_alarm.top", nil, []interface{}{string(aStartJson)}, nil)
	}
}

// getAlarmMessages read the timestamp, alarmEndType and alarmStartType, to construct the Struct_alarm_message.
// For each bit it checks if the value is 1 or 0
// It returns two object Struct_alarm_message for the start and the end alarm.
func getAlarmMessages(msg *MonitoringEvt) (map[string]*Struct_alarm_message, map[string]*Struct_alarm_message){
	//Get alarmEndType value
	alarmEnd := msg.AlarmStartEndMsg.AlarmEndType
	var ae uint32 = *alarmEnd
	//Get alarmStartType value
	alarmStart := msg.AlarmStartEndMsg.AlarmStartType
	var as uint32 = *alarmStart
	//Get timestamp value
	timestamp := *msg.AlarmStartEndMsg.Timestamp
	//Convert timestamp value to string
	var timeDate string = fmt.Sprint(timestamp)
	//Convert timestamp string value to int64
	tdInt, err := strconv.ParseInt(timeDate, 10, 64)
	if(err != nil){
		log.Print("["+time.Now().String()+"] - Log message : ")
		log.Println(err.Error())
	}

	var count uint = 0
	var alarmEBool []bool
	var alarmSBool []bool
	//Check if each bit is 1 or 0, if the bit is one, add true in a boolean array. If not, add false
	for i := 0; i < 16; i++{
		if(ae & (1 << count) != 0){
			alarmEBool = append(alarmEBool, true)
		}else{
			alarmEBool = append(alarmEBool, false)
		}
		if(as & (1 << count) != 0){
			alarmSBool = append(alarmSBool, true)
		}else{
			alarmSBool = append(alarmSBool, false)
		}
		count++
	}

	//Construct the alarm array to send
	var aEnd = map[string]*Struct_alarm_message{
		"AlarmEndMessages": &Struct_alarm_message{U1_overbase:alarmEBool[15],U2_overbase:alarmEBool[14],U3_overbase:alarmEBool[13],
		U1_underbase:alarmEBool[11], U2_underbase:alarmEBool[10],U3_underbase:alarmEBool[9], I1_overbase:alarmEBool[7],
		I2_overbase:alarmEBool[6], I3_overbase:alarmEBool[5], In_over:alarmEBool[4], Date_time:time.Unix(tdInt, 0).Format(time.RFC822Z)},
	}
	var aStart = map[string]*Struct_alarm_message{
		"AlarmStartMessages": &Struct_alarm_message{U1_overbase:alarmSBool[15],U2_overbase:alarmSBool[14],U3_overbase:alarmSBool[13],
			U1_underbase:alarmSBool[11], U2_underbase:alarmSBool[10],U3_underbase:alarmSBool[9], I1_overbase:alarmSBool[7],
			I2_overbase:alarmSBool[6], I3_overbase:alarmSBool[5], In_over:alarmSBool[4], Date_time:time.Unix(tdInt, 0).Format(time.RFC822Z)},
	}

	return aEnd, aStart

}