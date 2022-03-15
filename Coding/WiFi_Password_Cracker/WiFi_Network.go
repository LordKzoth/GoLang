package main

// === DECLARATION PART ===
type WiFiNetwork struct {
	// = Information
	SSID 					string 		// WLAN network name
	// = Authentification
	AuthentificationMethod 	string
	EncryprionMethod       	string
	Password             	string
	// = Connection
	BSSID        			[]string 	// MAC address
	Signal       			[]string
	WLANStandart 			[]string
}

type WiFiNetworkTools interface {
	// = Information
	PrintWiFiInfo()
	AddBSSIDInfo(_BSSID string, _Signal int, _WLANStandart string)
}

// === DEFINITION PART ===

func PrintWiFiInfo() {
	// Some code
}

func AddBSSIDInfo(_BSSID string, _Signal int, _WLANStandart string) {
	// Some code
}