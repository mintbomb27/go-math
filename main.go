package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

func add(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GET params were:", r.URL.Query())
	num1, _ := strconv.Atoi(r.URL.Query().Get("num1"))
	num2, _ := strconv.Atoi(r.URL.Query().Get("num2"))

	result := num1 + num2
	fmt.Println("Answer: ", result)

	fmt.Fprintf(w, "The answer is %d", result)
}

func main() {
	http.HandleFunc("/", helloWorld)
	http.HandleFunc("/add", add)
	http.ListenAndServe(":8010", nil)
}
