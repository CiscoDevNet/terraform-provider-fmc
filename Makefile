TEST?=$$(go list ./... | grep -v 'vendor')
HOSTNAME=registry.terraform.io
NAMESPACE=CiscoDevNet
NAME=fmc
BINARY=terraform-provider-${NAME}
VERSION=0.2
OS_ARCH=$(shell go env GOOS)_$(shell go env GOARCH)

default: install

build:
	go build -o ${BINARY}_${VERSION}_${OS_ARCH}

PLATFORMS := linux/amd64 windows/amd64 darwin/amd64

temp = $(subst /, ,$@)
OS = $(word 1, $(temp))
ARCH = $(word 2, $(temp))

release: $(PLATFORMS)
	tar cvzf release.tgz release
	zip -r release.zip release

$(PLATFORMS):
	mkdir -p ./release/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/${OS}_${ARCH}
	GOOS=${OS} GOARCH=${ARCH} go build -o ./release/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/${OS}_${ARCH}/${BINARY}_${VERSION}_${OS}_${ARCH}

.PHONY: release $(PLATFORMS)

install: build
	mkdir -p ~/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/${OS_ARCH}
	mv ${BINARY}_${VERSION}_${OS_ARCH} ~/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/${OS_ARCH}

generate:
	tfplugindocs

test: 
	go test -i $(TEST) || exit 1                                                   
	echo $(TEST) | xargs -t -n4 go test $(TESTARGS) -timeout=30s -parallel=4                    

testacc: 
	TF_ACC=1 go test $(TEST) -v $(TESTARGS) -timeout 120m

testinfra-testacc:
	FMC_INSECURE_SKIP_VERIFY=true \
	FMC_USERNAME="$$(cd testinfra; terraform output --raw fmc_username)" \
	FMC_PASSWORD="$$(cd testinfra; terraform output --raw fmc_password)" \
	FMC_HOST="$$(cd testinfra; terraform output --raw fmc_ip)" \
	$(info $$FMC_HOST is [${FMC_HOST}]) \
	TF_ACC=1 go test $(TEST) -v $(TESTARGS) -timeout 120m

