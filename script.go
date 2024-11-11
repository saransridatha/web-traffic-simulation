package main

import (
	"fmt"
	"net/http"
	"sort"
	"sync"
	"time"
)

var (
	successCount    int
	failureCount    int
	responseTimes   []float64
	statusCodeCount = make(map[int]int)
	errorLog        []string
	lock            sync.Mutex
)

func main() {
	var url string
	var numberOfRequests, concurrentRequests int

	fmt.Print("Enter the URL to test: ")
	fmt.Scan(&url)
	fmt.Print("Enter the total number of requests to send: ")
	fmt.Scan(&numberOfRequests)
	fmt.Print("Enter the number of concurrent requests to send: ")
	fmt.Scan(&concurrentRequests)

	startTime := time.Now()

	go printLiveMetrics(startTime, numberOfRequests)

	var wg sync.WaitGroup
	wg.Add(numberOfRequests)
	for i := 0; i < numberOfRequests; i++ {
		go sendRequest(url, &wg)
	}

	wg.Wait()

	printSummary(startTime, numberOfRequests)
}

func sendRequest(url string, wg *sync.WaitGroup) {
	defer wg.Done()

	startTime := time.Now()
	resp, err := http.Get(url)
	elapsedTime := time.Since(startTime).Seconds()

	lock.Lock()
	defer lock.Unlock()

	responseTimes = append(responseTimes, elapsedTime)

	if err != nil {
		failureCount++
		statusCodeCount[0]++
		errorLog = append(errorLog, err.Error())
	} else {
		defer resp.Body.Close()
		if resp.StatusCode >= 200 && resp.StatusCode < 300 {
			successCount++
		} else {
			failureCount++
		}
		statusCodeCount[resp.StatusCode]++
	}
}

func printLiveMetrics(startTime time.Time, totalRequests int) {
	spinner := []string{"|", "/", "-", "\\"}
	spinIdx := 0
	for successCount+failureCount < totalRequests {
		lock.Lock()
		elapsedTime := time.Since(startTime).Seconds()

		avgResponseTime := "N/A"
		if len(responseTimes) > 0 {
			avgResponseTime = fmt.Sprintf("%.4f", average(responseTimes))
		}
		percentile90 := "N/A"
		if len(responseTimes) > 0 {
			percentile90 = fmt.Sprintf("%.4f", percentile(responseTimes, 90))
		}
		percentile99 := "N/A"
		if len(responseTimes) > 0 {
			percentile99 = fmt.Sprintf("%.4f", percentile(responseTimes, 99))
		}

		fmt.Printf("\r%s Requests Sent: %d/%d | Success: %d | Failures: %d | Avg Resp Time: %s | 90th Pctl: %s | 99th Pctl: %s | Elapsed Time: %.2fs",
			spinner[spinIdx], successCount+failureCount, totalRequests, successCount, failureCount, avgResponseTime, percentile90, percentile99, elapsedTime)
		lock.Unlock()

		spinIdx = (spinIdx + 1) % len(spinner)
		time.Sleep(100 * time.Millisecond)
	}
}

func printSummary(startTime time.Time, totalRequests int) {
	elapsedTime := time.Since(startTime).Seconds()

	avgResponseTime := "N/A"
	if len(responseTimes) > 0 {
		avgResponseTime = fmt.Sprintf("%.4f", average(responseTimes))
	}
	percentile90 := "N/A"
	if len(responseTimes) > 0 {
		percentile90 = fmt.Sprintf("%.4f", percentile(responseTimes, 90))
	}
	percentile99 := "N/A"
	if len(responseTimes) > 0 {
		percentile99 = fmt.Sprintf("%.4f", percentile(responseTimes, 99))
	}
	successRate := (float64(successCount) / float64(totalRequests)) * 100
	failureRate := (float64(failureCount) / float64(totalRequests)) * 100

	fmt.Printf("\n\nLoad Test Summary\n==================\n")
	fmt.Printf("Total Requests Sent      : %d\n", totalRequests)
	fmt.Printf("Successful Requests      : %d\n", successCount)
	fmt.Printf("Failed Requests          : %d\n", failureCount)
	fmt.Printf("Success Rate             : %.2f%%\n", successRate)
	fmt.Printf("Failure Rate             : %.2f%%\n", failureRate)
	fmt.Printf("Total Time Taken         : %.2f seconds\n", elapsedTime)
	fmt.Printf("Requests per Second      : %.2f\n", float64(totalRequests)/elapsedTime)

	fmt.Printf("\nResponse Time Metrics (seconds)\n--------------------------------\n")
	fmt.Printf("Average Response Time    : %s\n", avgResponseTime)
	fmt.Printf("90th Percentile          : %s\n", percentile90)
	fmt.Printf("99th Percentile          : %s\n", percentile99)
	fmt.Printf("Minimum Response Time    : %.4f\n", min(responseTimes))
	fmt.Printf("Maximum Response Time    : %.4f\n", max(responseTimes))

	fmt.Printf("\nStatus Code Distribution\n------------------------\n")
	for code, count := range statusCodeCount {
		fmt.Printf("Status Code %d : %d responses\n", code, count)
	}

	if failureCount > 0 {
		fmt.Println("\nError Summary\n-------------")
		for i, err := range errorLog[:10] {
			fmt.Printf("%d. %s\n", i+1, err)
		}
		if len(errorLog) > 10 {
			fmt.Printf("...and %d more errors.\n", len(errorLog)-10)
		}
	}
}

func average(data []float64) float64 {
	if len(data) == 0 {
		return 0
	}
	sum := 0.0
	for _, value := range data {
		sum += value
	}
	return sum / float64(len(data))
}

func min(data []float64) float64 {
	if len(data) == 0 {
		return 0
	}
	minVal := data[0]
	for _, value := range data[1:] {
		if value < minVal {
			minVal = value
		}
	}
	return minVal
}

func max(data []float64) float64 {
	if len(data) == 0 {
		return 0
	}
	maxVal := data[0]
	for _, value := range data[1:] {
		if value > maxVal {
			maxVal = value
		}
	}
	return maxVal
}

func percentile(data []float64, p int) float64 {
	if len(data) == 0 {
		return 0
	}
	sort.Float64s(data)
	index := int(float64(len(data)) * float64(p) / 100)
	if index >= len(data) {
		index = len(data) - 1
	}
	return data[index]
}
