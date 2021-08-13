package template

var Makefile = `GOPATH:=$(shell go env GOPATH)

.PHONY: init
init:
	go get -u github.com/golang/protobuf/proto
	go get -u github.com/golang/protobuf/protoc-gen-go
	go get github.com/asim/go-micro/cmd/protoc-gen-micro/v3

.PHONY: proto
proto:
	protoc --proto_path=. --micro_out=. --go_out=:. proto/{{.Alias}}.proto

.PHONY: tidy
tidy:
	go mod tidy

.PHONY: build
build:
	go build -o {{.Alias}} *.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker:
	docker build -t {{.Alias}}:latest .
`
