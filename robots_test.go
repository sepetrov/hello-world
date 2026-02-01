package hello_world_test

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	hw "github.com/sepetrov/hello-world"
)

func TestRobotsHandler_ServeHTTP(t *testing.T) {
	t.Run("renders robots.txt for /robots.txt path", func(t *testing.T) {
		dh := http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write([]byte("bad request"))
		})
		rh := hw.RobotsHandler{Handler: dh}
		svr := httptest.NewServer(rh)
		defer svr.Close()

		resp, err := http.Get(svr.URL + "/robots.txt")
		if err != nil {
			t.Fatal(err)
		}
		defer func() { _ = resp.Body.Close() }()

		if resp.StatusCode != http.StatusOK {
			t.Errorf("expected status code %d, got %d", http.StatusOK, resp.StatusCode)
		}
		if ct := resp.Header.Get("Content-Type"); ct != "text/plain" {
			t.Errorf("expected content type 'text/plain', got '%s'", ct)
		}

		b, err := io.ReadAll(resp.Body)
		if err != nil {
			t.Fatal(err)
		}

		const expectedBody = "User-agent: *\nDisallow: /"
		if string(b) != expectedBody {
			t.Errorf("expected body '%s', got '%s'", expectedBody, string(b))
		}
	})
	t.Run("forwards other paths to underlying handler", func(t *testing.T) {
		const (
			expectedStatusCode = http.StatusTeapot
			expectedBody       = "I'm a teapot"
		)

		dh := http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
			w.WriteHeader(expectedStatusCode)
			_, _ = w.Write([]byte(expectedBody))
		})
		rh := hw.RobotsHandler{Handler: dh}
		svr := httptest.NewServer(rh)
		defer svr.Close()

		testCases := []struct {
			Method string
			Path   string
		}{
			{Method: http.MethodGet, Path: "/"},
			{Method: http.MethodPost, Path: "/robots.txt"},
			{Method: http.MethodGet, Path: "/some/other/path"},
		}

		for _, tc := range testCases {
			t.Run(fmt.Sprintf("%s %s", tc.Method, tc.Path), func(t *testing.T) {
				req, err := http.NewRequest(tc.Method, svr.URL+tc.Path, nil)
				if err != nil {
					t.Fatal(err)
				}

				resp, err := http.DefaultClient.Do(req)
				if err != nil {
					t.Fatal(err)
				}
				defer func() { _ = resp.Body.Close() }()

				if resp.StatusCode != expectedStatusCode {
					t.Errorf("expected status code %d, got %d", expectedStatusCode, resp.StatusCode)
				}

				b, err := io.ReadAll(resp.Body)
				if err != nil {
					t.Fatal(err)
				}

				if string(b) != expectedBody {
					t.Errorf("expected body '%s', got '%s'", expectedBody, string(b))
				}
			})
		}
	})
}
