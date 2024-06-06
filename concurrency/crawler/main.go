package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"sync"
	"time"
)

func get(url string) ([]byte, error) {
	r, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()

	data, _ := io.ReadAll(r.Body)
	return data, nil
}

func process(data []byte, name string) (*os.File, error) {
	file, err := os.OpenFile(name, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	_, err = file.Write(data)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func main() {
	t := time.Now().UTC().Format("06010215040506")
	urlTmpl := "https://github.com/search?q=go&type=repositories&p=%d"
	pages := 100

	wg := &sync.WaitGroup{}

	for i := 1; i <= pages; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()

			// Fetch Resource
			url := fmt.Sprintf(urlTmpl, i)
			d, err := get(url)
			if err != nil {
				fmt.Printf("Failed: %s (%s)\n", url, err.Error())
				return
			}

			// Handle Response
			n := fmt.Sprintf("Github-GO-%s-%03d.html", t, i)
			f, err := process(d, n)
			if err != nil {
				fmt.Printf("Failed: %s (%s)\n", url, err.Error())
				return
			}
			fmt.Printf("Success: %s -> %s\n", url, f.Name())
		}(i)
	}

	wg.Wait()
}
