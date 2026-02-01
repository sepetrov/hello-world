package hello_world_test

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	hw "github.com/sepetrov/hello-world"
)

func TestHandler_ServeHTTP(t *testing.T) {
	t.Run("responds with default values", func(t *testing.T) {
		handler := hw.Handler{
			ContentType:  "text/plain",
			StatusCode:   http.StatusOK,
			ResponseBody: "Hello World!",
		}

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		rec := httptest.NewRecorder()

		handler.ServeHTTP(rec, req)

		if rec.Code != 200 {
			t.Errorf("expected status code 200, got %d", rec.Code)
		}
		if ct := rec.Header().Get("Content-Type"); ct != "text/plain" {
			t.Errorf("expected Content-Type text/plain, got %s", ct)
		}
		if body := rec.Body.String(); body != "Hello World!" {
			t.Errorf("expected body 'Hello World!', got %s", body)
		}
	})

	t.Run("ignores request parameters if runtime configuration is disabled", func(t *testing.T) {
		handler := hw.Handler{
			ContentType:       "text/plain",
			StatusCode:        http.StatusOK,
			ResponseBody:      "Hello World!",
			WithRuntimeConfig: false,
		}

		form := url.Values{}
		form.Add("content_type", "application/xml")
		form.Add("status_code", "404")
		form.Add("response_body", "not found")

		req := httptest.NewRequest(http.MethodPost, "/?status_code=500", strings.NewReader(form.Encode()))
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		rec := httptest.NewRecorder()

		handler.ServeHTTP(rec, req)

		if rec.Code != handler.StatusCode {
			t.Errorf("expected status code %d, got %d", handler.StatusCode, rec.Code)
		}
		if ct := rec.Header().Get("Content-Type"); ct != handler.ContentType {
			t.Errorf("expected Content-Type %q, got %s", handler.ContentType, ct)
		}
		if body := rec.Body.String(); body != handler.ResponseBody {
			t.Errorf("expected body %q, got %s", handler.ResponseBody, body)
		}
	})

	t.Run("responds with query parameters if runtime configuration is enabled", func(t *testing.T) {
		handler := hw.Handler{
			WithRuntimeConfig: true,
		}

		req := httptest.NewRequest(http.MethodGet, "/?content_type=application/json&status_code=201&response_body=created", nil)
		rec := httptest.NewRecorder()

		handler.ServeHTTP(rec, req)

		if rec.Code != 201 {
			t.Errorf("expected status code 201, got %d", rec.Code)
		}
		if ct := rec.Header().Get("Content-Type"); ct != "application/json" {
			t.Errorf("expected Content-Type application/json, got %s", ct)
		}
		if body := rec.Body.String(); body != "created" {
			t.Errorf("expected body 'created', got %s", body)
		}
	})

	t.Run("responds with form parameters if runtime configuration is enabled", func(t *testing.T) {
		handler := hw.Handler{
			WithRuntimeConfig: true,
		}

		form := url.Values{}
		form.Add("content_type", "application/xml")
		form.Add("status_code", "404")
		form.Add("response_body", "not found")

		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(form.Encode()))
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		rec := httptest.NewRecorder()

		handler.ServeHTTP(rec, req)

		if rec.Code != 404 {
			t.Errorf("expected status code 404, got %d", rec.Code)
		}
		if ct := rec.Header().Get("Content-Type"); ct != "application/xml" {
			t.Errorf("expected Content-Type application/xml, got %s", ct)
		}
		if body := rec.Body.String(); body != "not found" {
			t.Errorf("expected body 'not found', got %s", body)
		}
	})
}
