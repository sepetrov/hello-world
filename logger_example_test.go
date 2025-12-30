package hello_world_test

import (
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"time"

	hw "github.com/sepetrov/hello-world"
)

type testTimer struct {
	now time.Time
}

func (t testTimer) Now() time.Time {
	return t.now
}

func ExampleRequestLogHandler_ServeHTTP() {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusCreated)
		_, _ = w.Write([]byte("Hello, World!"))
	})

	timer := testTimer{now: time.Date(2024, 10, 10, 13, 55, 36, 0, time.FixedZone("PDT", -7*3600))}
	logger := log.New(os.Stdout, "", 0)

	logHandler := hw.NewRequestLogHandler(handler, timer, logger)

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.RemoteAddr = "192.168.1.100:54321"
	w := httptest.NewRecorder()

	logHandler.ServeHTTP(w, req)

	// Output: 192.168.1.100:54321 - - [10/Oct/2024:13:55:36 -0700] "GET / HTTP/1.1" 201 13
}
