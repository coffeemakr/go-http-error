package http_error

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestJsonWriter(t *testing.T) {
	var badRequestType = NewHttpErrorType(http.StatusBadRequest, "Bad request")
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/", nil)
	badRequestType.CauseString("dummy error").Write(w, r)
	response := w.Result()
	decoder := json.NewDecoder(response.Body)
	var body map[string]interface{}
	err := decoder.Decode(&body)
	if err != nil {
		t.Fatal(err)
	}

	if body["description"] != "Bad request" {
		t.Fatalf("Got invalid description: %s", body["description"])
	}
	if body["error"] != true {
		t.Fatalf("Error field is not true: %s", body["true"])
	}
	if response.StatusCode != http.StatusBadRequest {
		t.Fatalf("Status code is not set to 400, it's set to %d", w.Code)
	}

	if contentType := w.Header().Get("Content-Type"); contentType != "application/json" {
		t.Fatalf("Invalid content type: %s", contentType)
	}
}
