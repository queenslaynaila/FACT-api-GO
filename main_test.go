package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandleRandomFactRequest_NoRepeatUntilExhausted(t *testing.T) {
	// Reset factManager so tests are predictable
	factManager = createFactManager()

	seen := make(map[string]bool)

	// Call handler once for each fact
	for i := 0; i < len(factsDatabase); i++ {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		rec := httptest.NewRecorder()

		handleRandomFactRequest(rec, req)

		if rec.Code != http.StatusOK {
			t.Fatalf("expected 200 OK, got %d", rec.Code)
		}

		var resp FactResponse
		if err := json.Unmarshal(rec.Body.Bytes(), &resp); err != nil {
			t.Fatalf("failed to parse JSON: %v", err)
		}

		// Check if we got a repeat before exhausting
		if seen[resp.Fact] {
			t.Fatalf("fact repeated before all facts were shown: %q", resp.Fact)
		}
		seen[resp.Fact] = true
	}
}

func TestHandleRandomFactRequest_ReshufflesAfterExhaustion(t *testing.T) {
	factManager = createFactManager()

	for i := 0; i < len(factsDatabase); i++ {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		rec := httptest.NewRecorder()
		handleRandomFactRequest(rec, req)
	}

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	handleRandomFactRequest(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected 200 OK after exhaustion, got %d", rec.Code)
	}
}

func TestHandleAllFactsRequest_ReturnsFullList(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/all", nil)
	rec := httptest.NewRecorder()

	handleAllFactsRequest(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected 200 OK, got %d", rec.Code)
	}

	var allFacts []string
	if err := json.Unmarshal(rec.Body.Bytes(), &allFacts); err != nil {
		t.Fatalf("failed to parse JSON: %v", err)
	}

	if len(allFacts) != len(factsDatabase) {
		t.Fatalf("expected %d facts, got %d", len(factsDatabase), len(allFacts))
	}
}