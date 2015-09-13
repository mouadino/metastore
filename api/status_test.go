package api

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestStatusHandler(t *testing.T) {
	req, _ := http.NewRequest("GET", "/", nil)
	resp := httptest.NewRecorder()

	testServer.Mux.ServeHTTP(resp, req)

	if resp.Code != http.StatusOK {
		t.Errorf("/status/ expected to return %v but was %v", http.StatusOK, resp.Code)
	}

	contentType := resp.Header().Get("content-type")
	if contentType != JSONContentType {
		t.Errorf("/status/ expected to return json content-type but was %v", contentType)
	}

	var body map[string]interface{}
	if err := json.Unmarshal(resp.Body.Bytes(), &body); err != nil {
		t.Errorf("/status/ returned invalid json body: %v", resp.Body)
	}
	expectedBody := map[string]interface{}{
		"store": map[string]interface{}{
			"driver": "dummy",
		},
	}
	if !reflect.DeepEqual(body, expectedBody) {
		t.Errorf("/status/ returned body was %v expected %v", body, expectedBody)
	}
}
