package main

import (
	"bufio"
	"encoding/json"
	"io/ioutil"
	"os"
)

type MyStruct struct {
    Students []struct {
        Rating []float64
    }
}

type Result struct {
    Average float64
}

func main() {
	data, err := ioutil.ReadAll(os.Stdin)
    if err != nil {
        panic(err)
    }

    var myStruct MyStruct
    json.Unmarshal(data, &myStruct)

    var marksCount, studentsCount float64
    for _, arr := range myStruct.Students {
        marksCount += float64(len(arr.Rating))
        studentsCount++
    }

    result := Result{
        Average: marksCount / studentsCount,
    }
    jsonData, _ := json.MarshalIndent(&result, "", "\t")
    wr := bufio.NewWriter(os.Stdout)
    wr.Write(jsonData)
    wr.Flush()
}