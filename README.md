# optional

Defines optional types based on generics with wrappers for JSON, XML and SQL

## receipts

Notes for some CLI commands

### install ginkgo

go get github.com/onsi/ginkgo/v2/ginkgo
go install github.com/onsi/ginkgo/v2/ginkgo
go get github.com/onsi/gomega/...

### bootstrap ginkgo

ginkgo.exe bootstrap

### create template for optional.go

ginkgo.exe generate optional

### run tests

ginkgo.exe

### get coverage tests

go test -v -coverpkg=./... -coverprofile=profile.cov ./...
go tool cover -func profile
