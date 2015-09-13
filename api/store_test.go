package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestStoreGetHandler(t *testing.T) {
	req, _ := http.NewRequest("GET", "/store/abc/", nil)
	resp := httptest.NewRecorder()

	testServer.Mux.ServeHTTP(resp, req)

	if resp.Code != http.StatusOK {
		t.Errorf("/store/ expected to return %v but was %v", http.StatusOK, resp.Code)
	}

	contentType := resp.Header().Get("content-type")
	if contentType != JSONContentType {
		t.Errorf("/store/ expected to return json content-type but was %v", contentType)
	}

	var body string
	if err := json.Unmarshal(resp.Body.Bytes(), &body); err != nil {
		t.Errorf("/store/ returned invalid json body: %v", resp.Body)
	}
	expectedBody := "foobar"
	if body != expectedBody {
		t.Errorf("/store/ returned body was %v expected %v", body, expectedBody)
	}
}

func TestStorePostHandler(t *testing.T) {
	body := bytes.NewBuffer([]byte(`{"key": "a", "value": "bar"}`))
	req, _ := http.NewRequest("POST", "/store/", body)
	req.Header.Set("Content-Type", JSONContentType)
	resp := httptest.NewRecorder()

	testServer.Mux.ServeHTTP(resp, req)

	if resp.Code != http.StatusCreated {
		t.Errorf("/store/ expected to return %v but was %v", http.StatusCreated, resp.Code)
	}
}
