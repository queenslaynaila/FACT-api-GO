package main

import "testing"

func TestAddition(t *testing.T) {
    sum := 2 + 2
    if sum != 4 {
        t.Errorf("Expected 2 + 2 to equal 4, got %d", sum)
    }
}
