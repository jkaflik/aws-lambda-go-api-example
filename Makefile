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
	rm -f handler/* ; \
	sh scripts/handlers.sh

package:
	cd ${BUILD_DIR}; \
	rm -f package.zip ;\
	cd handler ; \
    zip -r ../package.zip *

deploy:
	serverless deploy

build:
	cd ${BUILD_DIR}; \
	go build ${LDFLAGS} -o ${BINARY} cmd/main.go

clean:
	-rm -rf ${BUILD_DIR}/handler ${BUILD_DIR}/package.zip

.PHONY: deps handlers package build clean