package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

func worker(id int, url string) {
	request(id, url)
}

func main() {

	start := time.Now()
	url := "http://kpakozz96pyc.xyz:8080/product/getall"
	iterations := 1000
	//threads:= 10; //ToDo - limit threads

	fmt.Printf("Test started %.s \n requested url %s \n iterations: %d \n ", start.Format("15:04:05.100"), url, iterations)

	var wg sync.WaitGroup

	for i := 1; i <= iterations; i++ {
		wg.Add(1)

		i := i

		go func() {
			defer wg.Done()
			worker(i, url)
		}()
	}

	wg.Wait()
	fmt.Println("\033[32m", fmt.Sprintf("Test Ended in %.3fs, %d Requests done", time.Since(start).Seconds(), iterations))

}

func request(i int, url string) {
	start := time.Now()
	resp, _ := http.Get(url)
	_, _ = ioutil.ReadAll(resp.Body)
	fmt.Printf("Request %d  %s  %.3fs elapsed\n", i, url, time.Since(start).Seconds())
}
