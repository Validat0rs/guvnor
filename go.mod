module github.com/Validat0rs/guvnor

go 1.17

replace github.com/Validat0rs/guvnor => ./

require (
	github.com/BurntSushi/toml v1.1.0
	github.com/go-redis/redis/v8 v8.8.0
	github.com/go-redis/redismock/v8 v8.0.6
	github.com/gorilla/feeds v1.1.1
	github.com/gorilla/mux v1.8.0
	github.com/jaswdr/faker v1.10.2
	github.com/rs/zerolog v1.26.1
	github.com/urfave/negroni v1.0.0
	gopkg.in/check.v1 v1.0.0-20180628173108-788fd7840127
)

require (
	github.com/cespare/xxhash/v2 v2.1.1 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/kr/pretty v0.3.0 // indirect
	github.com/kr/text v0.2.0 // indirect
	github.com/onsi/ginkgo v1.16.5 // indirect
	github.com/onsi/gomega v1.19.0 // indirect
	github.com/rogpeppe/go-internal v1.6.1 // indirect
	go.opentelemetry.io/otel v0.19.0 // indirect
	go.opentelemetry.io/otel/metric v0.19.0 // indirect
	go.opentelemetry.io/otel/trace v0.19.0 // indirect
)
