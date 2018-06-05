GO_FILES      = $(shell find . -path ./vendor -prune -o -type f -name "*.go" -print)
IMPORT_PATH   = $(shell pwd | sed "s|^$(GOPATH)/src/||g")
GIT_HASH      = $(shell git rev-parse HEAD)
LDFLAGS       = -w -X $(IMPORT_PATH)/version.PreRelease=$(PRE_RELEASE)
APIB_FILES    = $(shell find . -type f -path "./*/*.apib" -not -path "./docs/*")

build: clean bindata test
	@go build -ldflags '$(LDFLAGS)'

clean:
	rm -f mdb

install:
	@godep restore

test: bindata
	go test -count=1 $(shell go list ./... | grep -v github.com/Bnei-Baruch/mdb/models)

lint:
	@golint $(GO_FILES) || true

fmt:
	@gofmt -w $(shell find . -type f -name '*.go' -not -path "./vendor/*" -not -path "./models/*")

bindata:
	@go-bindata data/... && sed -i 's/package main/package bindata/' bindata.go && mv bindata.go ./bindata

bindata_debug:
	@go-bindata -debug data/... && sed -i 's/package main/package bindata/' bindata.go && mv bindata.go ./bindata

docs:
	cd docs; \
	cp docs.tmpl docs.apib; \
	for f in ${APIB_FILES}; \
	do \
	cat ../$$f >> docs.apib; \
	done; \
	aglio -i docs.apib -o docs.html --theme-template triple
	cd ../

models:
	rm -rf models
	sqlboiler postgres
	go test ./models

.PHONY: all clean test lint fmt docs models bindata bindata_debug