#!/usr/bin/env bash
env GOOS=linux GOARCH=amd64 go test && go build -o tg_linux main.go shell.go