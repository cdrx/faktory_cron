NAME=faktory_worker
VERSION=0.5.0

deps:
	@go get github.com/golang/dep/cmd/dep
	@dep ensure
	@go get github.com/goreleaser/goreleaser

build: clean
	go build -o faktory-cron *.go

run: build
	./faktory-cron

clean:
	-rm -rf faktory-cron

fmt:
	go fmt ./...
