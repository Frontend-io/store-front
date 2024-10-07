package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		n, err := fmt.Fprintf(w, "Hello, world!")
		if err != nil {
			fmt.Sprintln("An error occured", err)
		}

		val := fmt.Sprintf("Numbers of byte written: %v", n)
		fmt.Println(val)
	})

	_ = http.ListenAndServe(":8080", nil)
}
