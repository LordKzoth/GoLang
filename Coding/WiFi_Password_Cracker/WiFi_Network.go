package main

import (
	"fmt"
	"os/exec"
	"strings"
)

// = Colors
var (
	colorReset	= "\033[0m"

	colorRed 	= "\033[31m"
	colorGreen 	= "\033[32m"
	colorYellow	= "\033[33m"
)

// === DECLARATION PART ===
type WiFiNetwork struct {
	// = Information
	SSID 					string 		// WLAN network name
	// = Authentification
	AuthenticationMethod 	string
	EncryprionMethod       	string
	Password             	string
	// = Connection
	BSSID        			[]string 	// MAC address
	Signal       			[]string
	WlanStandart 			[]string
}

type WiFiNetworkTools interface {
	// = Information
	PrintWiFiInfo()
	FindPassword() string
}

// === DEFINITION PART ===
func (wifi WiFiNetwork) PrintWiFiInfo() {
	// = Set Color
	var color string
	
	if wifi.Password != "" {
		color = colorGreen
	} else if wifi.AuthenticationMethod == "Open" {
		color = colorYellow
	} else {
		color = colorReset
	}

	// Print information
	fmt.Print(color)
	fmt.Printf("SSID: %s\n", wifi.SSID)
	if wifi.Password != "" {
		fmt.Printf("Pass: %s\n", wifi.Password)
	} else {
		fmt.Printf("Auth: %s\n", wifi.AuthenticationMethod)
		fmt.Printf("Encr: %s\n", wifi.EncryprionMethod)
	}
	for index := 0; index < len(wifi.BSSID); index++ {
		fmt.Printf("  = %d\n", index + 1)
		fmt.Printf("  BSSID:    %s\n", wifi.BSSID[index])
		fmt.Printf("  Signal:   %s\n", wifi.Signal[index])
		fmt.Printf("  Standart: %s\n", wifi.WlanStandart[index])
	}
	fmt.Print(colorReset)
}

func (wifi WiFiNetwork) FindPassword () string{
	// netsh wlan show profiles name=SSID key=clear
	// = Run command
	commandApp 	:= "netsh"
	commandArgs	:= strings.Fields("wlan show profiles name=" + wifi.SSID + " key=clear")

	command := exec.Command(commandApp, commandArgs...)
	info, _ := command.Output()

	// = Find password value
	for _, line := range strings.Split(string(info), "\r\n") {
		if strings.Contains(line, "Key Content") {
			temp := strings.Fields(line)
			return temp[len(temp) - 1]
		}
	}

	return ""
}