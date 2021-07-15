module github.com/resurfaceio/test-mux

go 1.16

require (
	github.com/99designs/gqlgen v0.13.0
	github.com/agnivade/levenshtein v1.1.1 // indirect
	github.com/gorilla/mux v1.8.0
	github.com/hashicorp/golang-lru v0.5.4 // indirect
	github.com/mitchellh/mapstructure v1.4.1 // indirect
	github.com/resurfaceio/logger-go v0.0.0-20210715202310-8a143a2013a9
	github.com/vektah/gqlparser/v2 v2.2.0
)

// added for testing with local version of go logger
replace github.com/resurfaceio/logger-go => ./../../logger-go
