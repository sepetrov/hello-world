package hello_world

import (
	"net/http"
	"time"
)

const timeFormat = "02/Jan/2006:15:04:05 -0700"

type Logger interface {
	Printf(format string, v ...any)
}

type Timer interface {
	Now() time.Time
}

type RequestLogHandler struct {
	handler http.Handler
	logger  Logger
	timer   Timer
}

type RealTimer struct{}

func (RealTimer) Now() time.Time {
	return time.Now()
}

func NewRequestLogHandler(handler http.Handler, timer Timer, logger Logger) http.Handler {
	return RequestLogHandler{
		handler: handler,
		logger:  logger,
		timer:   timer,
	}
}

func (h RequestLogHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	rw := newResponseWriter(w)

	h.handler.ServeHTTP(rw, r)

	h.logger.Printf("%s - - [%s] \"%s %s %s\" %d %d",
		r.RemoteAddr,
		h.timer.Now().Format(timeFormat),
		r.Method,
		r.URL.String(),
		r.Proto,
		rw.StatusCode,
		rw.Written,
	)
}

type responseWriter struct {
	http.ResponseWriter
	StatusCode int
	Written    int
}

func newResponseWriter(w http.ResponseWriter) *responseWriter {
	return &responseWriter{
		ResponseWriter: w,
		StatusCode:     http.StatusOK,
	}
}

func (rw *responseWriter) WriteHeader(statusCode int) {
	rw.StatusCode = statusCode
	rw.ResponseWriter.WriteHeader(statusCode)
}

func (rw *responseWriter) Write(b []byte) (int, error) {
	n, err := rw.ResponseWriter.Write(b)
	rw.Written += n
	return n, err
}
