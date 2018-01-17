BINARY = server

# Symlink into GOPATH
CURRENT_DIR=$(shell pwd)
BUILD_DIR=${GOPATH}/src/github.com/jkaflik/aws-lambda-go-api-example/build

# Build the project
all: deps build

deps:
	dep ensure

handlers:
	cd ${BUILD_DIR}; \
	sh scripts/handlers.sh ; \
	cd - >/dev/null

package:
	cd ${BUILD_DIR}; \
	cd handler ; \
    zip -r ../package.zip *

deploy:
	serverless deploy

build:
	cd ${BUILD_DIR}; \
	go build ${LDFLAGS} -o ${BINARY} cmd/main.go

vet:
	-cd ${BUILD_DIR}; \
	godep go vet ./... > ${VET_REPORT} 2>&1 ; \
	cd - >/dev/null

fmt:
	cd ${BUILD_DIR}; \
	go fmt $$(go list ./... | grep -v /vendor/) ; \
	cd - >/dev/null

clean:
	-rm -f ${BUILD_DIR}/*

.PHONY: deps handlers package build clean