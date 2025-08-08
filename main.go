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
	"Cat urine glows under a black-light."
}

type Fact struct {
	Fact string `json:"fact"`
}

func main() {
	rand.Seed(time.Now().UnixNano())

	http.HandleFunc("/", randomJokeHandler)

	println("Server running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func randomJokeHandler(w http.ResponseWriter, r *http.Request) {
	randomIndex := rand.Intn(len(coolFacts))
	fact := Fact{Fact: coolFacts[randomIndex]}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(fact)
}
