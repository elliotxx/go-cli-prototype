# Default variable
GOFORMATER 	?= goimports
GOCILINT 	?= golangci-lint
APP_NAME 	?= go-cli-prototype

default:
	@go run cmd/main.go

format:
	@$(GOFORMATER) -l -w .

test:
	@go test -timeout=10m `go list ./pkg/... ./cmd/...`

test-with-coverage:
	@go test -v -coverprofile=profile.cov -timeout=10m `go list ./pkg/... ./cmd/...`

lint:
	@$(GOCILINT) run --no-config --disable=errcheck ./...

gen-version: # Update version
	cd pkg/version/scripts && go run gen.go

build-all: build-darwin build-linux build-windows

build-darwin: gen-version
	-rm -rf ./build/darwin/bin
	# Build kusion
	GOOS=darwin GOARCH=amd64 CGO_ENABLED=0 \
		go build -o ./build/darwin/bin/$(APP_NAME) \
		./cmd/main.go

build-linux: gen-version
	-rm -rf ./build/linux/bin
	# Build kusion
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 \
		go build -o ./build/linux/bin/$(APP_NAME) \
		./cmd/main.go

build-windows: gen-version
	-rm -rf ./build/windows/bin
	# Build kusion
	GOOS=windows GOARCH=amd64 CGO_ENABLED=0 \
		go build -o ./build/windows/bin/$(APP_NAME).exe \
		./cmd/main.go
