package main

import (
	"archive/zip"
	"bufio"
	"fmt"
	"strings"
)

func main() {
	archive, err := zip.OpenReader("test.zip")

    if err != nil {
        panic(err)
    }

    for _, file := range archive.File {
        // Skip directories
        if file.FileInfo().IsDir() {
            continue
        }

        // Read .csv file
        fileInArchive, err := file.Open()
        if err != nil {
            panic(err)
        }
        defer fileInArchive.Close()

        ioScanner := bufio.NewScanner(fileInArchive)

        row := 0
        for ioScanner.Scan() {
            line := ioScanner.Text()

            // Search 5th row (Index = 4)
            if row != 4 {
                row++
                continue
            } else {
                fmt.Println(strings.Split(line, ",")[2])
                break
            }
        }

    }
}