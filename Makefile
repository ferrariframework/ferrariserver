NAME := ferrariserver
VERSION := 0.1.0
MANTAINER := "Otto Giron <ottog2486@gmail.com>"
SOURCE_URL := http://github.com/ferrariframework/ferrariserver.git 
DATE := $(shell date -u +%Y%m%d.%H%M%S)
COMMIT_ID := $(shell git rev-parse --short HEAD)
GIT_REPO := $(shell git config --get remote.origin.url)
# Go tools flags
LD_FLAGS := -X 	github.com/ferrariframework/ferrariserver/main.buildVersion=$(VERSION)
LD_FLAGS += -X github.com/ferrariframework/ferrariserver/main.buildCommit=$(COMMIT_ID)
LD_FLAGS += -X github.com/ferrariframework/ferrariserver/main.buildDate=$(DATE)
EXTRA_BUILD_VARS := CGO_ENABLED=0 GOARCH=amd64
SOURCE_DIRS := $(shell go list ./... | grep -v /vendor/)


all: test package-linux package-darwin


test: lint generate
	 @go test -v $(SOURCE_DIRS) -cover -bench . -race 

generate: install_dependencies 
	go generate

generate_grpc:
	@echo Generating grpc assets...
	@protoc -I grpc/ grpc/job.proto --go_out=plugins=grpc:grpc/gen
	@echo Done

lint:
	@go fmt $(SOURCE_DIRS)
	@go vet $(SOURCE_DIRS)

install_dependencies: 
	glide install

binaries: binary-darwin binary-linux

binary-darwin:
	@-rm -rf build/dist/darwin
	@-mkdir -p build/dist/darwin
	GOOS=darwin $(EXTRA_BUILD_VARS) go build -ldflags "$(LD_FLAGS)" -o build/dist/darwin/$(NAME)

binary-linux:
	@-rm -rf build/dist/linux
	@-mkdir -p build/dist/linux
	GOOS=linux $(EXTRA_BUILD_VARS) go build -ldflags "$(LD_FLAGS)" -o build/dist/linux/$(NAME)


package-darwin: binary-darwin
	@tar -czf build/dist/ferrariserver.darwin-amd64.tar.gz -C build/dist/darwin ferrariserver


package-linux: binary-linux
	@tar -czf build/dist/ferrariserver.linux-amd64.tar.gz -C build/dist/linux ferrariserver