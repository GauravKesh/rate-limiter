package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

const (
	URL            = "http://127.0.0.1:3000/"
	TotalRequests  = 200
	Concurrency    = 20
	RequestTimeout = 5 * time.Second
)

type Result struct {
	status  int
	Error   error
	Headers http.Header
}

func worker(wg *sync.WaitGroup, jobs <-chan int, results chan<- Result) {
	defer wg.Done()

	client := &http.Client{
		Timeout: RequestTimeout,
	}

	for range jobs {
		resp, err := client.Get(URL)
		if err != nil {
			results <- Result{Error: err}
			continue
		}

		results <- Result{status: resp.StatusCode,
			Headers: resp.Header,
		}
		resp.Body.Close()
	}
}

func main() {
	start := time.Now()

	jobs := make(chan int, TotalRequests)
	results := make(chan Result, TotalRequests)

	var wg sync.WaitGroup

	//spawn workers

	for i := 0; i < Concurrency; i++ {
		wg.Add(1)
		go worker(&wg, jobs, results)
	}

	// send jobs

	for i := 0; i < TotalRequests; i++ {
		jobs <- i
	}
	close(jobs)

	wg.Wait()
	close(results)

	success := 0
	rateLimited := 0
	failed := 0

	for res := range results {
		if res.Error != nil {
			fmt.Println("ERROR:", res.Error)
			failed++
			continue
		}
		switch res.status {
		case http.StatusOK:
			success++
		case http.StatusTooManyRequests:
			fmt.Println(res.Headers.Get("X-RateLimit-Limit"))
			fmt.Println(res.Headers.Get("X-RateLimit-Remaining"))
			rateLimited++
		default:
			failed++

		}
	}
	duration := time.Since(start)
	fmt.Println("------ Load Test Result ------")
	fmt.Printf("Total Requests: %d\n", TotalRequests)
	fmt.Printf("Success (200): %d\n", success)
	fmt.Printf("Rate Limited (429): %d\n", rateLimited)
	fmt.Printf("Failed: %d\n", failed)
	fmt.Printf("Time Taken: %v\n", duration)
	fmt.Printf("Throughput: %.2f req/sec\n", float64(TotalRequests)/duration.Seconds())
}
