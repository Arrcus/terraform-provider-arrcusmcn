TEST?=$$(go list ./... | grep -v 'vendor')
HOSTNAME=arrcus.com
NAMESPACE=arrcus
NAME=arrcusmcn
VERSION=1.0.0
OS_ARCH=darwin_arm64
OS=$(shell go env GOOS)
ARCH=$(shell go env GOARCH)
BINARY=terraform-provider-${NAME}
default: install

build:
	go build -o ${BINARY}

release:
	goreleaser release --rm-dist --snapshot --skip-publish  --skip-sign

install: build
	mkdir -p ~/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/${OS}_${ARCH}
	mv ${BINARY} ~/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/${OS}_${ARCH}

test: 
	go test -i $(TEST) || exit 1                                                   
	echo $(TEST) | xargs -t -n4 go test $(TESTARGS) -timeout=30s -parallel=4                    

testacc: 
	TF_ACC=1 go test $(TEST) -v $(TESTARGS) -timeout 120m   