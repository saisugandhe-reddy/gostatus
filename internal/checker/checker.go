package checker

import (
	"net/http"
	"sync"
	"time"
)

type Result struct {
	URL        string
	StatusCode int
	Latency    time.Duration
	Error      string
}

func CheckURLs(urls []string, timeout time.Duration) []Result {
	results := make([]Result, len(urls))
	client := &http.Client{Timeout: timeout}

	var wg sync.WaitGroup
	for i, url := range urls {
		wg.Add(1)
		go func(i int, url string) {
			defer wg.Done()
			results[i] = checkURL(client, url)
		}(i, url)
	}

	wg.Wait()
	return results
}

func checkURL(client *http.Client, url string) Result {
	start := time.Now()
	resp, err := client.Get(url)
	latency := time.Since(start)

	if err != nil {
		return Result{
			URL:     url,
			Latency: latency,
			Error:   err.Error(),
		}
	}
	defer resp.Body.Close()

	return Result{
		URL:        url,
		StatusCode: resp.StatusCode,
		Latency:    latency,
	}
}
