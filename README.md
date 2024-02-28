# SSRFuzzer
<p align="center">
  <img src="temple.jpeg">
</p>

## Introduction

This project implements a port scanning fuzzer utilizing SSRF (Server-Side Request Forgery) vulnerabilities. It's designed to probe internal server ports of a web application's hosting environment, thereby helping to identify improperly exposed services and other potential vulnerabilities linked to internal network configurations.

## How It Works

The fuzzer makes POST requests to a specified URL with payloads that attempt to access internal services through specific ports on the target host. The code supports concurrent execution of multiple requests to enhance scanning efficiency, using a configurable thread system.

### Features

- Scans ports from 21 to 65535 (range can be changed)
- Utilizes SSRF to test the accessibility of internal ports.
- Concurrent request execution using goroutines.
- Configurable number of threads to control concurrency.
- Response status detection to identify accessible ports.

## Installation

To use this script, you need to have [Go](https://golang.org/) installed on your system. Follow the instructions below to set up and run the fuzzer.

1. Clone the repository to your local machine:

```
git clone https://github.com/L3ster1337/SSRFuzzer.git
```
2. Set the thread number and modify the request accordingly to your SSRF vuln payload structure
4. Compile the Code
```
go build ssrfuzzer.go
```
5. Run it
```
./ssrfuzzer
```
