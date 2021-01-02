#!/bin/sh

go build cmd/response-server/response-server.go
./response-server $@