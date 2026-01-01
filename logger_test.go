package hello_world_test

import (
	"bytes"
	"log"
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"

	hw "github.com/sepetrov/hello-world"
)

func TestRequestLogHandler_ServeHTTP(t *testing.T) {
	tests := []struct {
		name           string
		method         string
		path           string
		handler        http.Handler
		wantStatusCode int
		wantBody       string
	}{
		{
			name:           "GET request with 200 response",
			method:         http.MethodGet,
			path:           "/?foo=bar",
			handler:        http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { _, _ = w.Write([]byte("Hello")) }),
			wantStatusCode: http.StatusOK,
			wantBody:       "Hello",
		},
		{
			name:           "POST request with 201 response",
			method:         http.MethodPost,
			path:           "/create",
			handler:        http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { w.WriteHeader(http.StatusCreated) }),
			wantStatusCode: http.StatusCreated,
			wantBody:       "",
		},
		{
			name:           "GET request with 404 response",
			method:         http.MethodGet,
			path:           "/notfound",
			handler:        http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { w.WriteHeader(http.StatusNotFound) }),
			wantStatusCode: http.StatusNotFound,
			wantBody:       "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer

			logger := log.New(&buf, "", 0)
			timer := hw.RealTimer{}

			handler := hw.NewRequestLogHandler(tt.handler, timer, logger)

			req := httptest.NewRequest(tt.method, tt.path, nil)
			req.Header.Set("User-Agent", "TestAgent/1.0")

			w := httptest.NewRecorder()

			handler.ServeHTTP(w, req)

			if w.Code != tt.wantStatusCode {
				t.Errorf("status code = %d, want %d", w.Code, tt.wantStatusCode)
			}

			if w.Body.String() != tt.wantBody {
				t.Errorf("body = %q, want %q", w.Body.String(), tt.wantBody)
			}

			logOutput := buf.String()
			t.Log(logOutput)

			pattern := `\d+\.\d+\.\d+\.\d+:\d+ - .* \[\d{2}/\w{3}/\d{4}:\d{2}:\d{2}:\d{2} [+-]\d{4}\] "` +
				tt.method + ` ` + regexp.QuoteMeta(tt.path) + ` HTTP/1\.1" ` +
				`\d+ \d+`

			matched, err := regexp.MatchString(pattern, logOutput)
			if err != nil {
				t.Fatalf("regex error: %v", err)
			}
			if !matched {
				t.Errorf("log output does not match expected format:\n%s", logOutput)
			}
		})
	}
}
