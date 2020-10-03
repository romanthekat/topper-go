

#-------------------------
# Build artefacts
#-------------------------
.PHONY: build build.linux build.mac

## Build all binaries
build:
	@go build -o topper-go main.go

## Execute development pipeline
build.linux: 
	@env GOOS=linux GOARCH=amd64 go test && go build -o tg_linux main.go
	

## Execute production pipeline
build.mac:
	@env GOOS=darwin GOARCH=amd64 go test && go build -o tg_mac main.go
