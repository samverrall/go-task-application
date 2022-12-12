#!/bin/bash

# exit when any command fails
set -e

echo Running golangci-lint
golangci-lint --enable gosec,misspell run ./...

echo Running go test
go test --cover --race --count=1 ./...

echo Building with race condition checking enabled.
go build --race -o ./tasker-api github.com/samverrall/task-service/cmd
