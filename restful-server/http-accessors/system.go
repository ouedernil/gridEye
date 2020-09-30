package http_accessors

import (
	"io/ioutil"
	"log"
	"strings"
	"syscall"
	"os/exec"
)


const drv string = "/home/grideye/drv/config/drv.cfg"
const mon string = "/home/grideye/mon/config/mon.cfg"
const net string = "/etc/network/interfaces.d/"


func getGSMParam() []string{
	var gsmParam []string
	input, err := ioutil.ReadFile(drv)
	if err != nil {
		gsmParam[0] = "error"
		gsmParam[1] = err.Error()
		log.Fatalln(err)
	}else{
		lines := strings.Split(string(input), "\n")
		var enable = ""
		var apn = ""
		for i := range lines {
			if strings.Contains(lines[i], "gsm-enable") {
				// Split on equal.
				result := strings.Split(lines[i], "=")
				// Display all elements.
				enable = result[1]
				rmAfter := strings.Split(enable, ";")
				enable = strings.Replace(rmAfter[0], ";", "", -1)
				enable = strings.Replace(enable, " ", "", -1)
				gsmParam = append(gsmParam, enable)
			}

			if strings.Contains(lines[i], "gsm-apn") {
				// Split on equal.
				result := strings.Split(lines[i], "=")
				// Display all elements.
				apn = strings.Replace(result[1], "\"", "", -1)
				rmAfter := strings.Split(apn, ";")
				apn = strings.Replace(rmAfter[0], ";", "", -1)
				apn = strings.Replace(apn, " ", "", -1)
				gsmParam = append(gsmParam, apn)
			}
		}
	}
	return gsmParam


}

func getNetworkParam(iface string) []string {
	var networkParam []string
	var address string = ""
	var netmask string = ""
	var gateway string = ""
	if(iface == "eth0" || iface == "eth1"){
		if isUp(iface) {
			input, err := ioutil.ReadFile(net+iface)
			if err != nil {
				networkParam[0] = "error"
				networkParam[1] = err.Error()
			}else{
				lines := strings.Split(string(input), "\n")
				if strings.Contains(lines[1], "dhcp") {
					networkParam = append(networkParam, "dhcp")
					out, err := exec.Command("bash", "-c", "ifconfig "+iface).Output()
					if err != nil {
						networkParam[0] = "error"
						networkParam[1] = err.Error()
					}else{
						networkParam = append(networkParam, "dhcp")
						linesIfconfig := strings.Split(string(out), "\n")
						for i := range lines {
							if strings.Contains(linesIfconfig[i], "inet") {
								// Get ip address
								resultAddress1 := strings.Split(linesIfconfig[i], "inet ")
								firstSplitIp := resultAddress1[1]
								resultAddress2 := strings.Split(firstSplitIp, " ")
								address = resultAddress2[0]
								networkParam = append(networkParam, address)

								resultNetmask1 := strings.Split(linesIfconfig[i], "netmask ")
								firstSplitNetmask := resultNetmask1[1]
								resultNetmask2 := strings.Split(firstSplitNetmask, " ")
								netmask = resultNetmask2[0]
							}
						}
					}
				}else{
					networkParam = append(networkParam, "manual")
					for i := range lines {
						if strings.Contains(lines[i], "address") {
							// Split on space.
							result := strings.Split(lines[i], "address ")
							address = result[1]
							networkParam = append(networkParam, address)
						}
						if strings.Contains(lines[i], "netmask") {
							// Split on space.
							result := strings.Split(lines[i], "netmask ")
							netmask = result[1]
							networkParam = append(networkParam, netmask)
						}
						if strings.Contains(lines[i], "gateway") {
							// Split on space.
							result := strings.Split(lines[i], "gateway ")
							gateway = result[1]
							networkParam = append(networkParam, gateway)
						}
					}
				}
			}

		}else{
			networkParam = append(networkParam, "down")
		}
	}
	return networkParam
}

func checkPinCode(currentPin string) bool{
	input, err := ioutil.ReadFile(drv)
	if err != nil {
		log.Fatalln(err)
		return false
	}
	var pin = ""
	lines := strings.Split(string(input), "\n")
	for i := range lines {
		if strings.Contains(lines[i], "gsm-pin") {
			// Split on comma.
			result := strings.Split(lines[i], "=")
			// Display all elements.
			pin = strings.Replace(result[1], "\"", "", -1)
			rmAfter := strings.Split(pin, ";")
			pin = strings.Replace(rmAfter[0], ";", "", -1)
			pin = strings.Replace(pin, " ", "", -1)
		}
	}
	if pin == currentPin {
		return true
	}

	return false
}

func configureNetwork(network map[string] string) bool{
	if(network["inter"] == "eth1" || network["inter"]  == "eth0"){
		var fileToRead string = net+network["inter"]
		input, err := ioutil.ReadFile(fileToRead)
		lines := strings.Split(string(input), "\n")
		for i := range lines {
			lines[i] = ""
		}
		if network["configur"] == "manual"{
			lines[0] = "auto "+ network["inter"]+"\n iface "+ network["inter"] + " inet static \n address  "+ network["ip"]+"\n netmask "+ network["netmask"]+"\n gateway "+ network["gateway"]
		}else if network["configur"] == "dhcp"{
			lines[0] = "auto "+ network["inter"]
			lines[1] = "iface "+ network["inter"] + " inet dhcp"
		}

		output := strings.Join(lines, "\n")
		err = ioutil.WriteFile(net+network["inter"], []byte(output), 0644)
		if err != nil {
			log.Fatalln(err)
			return false
		}
	}else{
		return false
	}
	return true
}

func rebootSystem(){
	syscall.Sync()
	syscall.Reboot(syscall.LINUX_REBOOT_CMD_RESTART)
}

func configurePinCode(pin string) bool{
	input, err := ioutil.ReadFile(drv)
	if err != nil {
		log.Fatalln(err)
		return false
	}
	lines := strings.Split(string(input), "\n")

	for i, line := range lines {
		if strings.Contains(line, "gsm-pin") {
			lines[i] = "gsm-pin = \""+pin+"\";"
		}
	}
	output := strings.Join(lines, "\n")
	err = ioutil.WriteFile(drv, []byte(output), 0644)
	if err != nil {
		log.Fatalln(err)
		return false
	}
	return true
}


func configureApn(apn string) bool{
	input, err := ioutil.ReadFile(drv)
	if err != nil {
		log.Fatalln(err)
		return false
	}
	lines := strings.Split(string(input), "\n")

	for i, line := range lines {
		if strings.Contains(line, "gsm-apn") {
			lines[i] = "gsm-apn = \""+apn+"\";"
		}
	}
	output := strings.Join(lines, "\n")
	err = ioutil.WriteFile(drv, []byte(output), 0644)
	if err != nil {
		log.Fatalln(err)
		return false
	}
	return true
}

func configureEnable(enable string) bool{
	input, err := ioutil.ReadFile(drv)
	if err != nil {
		log.Fatalln(err)
		return false
	}
	lines := strings.Split(string(input), "\n")

	for i, line := range lines {
		if strings.Contains(line, "gsm-enable") {
			lines[i] = "gsm-enable = \""+enable+"\"; #0 = not start, 1 = start"
		}
	}
	output := strings.Join(lines, "\n")
	err = ioutil.WriteFile(drv, []byte(output), 0644)
	if err != nil {
		log.Fatalln(err)
		return false
	}
	return true
}


func isUp (inet string) bool{
	out, err := exec.Command("ethtool", inet).Output()
	if err != nil {
		log.Fatalln(err)
		return false
	}
	lines := strings.Split(string(out), "\n")
	for i := range lines {
		if strings.Contains(string(lines[i]), "Link") {
			if strings.Contains(string(lines[i]), "no") {
				return false
			}
		}
	}
	return true
}

func getMonitorFrequency() string {
	var frequency string = ""
	input, err := ioutil.ReadFile(mon)
	if err != nil {
		log.Fatalln(err)
		return "error"
	}
	lines := strings.Split(string(input), "\n")

	for i := range lines {
		if strings.Contains(lines[i], "monitor_freq") {
			result := strings.Split(lines[i], "=")
			frequency = strings.Replace(result[1], " ", "", -1)
			frequency = strings.Replace(result[1], ";", "", -1)
		}
	}
	return frequency
}

func getAlarmParam(paramType string) [3]string{
	var alarmParam [3]string
	var threshold string = ""
	var hysteresis string = ""
	var violtime string = ""
	var unit string = ""
		input, err := ioutil.ReadFile("/home/grideye/mon/config/mon.cfg")
		if err != nil {
			alarmParam[0] = "error"
			alarmParam[1] = err.Error()
		}else {
			if paramType == "ov" {
				unit = "u_over_"
			}
			if paramType == "uv" {
				unit = "u_under_"
			}
			if paramType == "oc" {
				unit = "i_over_"
			}
			if paramType == "ocn" {
				unit = "in_over_"
			}
			lines := strings.Split(string(input), "\n")
			for i := range lines {
				if strings.Contains(lines[i], unit+"thresh") {
					resultT := strings.Split(lines[i], "=")
					resultT2 := strings.Replace(resultT[1], " ", "", -1)
					threshold = strings.Replace(resultT2, ";", "", -1)
					alarmParam[0] = threshold
				}
				if strings.Contains(lines[i], unit+"hysteresis") {
					resultH := strings.Split(lines[i], "=")
					resultH2 := strings.Replace(resultH[1], " ", "", -1)
					hysteresis = strings.Replace(resultH2, ";", "", -1)
					alarmParam[1] = hysteresis
				}
				if strings.Contains(lines[i], unit+"viol_time") {
					resultV := strings.Split(lines[i], "=")
					resultV2 := strings.Replace(resultV[1], " ", "", -1)
					violtime = strings.Replace(resultV2, ";", "", -1)
					alarmParam[2] = violtime
				}

			}
		}
	return alarmParam
}