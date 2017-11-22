# health-check

[![Build Status](https://travis-ci.org/robertoduessmann/health-check.svg?branch=master)](https://travis-ci.org/robertoduessmann/health-check)
[![Go Report Card](https://goreportcard.com/badge/github.com/robertoduessmann/health-check)](https://goreportcard.com/report/github.com/robertoduessmann/health-check)
[![GoDoc](https://godoc.org/github.com/robertoduessmann/health-check?status.svg)](https://godoc.org/github.com/robertoduessmann/health-check)

> Check your server health easily

## Build
```sh
go build
```
## Run
Build (above) or download [health-check](https://github.com/robertoduessmann/health-check/blob/master/health-check) and run:
```sh
./health-check
```

## Usage

Just open http://localhost:3000/status in your server :)

## Response example
```json
{
	"memory": {
		"total": "12349032 Kb",
		"used": "3490432 Kb",
		"free": "8858600 Kb"
	},
	"disks": [{
		"name": "sda1",
		"total": "1623248 Kb",
		"used": "643480 Kb",
		"free": "979768 Kb"
	}, {
		"name": "sda2",
		"total": "5483946 Kb",
		"used": "5235354 Kb",
		"free": "248592 Kb"
	}],
	"processor": {
		"total": "8.09 %"
	}
}
```

## Installation

If you want to hack on health-check

`go get github.com/robertoduessmann/health-check`

To make sure all the dependencies are installed

`dep ensure`
