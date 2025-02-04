default: testacc 
GOOS=$(shell go env GOOS)
GOARCH=$(shell go env GOARCH)
INSTALL_PATH=~/.local/share/terraform/plugins/localhost/providers/boundary/0.0.1/linux_$(GOARCH)
BUILD_ALL_PATH=${PWD}/bin

ifeq ($(GOOS), darwin)
	INSTALL_PATH=~/Library/Application\ Support/io.terraform/plugins/localhost/providers/boundary/0.0.1/darwin_$(GOARCH)
endif
ifeq ($(GOOS), "windows")
	INSTALL_PATH=%APPDATA%/HashiCorp/Terraform/plugins/localhost/providers/boundary/0.0.1/windows_$(GOARCH)
endif

tools:
	go generate -tags tools tools/tools.go

fmtcheck:
	echo "Placeholder"

test:
	echo "Placeholder"

# Run acceptance tests
testacc:
	TF_ACC=1 go test ./... -v $(TESTARGS) -timeout 120m

testacc-ci: install-go
	git config --global --add url."git@github.com:".insteadOf "https://github.com/"
	TF_ACC=1 ~/.go/bin/go test ./... -v $(TESTARGS) -timeout 120m

install-go:
	./ci/goinstall.sh

dev:
	mkdir -p $(INSTALL_PATH)	
	go build -o $(INSTALL_PATH)/terraform-provider-boundary main.go

all:
	mkdir -p $(BUILD_ALL_PATH)
	GOOS=darwin go build -o $(BUILD_ALL_PATH)/terraform-provider-boundary_darwin-amd64 main.go
	GOOS=windows go build -o $(BUILD_ALL_PATH)/terraform-provider-boundary_windows-amd64 main.go
	GOOS=linux go build -o $(BUILD_ALL_PATH)/terraform-provider-boundary_linux-amd64 main.go

docs: 
	go run github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs

rm-id-flag-from-docs:
	find docs/ -name "*.md" -type f | xargs sed -i -e '/- \*\*id\*\*/d'

.PHONY: testacc tools docs
