package main

import (
	"fmt"
	"os/exec"
	"regexp"
	"sort"
	"strings"
)

var (
	wlans []WiFiNetwork
)


func main() {
	// = Information about Project
	fmt.Printf(`
		================================
		#  ##  #####   ###  #####  #   #
		# #        #  #   #   #    #   #
		##      ###   #   #   #    #####
		# #    #      #   #   #    #   #
		#  ##  #####   ###    #    #   #
		================================

	`)

	ChangeCmdLanguage("437")
	defer ChangeCmdLanguage("866")
	
	// = Get Information about local WLANS
	// == Read information from Netsh Output
	commandApp 	:= "netsh"
	commandArgs := strings.Fields("wlan show networks mode=bssid")

	command 	:= exec.Command(commandApp, commandArgs...)
	wlanInfo, _ := command.Output()

	// == Parse Data
	for _, line := range strings.Split(string(wlanInfo), "\r\n") {
		if regexp.MustCompile("^SSID").MatchString(line) {
			// === Add new WLAN to list
			temp := strings.Fields(line)
			if temp[len(temp) - 1] == ":" {
				temp[len(temp) - 1] = ""
			}
			wlans = append(wlans, WiFiNetwork{SSID: temp[len(temp) - 1]})
			
			// === Find Password (In PC memory)
			wlans[len(wlans) - 1].Password = wlans[len(wlans) - 1].FindPassword()
		} else {
			switch {
			case strings.Contains(line, "Authentication"):
				temp := strings.Fields(line)
				wlans[len(wlans) - 1].AuthenticationMethod = temp[len(temp) - 1]

			case strings.Contains(line, "Encryption"):
				temp := strings.Fields(line)
				wlans[len(wlans) - 1].EncryprionMethod = temp[len(temp) - 1]

			case strings.Contains(line, "BSSID"):
				temp := strings.Fields(line)
				wlans[len(wlans) - 1].BSSID = append(wlans[len(wlans) - 1].BSSID, temp[len(temp) - 1])

			case strings.Contains(line, "Signal"):
				temp := strings.Fields(line)
				wlans[len(wlans) - 1].Signal = append(wlans[len(wlans) - 1].Signal, temp[len(temp) - 1])

			case strings.Contains(line, "Radio type"):
				temp := strings.Fields(line)
				wlans[len(wlans) - 1].WlanStandart = append(wlans[len(wlans) - 1].WlanStandart, temp[len(temp) - 1])

			}
		}
	}

	// == Sort WLANs by Signal
	sort.Slice(wlans, func(i, j int) bool {
		return wlans[i].Signal[0] > wlans[j].Signal[0]
	})

	// == Print Information about WLANs
	for index, wlan := range wlans {
		fmt.Printf("\n= %d =\n", index + 1)
		wlan.PrintWiFiInfo()
	}
}

func ChangeCmdLanguage(charset string) {
	// RU - 866
	// EN - 437

	command := exec.Command("chcp", charset) 
	command.Run()
}