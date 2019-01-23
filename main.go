package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func sendResponse(w http.ResponseWriter, r *http.Request) {
	keys, ok := r.URL.Query()["a"]
	if !ok || len(keys[0]) < 1 {
		log.Println("Url Param 'a' is missing")
		return
	}

	a, err := strconv.Atoi(keys[0])

	if err != nil {
		return
	}

	keys, ok = r.URL.Query()["b"]
	if !ok || len(keys[0]) < 1 {
		log.Println("Url Param 'b' is missing")
		return
	}
	b, err := strconv.Atoi(keys[0])
	if err != nil {
		log.Println("Bad number!")
		return
	}

	fmt.Fprintf(w, "%d + %d = %d", a, b, sum(a, b))
}

func main() {
	http.HandleFunc("/", sendResponse)
	log.Fatal(http.ListenAndServe(":5000", nil))
}

func sum(a, b int) int {
	return a + b
}
