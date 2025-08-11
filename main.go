package main

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"time"
)


var factsDatabase = []string{
	"Like fingerprints, everyone's tongue print is different.",
	"Penguins propose with pebbles, yes pebble stones. Cheap and effective, maybe, but don't try that.",
	"You can't hum while holding your nose. Go on, try it.",
	"Taping your mouth while sleeping can help with snoring, but don't tell people you do it.",
	"There's a rare condition called Auto-Brewery Syndrome — your gut literally brews alcohol from carbs you ate. Basically, a walking alcohol brewery.",
	"Octopuses have three hearts.",
	"There's a real condition where you can't control continuous farting — it's called flatulence incontinence. May all my enemies experience it.",
	"There's this animal called a wombat. It has cube-shaped poop. Yes, good data structures?",
	"You hear that voice in your head when you read? Some people don't have it.",
	"Cat urine glows under a black-light.",
}

type FactResponse struct {
	Fact      string `json:"fact"`
	Timestamp string `json:"timestamp"`
}

type FactManager struct {
	shuffledIndices []int
	currentIndex    int
	totalFacts      int
}

var factManager = createFactManager()

func shuffleFactOrder(manager *FactManager) {
	manager.shuffledIndices = rand.Perm(manager.totalFacts)
	manager.currentIndex = 0
}

func createFactManager() FactManager {
	rand.Seed(time.Now().UnixNano())
	
	manager := FactManager{
		totalFacts: len(factsDatabase),
	}
	shuffleFactOrder(&manager)
	return manager
}

func handleRandomFactRequest(w http.ResponseWriter, r *http.Request) {
	// If we've gone through all facts, shuffle again
	if factManager.currentIndex >= len(factManager.shuffledIndices) {
		shuffleFactOrder(&factManager)
	}

	// Get the next fact from the shuffled order
	factIndex := factManager.shuffledIndices[factManager.currentIndex]
	factManager.currentIndex++

	fact := FactResponse{
		Fact:      factsDatabase[factIndex],
		Timestamp: time.Now().Format(time.RFC3339),
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(fact); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}

func handleAllFactsRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(factsDatabase); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}


func main() {
	http.HandleFunc("/", handleRandomFactRequest)
	http.HandleFunc("/all", handleAllFactsRequest)
	
	println("Random Facts API server running on http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic("Failed to start server: " + err.Error())
	}
}