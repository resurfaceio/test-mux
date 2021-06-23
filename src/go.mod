module github.com/resurfaceio/test-mux

go 1.16

require (
	github.com/99designs/gqlgen v0.13.0
	github.com/agnivade/levenshtein v1.1.1 // indirect
	github.com/asaskevich/govalidator v0.0.0-20210307081110-f21760c49a8d // indirect
	github.com/gorilla/mux v1.8.0
	github.com/gorilla/websocket v1.4.2 // indirect
	github.com/hashicorp/golang-lru v0.5.4 // indirect
	github.com/mitchellh/mapstructure v1.4.1 // indirect
	github.com/resurfaceio/logger-go v0.0.0-20210617234230-0164f5c405a1
	github.com/vektah/gqlparser/v2 v2.2.0
)

// added for testing with local version of go logger
// replace github.com/resurfaceio/logger-go => ./../../logger-go
