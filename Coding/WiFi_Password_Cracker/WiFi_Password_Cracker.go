package main

import (
	"fmt"
	"os/exec"
	"regexp"
	"strings"
)

func main() {
	// = Information about Project
	fmt.Printf("= Welcome in SpideRUS WiFi Password Cracker! =\n\n")

	// = Change cmd Language (RU - EN)
	command 		:= exec.Command("chcp", "437")
	_, issueChcp	:= command.Output()

	if issueChcp != nil {
		fmt.Println(issueChcp.Error())
		return
	}

	// == Trash Code for Antivirus evasion
	if true {
		fmt.Printf("\n==============================================================\n")
		var a, b int = 1, 1
		for index := 0; index < 100; index++ {
			b, a = a + b, b
			fmt.Printf("%d => ", a)
		}
		fmt.Printf("\n==============================================================\n\n")
	}
	
	// = Get Information about local WLANS
	// == Read information from Netsh Output
	commandApp 	:= "netsh"
	commandArgs := strings.Fields("wlan show networks mode=bssid")

	command 	= exec.Command(commandApp, commandArgs...)
	WlanInfo, _ := command.Output()

	// == Parse Data
	for _, line := range strings.Split(string(WlanInfo), "\r\n") {
		if regexp.MustCompile("^SSID").MatchString(line) {
			fmt.Println(line)
		}
	}

	// = Change cmd Language (EN - RU)
	command 		= exec.Command("chcp", "866")
	_, issueChcp	= command.Output()

	if issueChcp != nil {
		fmt.Println(issueChcp.Error())
		return
	}
}