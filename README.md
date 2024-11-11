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

        URL: The target URL you want to test.
        Total Number of Requests: The number of HTTP GET requests to send.
        Concurrent Requests: The number of requests to send concurrently.   
















   
