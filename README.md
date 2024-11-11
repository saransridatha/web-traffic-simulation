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
