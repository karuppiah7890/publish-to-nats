VERSION=$(shell git describe --abbrev --tags --dirty)

# TODO: Build using https://goreleaser.com/
build:
	CGO_ENABLED=0 go build -v -ldflags "-X main.version=$(VERSION)"

build-linux:
	GOARCH=amd64 GOOS=linux CGO_ENABLED=0 go build -v -ldflags "-X main.version=$(VERSION)"

release-linux: build-linux
	tar cvzf publish-to-nats-linux-amd64.tar.gz publish-to-nats

# TODO: Lint using golangci-lint
