package hello_world_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"

	hw "github.com/sepetrov/hello-world"
)

func ExampleHandler_ServeHTTP() {
	handler := hw.Handler{
		ContentType:  "text/plain",
		StatusCode:   http.StatusOK,
		ResponseBody: "Hello World!",
	}

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()

	handler.ServeHTTP(rec, req)

	fmt.Println(rec.Code)
	fmt.Println(rec.Header().Get("Content-Type"))
	fmt.Println(rec.Body.String())
	// Output:
	// 200
	// text/plain
	// Hello World!
}
