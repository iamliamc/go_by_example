package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

// https://pkg.go.dev/fmt#Printf
// Instead of writer being *bytes.Buffer, we can use io.Writer
func Greet(writer io.Writer, name string) {
	// Its not easy to test fmt.Printf here so we use fmt.Fprintf and pass the buffer
	// fmt.Printf("Hello, %s", name)
	fmt.Fprintf(writer, "Hello, %s", name)
}

// MyGreeterHandler says Hello, world over HTTP.
func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "world")
}

// go run di.go
// http://localhost:5000/
func main() {
	Greet(os.Stdout, "Aoife")
	log.Fatal(http.ListenAndServe(":5000", http.HandlerFunc(MyGreeterHandler)))
}
