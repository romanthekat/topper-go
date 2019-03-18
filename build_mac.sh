#!/usr/bin/env bash
env GOOS=darwin GOARCH=amd64 go test && go build -o tg_mac main.go shell.go