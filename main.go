package main

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"time"
)

var coolFacts = []string{
	"Like fingerprints, everyone's tongue print is different.",
	"Penguins propose with pebbles, yes pebble stones. Cheap and effective, maybe, but don’t try that.",
	"You can’t hum while holding your nose. Go on, try it.",
	"Taping your mouth while sleeping can help with snoring, but don’t tell people you do it.",
	"There's a rare condition called Auto-Brewery Syndrome — your gut literally brews alcohol from carbs you ate. Basically, a walking alcohol brewery.",
	"Octopuses have three hearts.",
	"There’s a real condition where you can’t control continuous farting — it’s called flatulence incontinence. May all my enemies experience it.",
	"There's this animal called a wombat. It has cube-shaped poop. Yes, good data structures?",
	"You hear that voice in your head when you read? Some people don’t have it.",
	"Cat urine glows under a black-light.",
}

type Fact struct {
	Fact      string `json:"fact"`
	Timestamp string `json:"timestamp"`
}

var servedFacts map[int]bool
var servedCount int
var shuffledIndexes []int
var currentIndex int

func main() {
	rand.Seed(time.Now().UnixNano())

	servedFacts = make(map[int]bool)
	servedCount = 0
	shuffledIndexes = rand.Perm(len(coolFacts))
	currentIndex = 0

	http.HandleFunc("/", randomFactHandler)
	http.HandleFunc("/all", allFactsHandler)

	println("Server running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func randomFactHandler(w http.ResponseWriter, r *http.Request) {
	// No repeat until all facts have been shown
	if currentIndex >= len(shuffledIndexes) {
		// Reshuffle for a new round
		shuffledIndexes = rand.Perm(len(coolFacts))
		currentIndex = 0
		servedFacts = make(map[int]bool)
		servedCount = 0
	}
	idx := shuffledIndexes[currentIndex]
	currentIndex++
	servedFacts[idx] = true
	servedCount = len(servedFacts)

	fact := Fact{
		Fact:      coolFacts[idx],
		Timestamp: time.Now().Format(time.RFC3339),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(fact)
}

func allFactsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(coolFacts)
}
