package main

import (
	"context"
	"fmt"
	"github.com/influxdata/influxdb-client-go/v2"
	"math/rand"
	"time"
)

// Connection Data
const token = "B7vskD5298p5TwvRHrQpnzTaw7tXMtR3WLKwv4qIcOu-AGe0yuB5DV_dAalYczs6VtgaP40G7SWSyRd4ViBdpw=="
const bucket = "Entry"
const org = "KLegion"

func main() {
	dbConnection, err := connectToInfluxDBClient()
	defer dbConnection.Close()
	if err != nil {
		panic("connection error occurred")
	} else {
		fmt.Println("[OK] Connection info: ", dbConnection)
	}

	// Write data - Test
	for i := 0; i < 10; i++ {
		err := writeData(dbConnection, map[string]interface{}{
			"temperature": 35 + rand.Float64()*(50-35),
		})
		if err != nil {
			panic("data writing error")
		}
	}

	// Read first 10 rows (from added data)
	fmt.Println("\n1st Query:")
	firstQueryResponse, err := executeQuery(dbConnection, `from (bucket: "Entry")
																	|> range(start: -1h)
																	|> filter(fn: (r) => r._field == "temperature" and r._value <= 50.0)
																	|> group()
																	|> sort(columns: ["_value"], desc: true)
																	|> limit(n: 10, offset: 0)
																	|> keep(columns: ["_time", "_value", "sensor_id"])`)
	if err != nil {
		panic(err)
	}
	for _, row := range firstQueryResponse {
		fmt.Println(row)
	}

	// Get mean temperature in range(-1d, now)
	fmt.Println("\n2nd Query:")
	secondQueryResponse, err := executeQuery(dbConnection, `from (bucket: "Entry")
																	|> range(start: -1d)
																	|> filter(fn: (r) => r._field == "temperature")
																	|> keep(columns: ["_time", "_value", "sensor_id"])
																	|> mean()`)
	if err != nil {
		panic(err)
	}
	for _, row := range secondQueryResponse {
		fmt.Println(row)
	}
}

func connectToInfluxDBClient() (influxdb2.Client, error) {
	client := influxdb2.NewClient("http://localhost:8086", token)
	_, err := client.Health(context.Background())

	return client, err
}

func writeData(client influxdb2.Client, data map[string]interface{}) error {
	p := influxdb2.NewPoint("airSensors",
		map[string]string{"sensor_id": "TLM1337"},
		data,
		time.Now(),
	)

	writeAPI := client.WriteAPIBlocking(org, bucket)
	err := writeAPI.WritePoint(context.Background(), p)

	return err
}

func executeQuery(client influxdb2.Client, query string) ([]string, error) {
	queryAPI := client.QueryAPI(org)
	response, err := queryAPI.Query(context.Background(), query)
	if err != nil {
		return []string{""}, err
	}
	if response.Err() != nil {
		return []string{""}, response.Err()
	}

	result := make([]string, 0, 100)
	for response.Next() {
		result = append(result, response.Record().String())
	}

	return result, nil
}
