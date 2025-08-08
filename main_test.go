package main

import (
	   "encoding/json"
	   "math/rand"
	   "net/http/httptest"
	   "testing"
)

func TestNoImmediateRepetition(t *testing.T) {
	// Reset the shuffle state to simulate fresh server state
	servedFacts = make(map[int]bool)
	servedCount = 0
	shuffledIndexes = []int{}
	currentIndex = 0

	// Seed a new shuffle
	shuffledIndexes = shuffledIndexes[:0]
	shuffledIndexes = rand.Perm(len(coolFacts))
	currentIndex = 0

	seen := make(map[string]bool)

	// Hit the endpoint as many times as there are facts
	for i := 0; i < len(coolFacts); i++ {
		req := httptest.NewRequest("GET", "/", nil)
		w := httptest.NewRecorder()

		randomFactHandler(w, req)

		resp := w.Result()
		defer resp.Body.Close()

		var result Fact
		err := json.NewDecoder(resp.Body).Decode(&result)
		if err != nil {
			t.Fatalf("Failed to decode response: %v", err)
		}

		if seen[result.Fact] {
			t.Errorf("Duplicate fact detected before all facts were served: %s", result.Fact)
		}
		seen[result.Fact] = true
	}
}
