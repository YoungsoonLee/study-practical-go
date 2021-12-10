package middleware

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/YoungsoonLee/practical-go/ch05/complex-server/config"
	"github.com/YoungsoonLee/practical-go/ch05/complex-server/handlers"
)

func TestPanicMiddleware(t *testing.T) {
	b := new(bytes.Buffer)
	c := config.InitConfig(b)

	m := http.NewServeMux()
	handlers.Register(m, c)

	h := panicMiddleware(m, c)

	r := httptest.NewRequest("GET", "/panic", nil)
	w := httptest.NewRecorder()
	h.ServeHTTP(w, r)

	resp := w.Result()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("Error reading response body: %v", err)
	}

	if resp.StatusCode != http.StatusInternalServerError {
		t.Errorf(
			"Expected resonse status: %v, Got: %v\n",
			http.StatusOK,
			resp.StatusCode,
		)
	}

	expectedResponseBody := "unexpected server error occurred"

	if string(body) != expectedResponseBody {
		t.Errorf(
			"Expected response: %s, Got: %s\n",
			expectedResponseBody,
			string(body),
		)
	}
}
