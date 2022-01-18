GOPATH:=$(shell go env GOPATH)
VERSION=$(shell git describe --tags --always)

API_PROTO_FILES=$(shell find api -name *.proto)
MODEL_PATH=$(shell find . -name model -type d)
INTERNAL_PATH=$(shell find . -name internal -type d)
INTERNAL_PROTO_FILES=$(shell find . -name conf.proto)
.PHONY: init
# init env
init:
	go get -u github.com/go-kratos/kratos/cmd/kratos/v2
	go get -u google.golang.org/protobuf/cmd/protoc-gen-go
	go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc
	go get -u github.com/go-kratos/kratos/cmd/protoc-gen-go-http/v2
	go get -u github.com/go-kratos/kratos/cmd/protoc-gen-go-errors/v2

.PHONY: errors
# generate errors code
errors:
	protoc --proto_path=. \
               --proto_path=./third_party \
               --go_out=paths=source_relative:. \
               --go-errors_out=paths=source_relative:. \
               $(API_PROTO_FILES)

.PHONY: config
# generate internal proto
config:
	protoc --proto_path=. \
	       --proto_path=./third_party \
 	       --go_out=paths=source_relative:. \
	       ./app/order/service/internal/conf/conf.proto

.PHONY: api
# generate api proto
api:
	protoc --proto_path=. \
	       --proto_path=./third_party \
 	       --go_out=paths=source_relative:. \
 	       --go-http_out=paths=source_relative:. \
 	       --go-grpc_out=paths=source_relative:. \
	       $(API_PROTO_FILES)

.PHONY: build
# build
build:
	mkdir -p bin/ && go build -ldflags "-X main.Version=$(VERSION)" -o ./bin/ ./...

.PHONY: app
app:
	cd 	app/app/service && mkdir -p bin/ && \
	go build -ldflags "-X main.Version=$(VERSION)" -o ./bin/ ./... && \
	 ./bin/server -conf ./configs

.PHONY: user
user:
	cd 	app/user/service && mkdir -p bin/ && \
	go build -ldflags "-X main.Version=$(VERSION)" -o ./bin/ ./... && \
	 ./bin/server -conf ./configs
.PHONY: order
order:
	cd 	app/order/service && mkdir -p bin/ && \
	go build -ldflags "-X main.Version=$(VERSION)" -o ./bin/ ./... && \
	 ./bin/server -conf ./configs

.PHONY: spu
spu:
	cd 	app/spu/service && mkdir -p bin/ && \
	go build -ldflags "-X main.Version=$(VERSION)" -o ./bin/ ./... && \
	 ./bin/server -conf ./configs
.PHONY: generate
# generate
generate:
	go generate ./...


.PHONY: ent
ent:
	go generate ./internal/data/ent

.PHONY: validate
# validate
validate:
	protoc --proto_path=. \
           --proto_path=./third_party \
           --go_out=paths=source_relative:. \
           --validate_out=paths=source_relative,lang=go:. \
           $(API_PROTO_FILES)

.PHONY: all
# generate all
all:
	make api;
	make errors;
	make config;
	make generate;

.PHONY: run
run:
	make user;
	make order;
	make spu;
# show help
help:
	@echo ''
	@echo 'Usage:'
	@echo ' make [target]'
	@echo ''
	@echo 'Targets:'
	@awk '/^[a-zA-Z\-\_0-9]+:/ { \
	helpMessage = match(lastLine, /^# (.*)/); \
		if (helpMessage) { \
			helpCommand = substr($$1, 0, index($$1, ":")-1); \
			helpMessage = substr(lastLine, RSTART + 2, RLENGTH); \
			printf "\033[36m%-22s\033[0m %s\n", helpCommand,helpMessage; \
		} \
	} \
	{ lastLine = $$0 }' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help
