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

	config, err := hw.NewConfig()
	if err != nil {
		logger.Fatal(err)
	}

	logger.Printf("config: default content type: %s\n", config.ContentType)
	logger.Printf("config: default status code: %d\n", config.StatusCode)
	logger.Printf("config: default response body: %s\n", config.ResponseBody)
	logger.Printf("config: with robots.txt: %v\n", config.WithRobots)

	var handler http.Handler

	handler = hw.NewHandler(config)
	handler = hw.NewRequestLogHandler(handler, hw.RealTimer{}, logger)

	logger.Printf("start listening on port %v\n", config.ServerPort)

	if err := http.ListenAndServe(fmt.Sprintf(":%d", config.ServerPort), handler); !errors.Is(err, http.ErrServerClosed) {
		logger.Fatalf("ListenAndServe: %v", err)
	}
}
