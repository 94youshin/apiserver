SHELL := /bin/bash
BASEDIR = $(shell pwd)
export CGO_ENABLED=0

versionDir = "github.com/youshintop/apiserver/pkg/version"
gitTag = $(shell if [ "`git describe --tags --abbrev=0 2>/dev/null`" != "" ];then git describe --tags --abbrev=0; else git log --pretty=format:'%h' -n 1; fi)
buildDate = $(shell TZ=Asia/Shanghai date +%FT%T%z)
gitCommit = $(shell git log --pretty=format:'%H' -n 1)
gitTreeState = $(shell if git status|grep -q 'clean';then echo clean; else echo dirty; fi)

ldflags="-w -X ${versionDir}.gitTag=${gitTag} -X ${versionDir}.buildDate=${buildDate} -X ${versionDir}.gitCommit=${gitCommit} -X ${versionDir}.gitTreeState=${gitTreeState}"

all: fmt vet build

.PHONY: build
build:
	@go build -v -ldflags ${ldflags} .

cert:
	openssl req -new -nodes -x509 -out conf/server.crt -keyout conf/server.key -days 3650 -subj "/C=DE/ST=NRW/L=Earth/O=Random Company/OU=IT/CN=127.0.0.1,10.92.119.242/emailAddress=yanginbot@gmail.com"

clean:
	@rm -f apiserver
fmt:
	go fmt ./...
vet:
	go vet ./...
run:
	./apiserver --config=./conf/config.yaml
