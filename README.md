# Web Traffic Simulation
This project is a Go-based tool for simulating high levels of web traffic to test how your web infrastructure handles a large number of concurrent HTTP requests. It provides real-time feedback on success/failure rates, response times, and other performance metrics during the load test. The tool is useful for stress testing your website or web application.

## Features
1. Simulate a large number of HTTP GET requests to a target URL.
2. Control the number of total requests and the number of concurrent requests.
3. Real-time monitoring of request progress, including:
     1. Requests sent.
     2. Success and failure counts.
     3. Average, 90th percentile, and 9th percentile response times.
4. Post-test summary including:
     1. Total requests.
     2. Success rate.
     3. Failure rate.
     4. Response time metrics.
     5. Status code distribution.
        
## Prerequisites
1. Go 1.18+ installed on your machine.
2. A terminal (Linux, macOS, or Windows with WSL).

## Installation
1. Clone the repository to your local machine:
   ```bash
   git clone https://github.com/saransridatha/web-traffic-simulation.git
   ```
2. Navigate to the project directory:
   ```bash
   cd web-traffic-simulation
   ```
3. Build the Go program (optional, since go run can be used directly):
   ```bash
   go build script.go
   ```
## Usage
1. To run the web traffic simulation tool, follow these steps:
   ```bash
   go run script.go
   ```
2. The program will prompt you for the following inputs:

      1. **URL:** The target URL you want to test.
      2. **Total Number of Requests:** The number of HTTP GET requests to send.
      3. **Concurrent Requests:** The number of requests to send concurrently.
3. The program will start sending requests and show real-time metrics, including the success rate, response time, and percentiles.
4. Once all requests have been completed, the program will display a summary, including:

      1. Total requests sent.
      2. Successful and failed requests.
      3. Average response time, 90th and 99th percentiles.
      4. Status code distribution.

## Exaple Output
```bash
Enter the URL to test: https://example.com
Enter the total number of requests to send: 1000
Enter the number of concurrent requests to send: 50

| Requests Sent: 500/1000 | Success: 450 | Failures: 50 | Avg Resp Time: 0.3234s | 90th Pctl: 0.7453 | 99th Pctl: 1.0234s | Elapsed Time: 10.43s
```
After all requests are completed:
```bash
Load Test Summary
==================
Total Requests Sent      : 1000
Successful Requests      : 950
Failed Requests          : 50
Success Rate             : 95.00%
Failure Rate             : 5.00%
Total Time Taken         : 15.23 seconds
Requests per Second      : 65.68

Response Time Metrics (seconds)
--------------------------------
Average Response Time    : 0.3125
90th Percentile          : 0.7453
99th Percentile          : 1.0234
Minimum Response Time    : 0.1234
Maximum Response Time    : 2.3456

Status Code Distribution
------------------------
Status Code 200 : 950 responses
Status Code 404 : 50 responses
```

## Disclaimer

**This tool is intended for legitimate testing of your own web infrastructure or websites that you have explicit permission to test. It should only be used for performance testing, load testing, and stress testing purposes on servers and websites that you own or have permission to test.**

**Do not use this tool to perform Denial of Service (DoS) attacks or Distributed Denial of Service (DDoS) attacks on any website or server without the owner's consent. Unauthorized use of this tool for DoS or DDoS attacks is illegal and unethical, and can result in serious legal consequences.**











   
