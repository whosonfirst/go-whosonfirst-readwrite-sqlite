CWD=$(shell pwd)
GOPATH := $(CWD)

prep:
	if test -d pkg; then rm -rf pkg; fi

self:   prep rmdeps
	if test -d src; then rm -rf src; fi
	mkdir -p src/github.com/whosonfirst/go-whosonfirst-readwrite-sqlite
	cp -r database src/github.com/whosonfirst/go-whosonfirst-readwrite-sqlite/
	cp -r reader src/github.com/whosonfirst/go-whosonfirst-readwrite-sqlite/
	cp -r writer src/github.com/whosonfirst/go-whosonfirst-readwrite-sqlite/
	cp -r vendor/* src/

rmdeps:
	if test -d src; then rm -rf src; fi 

build:	fmt bin

docker-build:
	docker build -t wof-readwrited .

deps:
	@GOPATH=$(GOPATH) go get -u "github.com/whosonfirst/go-whosonfirst-readwrite/..."
	@GOPATH=$(GOPATH) go get -u "github.com/whosonfirst/go-whosonfirst-sqlite"
	@GOPATH=$(GOPATH) go get -u "github.com/whosonfirst/go-whosonfirst-sqlite-features"
	@GOPATH=$(GOPATH) go get -u "github.com/whosonfirst/go-whosonfirst-uri"
	rm -rf src/github.com/whosonfirst/go-whosonfirst-sqlite-features/vendor/github.com/whosonfirst/go-whosonfirst-sqlite
	cp -r http src/github.com/whosonfirst/go-whosonfist-readwrite/

vendor-deps: rmdeps deps
	if test ! -d vendor; then mkdir vendor; fi
	if test -d vendor; then rm -rf vendor; fi
	cp -r src vendor
	find vendor -name '.git' -print -type d -exec rm -rf {} +
	rm -rf src

fmt:
	go fmt database/*.go
	go fmt http/*.go
	go fmt reader/*.go
	go fmt writer/*.go

bin: 	self
	GOPATH=$(GOPATH) go build -o bin/wof-readerd cmd/wof-readerd.go
