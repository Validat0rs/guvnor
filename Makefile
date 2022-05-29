GOBIN?=${GOPATH}/bin

all: lint install

lint-pre:
	@test -z $(gofmt -l .)
	@go mod verify

lint: lint-pre
	@golangci-lint run

lint-verbose: lint-pre
	@golangci-lint run -v --timeout=5m

install: go.sum
	GO111MODULE=on go install -v ./cmd/guvnord

clean:
	rm -f ${GOBIN}/{guvnord}

tests:
	@go test -v -coverprofile .testCoverage.txt ./...

docker-redis:
	@docker-compose -f docker-compose.yml run --rm -p 6379:6379 --no-deps -d redis

setup-config:
	@mkdir -p ~/.guvnor/config
	@cp config.toml.example ${HOME}/.guvnor/config/config.toml
	@export GUVNOR_CONFIG=${HOME}/.guvnor/config/config.toml

run-guvnor-service:
	@guvnord
