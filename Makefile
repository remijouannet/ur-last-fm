TEST?=$$(go list ./... |grep -v 'vendor')
GOFMT_FILES?=$$(find . -name '*.go' |grep -v vendor)
XC_ARCH="amd64 arm"
XC_OS="linux darwin windows freebsd"
XC_EXCLUDE_OSARCH="!darwin/arm !darwin/386"
VERSION=$$(git describe --abbrev=0 --tags)
PWD=$$(pwd)

COMMIT=$$(git rev-parse HEAD)
GOOS=$$(go env GOOS)
GOARCH=$$(go env GOARCH)

default: build

build: fmt
	go install

pkg: fmt
	mkdir -p ./pkg
	rm -rf ./pkg/*
	echo "==> Building..."
	CGO_ENABLED=0 gox -os=$(XC_OS) -arch=$(XC_ARCH) \
				-osarch=$(XC_EXCLUDE_OSARCH) \
				-output ./pkg/ur-last-fm_{{.OS}}_{{.Arch}}_$(VERSION)/ur-last-fm_$(VERSION) .

bin: fmt
	mkdir -p ./bin
	echo "==> Building..."
	CGO_ENABLED=0 go build -o ./bin/ur-last-fm_$(VERSION) .
	chmod 777 ./bin/ur-last-fm_$(VERSION)

bin-darwin: fmt
	mkdir -p ./bin
	echo "==> Building..."
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o ./bin/ur-last-fm_$(VERSION) .
	chmod 777 ./bin/ur-last-fm_$(VERSION)

fmt:
	gofmt -w $(GOFMT_FILES)

release:
	bash scripts/github-releases.sh

docker-bin: docker-image
	docker run  \
		-v $(PWD)/bin:/go/src/github.com/remijouannet/ur-last-fm/bin \
		ur-last-fm:$(VERSION) bin

docker-bin-darwin: docker-image
	docker run  \
		-v $(PWD)/bin:/go/src/github.com/remijouannet/ur-last-fm/bin \
		ur-last-fm:$(VERSION) bin-darwin

docker-image:
	docker build -t ur-last-fm:$(VERSION) .

docker-build:
	docker run \
		-v $(PWD)/pkg:/go/src/github.com/remijouannet/ur-last-fm/pkg \
		ur-last-fm:$(VERSION) pkg

docker-release:
	docker run \
		-v $(PWD)/pkg:/go/src/github.com/remijouannet/ur-last-fm/pkg \
		-e "GITHUB_TOKEN=$(GITHUB_TOKEN)" \
		ur-last-fm:$(VERSION) release

.PHONY: build test testacc vet fmt fmtcheck errcheck vendor-status test-compile
