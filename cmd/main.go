package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	hw "github.com/sepetrov/hello-world"
)

func main() {
	h, err := hw.NewHandler()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("default content type: %s\n", h.ContentType)
	log.Printf("default status code: %d\n", h.StatusCode)
	log.Printf("default response body: %s\n", h.ResponseBody)

	log.Printf("start listening on port %v\n", h.ServerPort)

	if err := http.ListenAndServe(fmt.Sprintf(":%d", h.ServerPort), h); !errors.Is(err, http.ErrServerClosed) {
		log.Fatalf("ListenAndServe: %v", err)
	}
}
