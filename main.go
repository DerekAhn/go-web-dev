package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Hello, Go Web Development")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, Go Web Development")
	})

	fmt.Println(http.ListenAndServe(":3000", nil))
}
