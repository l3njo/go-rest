# go-rest

[![Build Status](https://travis-ci.org/l3njo/go-rest.svg?branch=master)](https://travis-ci.org/l3njo/go-rest)
[![Go Report Card](https://goreportcard.com/badge/github.com/l3njo/go-rest)](https://goreportcard.com/report/github.com/l3njo/go-rest)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

This is a simple application that interfaces with a PostgreSQL database and exposes a REST API with GET, POST, PUT and DELETE methods to access and manipulate it.

## Configuration

The database details are fetched from the following ennvironment variables:
DBHOST - The database server host, e.g localhost
DBUSER - The database server user, e.g root
DBPORT - The port used to access the database server, e.g 5432
DBPASS - The database password, e.g 1q2w44f5g
DBNAME - The database name, e.g my_database

## Running

Simply run 'rest.exe'.
The program will be launched on localhost:8000/users.

## Building from source

### Prerequisites

You must have a recent version of the Go compiler installed. GOROOT, GOPATH and GOBIN environment variables should be set up for convenience. Visit https://www.golang.org to learn how to set up Go.

The following third-party libraries are required to build the program from the source.
1. mux
2. gorm
3. pq

### Process

Navigate to the program folder and run the ```go build``` command.
You may also use ```go install path/to/package``` to create a binary, or ```go run path/to/package``` to compile and run.
