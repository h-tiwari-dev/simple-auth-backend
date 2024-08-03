package utils

import (
	"fmt"
	"math/rand"
	"strings"
)

func GenerateRandomUsername() string {
	adjectives := []string{
		"brave", "calm", "eager", "fierce", "gentle", "happy", "jolly", "kind", "lively", "mighty",
	}
	nouns := []string{
		"lion", "tiger", "eagle", "wolf", "fox", "bear", "shark", "hawk", "panther", "whale",
	}

	// Select a random adjective and noun
	adj := adjectives[rand.Intn(len(adjectives))]
	noun := nouns[rand.Intn(len(nouns))]

	// Generate a random number between 1000 and 9999
	number := rand.Intn(9000) + 1000

	// Combine them into a username
	username := fmt.Sprintf("%s%s%d", adj, noun, number)

	// Return the username in lowercase
	return strings.ToLower(username)
}
