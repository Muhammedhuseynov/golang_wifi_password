package pck

import (
	"fmt"
	"os/exec"
	"strings"
)

func GetWifiNames() []string {
	command, err := exec.Command("netsh", "wlan", "show", "profiles").Output()
	if err != nil {
		fmt.Println("Error")
	}
	//convert byte code 2 string
	cmdStr := fmt.Sprintf("%s", command)

	// start parsing string
	cmdSplited := strings.Split(cmdStr, "All User Profile")
	var wifi_name []string
	for i := 1; i < len(cmdSplited); i++ {
		// your wifi name will be :wifiname. So I  want ":" split it
		for _, v := range strings.Split(cmdSplited[i], ":") {
			if v != "\n" {
				// got names but inside have some empty strings. In line 33 we will remove it
				wifi_name = append(wifi_name, strings.TrimSpace(v))
			}
		}
	}
	//remove empty strings,get all clean names
	var cleanWifiNames []string
	for _, v := range wifi_name {
		if v != "" {
			cleanWifiNames = append(cleanWifiNames, v)
		}
	}
	return cleanWifiNames
}

func WifiPasswords(wifinames []string) map[string]string {
	wifiPasws := map[string]string{}
	for _, wifiname := range wifinames {
		// get each wifi infos
		command, err := exec.Command("netsh", "wlan", "show", "profile", wifiname, "key=clear").CombinedOutput()
		//if on wifi no wireless interface
		if err != nil {
			fmt.Printf("%s - There is no such wireless interface on the system! Trying again...\n", wifiname)
		} else {
			//convert byte code 2 string
			cmdStr := fmt.Sprintf("%s", command)
			//start parsing string(infos)
			if strings.Contains(cmdStr, "Key Content") == true {
				getSecSetting := strings.Split(cmdStr, "Security settings")[1]
				splitPaswFromSecSetting := strings.Split(getSecSetting, "Key Content")[1]
				//remove last Cost setting
				getPasw := strings.TrimSpace(strings.Split(splitPaswFromSecSetting, "Cost settings")[0])
				wifiPasws[wifiname] = strings.Split(getPasw, ":")[1]
			}
		}
	}
	//for k, v := range wifiPasws {
	//	fmt.Printf("%s : %s\n", k, v)
	//	//if v == "" {
	//	//	fmt.Printf("Empty: %s\n", k)
	//	//	delete(wifiPasws, k)
	//	//}
	//}
	return wifiPasws
}
