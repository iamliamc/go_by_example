go mod init your-module-name
# https://gobyexample.com/maps
# https://github.com/quii/learn-go-with-tests/blob/main

Run all tests
go test ./...
go test ./... -bench=. 
go test -cover