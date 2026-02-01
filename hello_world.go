package hello_world

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/kelseyhightower/envconfig"
)

const headerContentType = "Content-Type"

type Config struct {
	ServerPort   int    `envconfig:"SERVER_PORT" default:"8080"`
	ContentType  string `envconfig:"CONTENT_TYPE" default:"text/plain"`
	StatusCode   int    `envconfig:"STATUS_CODE" default:"200"`
	ResponseBody string `envconfig:"RESPONSE_BODY" default:"Hello World!"`
	WithRobots   bool   `envconfig:"WITH_ROBOTS_TXT" default:"false"`
}

func NewConfig() (Config, error) {
	var c Config
	if err := envconfig.Process("", &c); err != nil {
		return Config{}, fmt.Errorf("NewConfig: parse configuration: %w", err)
	}

	return c, nil
}

func NewHandler(c Config) http.Handler {
	var h http.Handler

	h = Handler{
		ContentType:  c.ContentType,
		StatusCode:   c.StatusCode,
		ResponseBody: c.ResponseBody,
	}

	if c.WithRobots {
		h = RobotsHandler{Handler: h}
	}

	return h
}

type Handler struct {
	ContentType  string
	StatusCode   int
	ResponseBody string
}

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(headerContentType, param("content_type", r, h.ContentType))

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
