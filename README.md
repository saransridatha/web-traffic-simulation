# Web Traffic Simulation
This project is a Go-based tool for simulating high levels of web traffic to test how your web infrastructure handles a large number of concurrent HTTP requests. It provides real-time feedback on success/failure rates, response times, and other performance metrics during the load test. The tool is useful for stress testing your website or web application.

## Features
1. Simulate a large number of HTTP GET requests to a target URL.
2. Control the number of total requests and the number of concurrent requests.
3. Real-time monitoring of request progress, including:
      Requests sent.
      Success and failure counts.
      Average, 90th percentile, and 99th percentile response times.
4. Post-test summary including:
      Total requests.
      Success rate.
      Failure rate.
      Response time metrics.
      Status code distribution.
