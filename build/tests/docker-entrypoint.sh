#!/bin/bash

set -xe

# Testing Go client.
gotenberg &
sleep 10
go test -race -cover -covermode=atomic github.com/meateam/gotenberg-go-client/v6
sleep 10 # allows Gotenberg to remove generated files.