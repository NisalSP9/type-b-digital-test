package handler

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func decodeResponse(t *testing.T, rr *httptest.ResponseRecorder) map[string]string {
	t.Helper()
	var body map[string]string
	if err := json.Unmarshal(rr.Body.Bytes(), &body); err != nil {
		t.Fatalf("Failed to decode JSON: %v", err)
	}
	return body
}

func TestCheckName(t *testing.T) {
	tests := []struct {
		name           string
		queryParam     string
		expectedStatus int
		expectedKey    string
	}{
		{"ValidFirstHalf #1", "Alice", http.StatusOK, "message"},
		{"ValidFirstHalf #2", "Mike", http.StatusOK, "message"},
		{"InvalidSecondHalf #1", "Nisal", http.StatusBadRequest, "error"},
		{"InvalidSecondHalf #2", "Zara", http.StatusBadRequest, "error"},
		{"InvalidSymbol #1", "123", http.StatusBadRequest, "error"},
		{"InvalidSymbol #2", "!@#$%^&", http.StatusBadRequest, "error"},
		{"EmptyName", "", http.StatusBadRequest, "error"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "/check?name="+tc.queryParam, nil)
			rr := httptest.NewRecorder()

			CheckName(rr, req)

			res := rr.Result()
			defer res.Body.Close()

			if res.StatusCode != tc.expectedStatus {
				t.Errorf("Expected status %d, got %d", tc.expectedStatus, res.StatusCode)
			}

			body := decodeResponse(t, rr)
			if _, ok := body[tc.expectedKey]; !ok {
				t.Errorf("Expected key %q in response, got %v", tc.expectedKey, body)
			}
		})
	}
}
