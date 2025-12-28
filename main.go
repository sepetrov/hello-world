package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/kelseyhightower/envconfig"
)

func main() {
	h, err := NewHandler()
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

type Handler struct {
	ServerPort   int    `envconfig:"SERVER_PORT" default:"8080"`
	ContentType  string `envconfig:"CONTENT_TYPE" default:"text/plain"`
	StatusCode   int    `envconfig:"STATUS_CODE" default:"200"`
	ResponseBody string `envconfig:"RESPONSE_BODY" default:"Hello World!"`
}

func NewHandler() (Handler, error) {
	var h Handler
	if err := envconfig.Process("", &h); err != nil {
		return Handler{}, fmt.Errorf("NewHandler: parse configuration: %w", err)
	}

	return h, nil
}

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", param("content_type", r, h.ContentType))

	if i, err := strconv.Atoi(param("status_code", r, strconv.Itoa(h.StatusCode))); err == nil {
		w.WriteHeader(i)
	}

	_, _ = w.Write([]byte(param("response_body", r, h.ResponseBody)))
}

func param(name string, r *http.Request, def string) string {
	if s := r.FormValue(name); s != "" {
		return s
	}

	return def
}
