go mod init your-module-name
# https://gobyexample.com/maps
# https://github.com/quii/learn-go-with-tests/blob/main

Run all tests
`go test ./...`

Run all tests with benchmark
`go test ./... -bench=.` 

Run tests with coverage
`go test -cover`

Check for race conditions
`go test race`

Run tests with logging
`go test -v`

Don't remember
`go vet` 