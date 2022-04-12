package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type MyJsonFile struct {
    ID int64 `json:"global_id"`
}

func main() {
	var myJsonFile []MyJsonFile

    file, err := os.Open("data-20190514T0100.json")
    if err != nil {
        panic(err)
    }
    defer file.Close()
    
    data, _ := ioutil.ReadAll(file)
    json.Unmarshal(data, &myJsonFile)

    var sum int64
    for _, value := range myJsonFile {
        sum += value.ID
    }

    fmt.Println(sum)
}