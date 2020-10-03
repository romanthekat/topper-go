#-------------------------
# Build artefacts
#-------------------------
.PHONY: build build.linux build.mac

## Build binary
build:
	@go build -o topper-go main.go

## Build linux binary
build.linux: 
	@env GOOS=linux GOARCH=amd64 go test && go build -o tg_linux main.go
	

## Build macos binary
build.mac:
	@env GOOS=darwin GOARCH=amd64 go test && go build -o tg_mac main.go
