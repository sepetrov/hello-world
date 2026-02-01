package hello_world

import (
	"net/http"
)

const (
	robotsPath        = "/robots.txt"
	robotsContentType = "text/plain"
	robotsBody        = "User-agent: *\nDisallow: /"
)

type RobotsHandler struct {
	Handler http.Handler
}

func (h RobotsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet && r.URL.Path == robotsPath {
		w.Header().Set(headerContentType, robotsContentType)
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(robotsBody))

		return
	}

	h.Handler.ServeHTTP(w, r)
}
