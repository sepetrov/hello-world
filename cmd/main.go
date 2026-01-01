package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	hw "github.com/sepetrov/hello-world"
)

func main() {
	logger := log.New(log.Writer(), "", 0)

	h, err := hw.NewHandler()
	if err != nil {
		logger.Fatal(err)
	}

	lh := hw.NewRequestLogHandler(h, hw.RealTimer{}, logger)

	logger.Printf("default content type: %s\n", h.ContentType)
	logger.Printf("default status code: %d\n", h.StatusCode)
	logger.Printf("default response body: %s\n", h.ResponseBody)

	logger.Printf("start listening on port %v\n", h.ServerPort)

	if err := http.ListenAndServe(fmt.Sprintf(":%d", h.ServerPort), lh); !errors.Is(err, http.ErrServerClosed) {
		logger.Fatalf("ListenAndServe: %v", err)
	}
}
