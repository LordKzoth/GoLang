package main

import (
	"context"
	"fmt"
	"github.com/influxdata/influxdb-client-go/v2"
)

// Connection Data
const token = "B7vskD5298p5TwvRHrQpnzTaw7tXMtR3WLKwv4qIcOu-AGe0yuB5DV_dAalYczs6VtgaP40G7SWSyRd4ViBdpw=="

//const bucket = "Entry"
//const org = "KLegion"

func main() {
	dbConnection, err := connectToInfluxDbClient()
	defer dbConnection.Close()
	if err != nil {
		panic("connection error occurred")
	} else {
		fmt.Println("[OK] Connection info: ", dbConnection)
	}
}

func connectToInfluxDbClient() (influxdb2.Client, error) {
	client := influxdb2.NewClient("http://localhost:8086", token)
	_, err := client.Health(context.Background())

	return client, err
}
