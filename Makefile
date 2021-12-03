# Default variable
GOIMPORTS ?= goimports
GOCILINT ?= golangci-lint
APP_NAME ?= go-cli-prototype

default:
	@go run cmd/main.go

build-all: build-darwin build-linux build-windows

build-darwin:
	-rm -rf ./build/darwin/bin
	# Update version
	cd pkg/version/scripts && go run gen.go
	# Build kusion
	GOOS=darwin GOARCH=amd64 CGO_ENABLED=0 \
		go build -o ./build/darwin/bin/$(APP_NAME) \
		./cmd/main.go

build-linux:
	-rm -rf ./build/linux/bin
	# Update version
	cd pkg/version/scripts && go run gen.go
	# Build kusion
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 \
		go build -o ./build/linux/bin/$(APP_NAME) \
		./cmd/main.go

build-windows:
	-rm -rf ./build/windows/bin
	# Update version
	cd pkg/version/scripts && go run gen.go
	# Build kusion
	GOOS=windows GOARCH=amd64 CGO_ENABLED=0 \
		go build -o ./build/windows/bin/$(APP_NAME).exe \
		./cmd/main.go

test:
	@go test -timeout=10m `go list ./pkg/... ./cmd/...`

test-with-coverage:
	@go test -v -coverprofile=profile.cov -timeout=10m `go list ./pkg/... ./cmd/...`

lint:
	@$(GOCILINT) run --no-config --disable=errcheck ./...

gen-version: # Update version
	cd pkg/version/scripts && go run gen.go
