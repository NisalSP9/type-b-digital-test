package main

import (
	"log"
	"net/http"

	"github.com/NisalSP9/type-b-digital-test/internal/router"
)

func main() {

	port := "8080"
	mux := router.New()

	log.Printf("Server running on port %v", port)
	log.Fatal(http.ListenAndServe(":"+port, mux))
}
