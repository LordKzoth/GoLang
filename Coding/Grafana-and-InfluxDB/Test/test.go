package main

import (
	"fmt"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

func main() {
	name1, err := connectToInfluxDbClient()
	if err != nil {
		panic("Error!")
	}

	fmt.Println(name1)
}

func connectToInfluxDbClient() (influxdb2.Client, error) {

	return influxdb2.NewClient("kek", "lol"), nil
}
