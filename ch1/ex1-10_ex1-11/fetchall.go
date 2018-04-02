// fetch [URL ...]

// Observation: unresponsive URLs block the channel mechanism, see section 8.4

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	info := make(chan string)
	data := make(chan []byte)

	file, err := os.Create("data")
	defer file.Close()
	if err != nil {
		fmt.Println("error opening output file")
		os.Exit(1)
	}

	for _, url := range os.Args[1:] {
		go fetch(url, info, data)
	}

	for range os.Args[1:] {
		file.Write(<-data)
		fmt.Println(<-info)
	}

	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

// fetch writes the response time and size to info, the payload to data
func fetch(url string, info chan<- string, data chan<- []byte) {
	if !strings.HasPrefix(url, "http://") {
		url = "http://" + url
	}
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		info <- fmt.Sprint(err)
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		info <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	data <- body
	info <- fmt.Sprintf("%.2fs %7d %s", secs, len(body), url)
}
