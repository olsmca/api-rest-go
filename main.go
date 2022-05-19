package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	router := NewRouter()

	fmt.Println("Server Start:", 3333)
	server := http.ListenAndServe(":3333", router)

	log.Fatal(server)
}
