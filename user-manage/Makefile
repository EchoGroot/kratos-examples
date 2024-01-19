GOHOSTOS:=$(shell go env GOHOSTOS)
GOPATH:=$(shell go env GOPATH)
VERSION=$(shell git describe --tags --always)
CMD_PATH := $(shell pwd)/cmd
WIRE_PATH := $(shell find $(CMD_PATH) -maxdepth 1 -type d)
WIRE_PATH := $(filter-out $(CMD_PATH), $(WIRE_PATH))
KRATOS_VERSION=$(shell go mod graph |grep go-kratos/kratos/v2 |head -n 1 |awk -F '@' '{print $$2}')
KRATOS=$(GOPATH)/pkg/mod/github.com/go-kratos/kratos/v2@$(KRATOS_VERSION)

ifeq ($(GOHOSTOS), windows)
	#the `find.exe` is different from `find` in bash/shell.
	#to see https://docs.microsoft.com/en-us/windows-server/administration/windows-commands/find.
	#changed to use git-bash.exe to run find cli or other cli friendly, caused of every developer has a Git.
	Git_Bash= $(subst cmd\,bin\bash.exe,$(dir $(shell where git)))
	INTERNAL_PROTO_FILES=$(shell $(Git_Bash) -c "find internal -name *.proto")
	API_PROTO_FILES=$(shell $(Git_Bash) -c "find api -name *.proto")
else
	INTERNAL_PROTO_FILES=$(shell find internal -name *.proto)
	API_PROTO_FILES=$(shell find api -name *.proto)
endif

# 帮助信息
help:
	@echo ''
	@echo '用法:'
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


.PHONY: init
# 初始化依赖，clone项目后，需要运行一次
init:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	go install github.com/go-kratos/kratos/cmd/kratos/v2@latest
	go install github.com/go-kratos/kratos/cmd/protoc-gen-go-http/v2@latest
	go install github.com/google/gnostic/cmd/protoc-gen-openapi@latest
	go install github.com/google/wire/cmd/wire@latest
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	go install mvdan.cc/gofumpt@latest
	go mod tidy


.PHONY: conf
# 生成配置文件相关代码
conf:
	protoc --proto_path=./internal \
	       --proto_path=$(KRATOS)/third_party \
 	       --go_out=paths=source_relative:./internal \
           --validate_out=paths=source_relative,lang=go:./internal \
	       $(INTERNAL_PROTO_FILES)

.PHONY: api
# 生成API相关代码
api:
	protoc --proto_path=. \
	       --proto_path=$(KRATOS)/third_party \
 	       --go_out=paths=source_relative:. \
 	       --go-http_out=paths=source_relative:. \
 	       --go-grpc_out=paths=source_relative,require_unimplemented_servers=false:. \
           --go-errors_out=paths=source_relative:. \
           --validate_out=paths=source_relative,lang=go:. \
           --openapiv2_out=allow_repeated_fields_in_body=true:. \
	       $(API_PROTO_FILES)

.PHONY: lint
# 代码检查并修复部分问题，规则文件 ./.golangci.yml
lint:
	gofumpt -e -d -w ./
	golangci-lint run -v


.PHONY: wire
# 生成依赖注入代码
wire:
	@for dir in ${WIRE_PATH}; do cd $$dir && wire && cd - > /dev/null; done


.PHONY: test
# 单元测试
test:
	go test -short `go list ./...`


.PHONY: updateGoDependencies
# 更新 go 依赖
updateGoDependencies:
	go mod tidy

.PHONY: build
# 构建二进制可执行文件
build:
	go mod tidy
	mkdir -p bin
	cd cmd/${SUB_MODULE};CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-X main.Version=${VERSION}" -o ../../bin/

.PHONY: buildImage
# 构建镜像
buildImage:
	docker build --build-arg SUB_MODULE=${SUB_MODULE} -t ${MODULE_PREFIX}-${SUB_MODULE}:${VERSION} -f Dockerfile .;




