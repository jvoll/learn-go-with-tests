package main

import (
	"fmt"
	"io"
	"net/http"
)

// Greet writes a hello message to the supplied buffer
func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

// MyGreeterHandler is an HTTP handler to show that Greet can also
// write HTTP responses.
func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "world")
}

func main() {
	// Greet(os.Stdout, "Elodie")
	http.ListenAndServe(":5002", http.HandlerFunc(MyGreeterHandler))
}
