package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
	// Return a greeting that embeds the name in a message.
	if name == "" {
		return name, errors.New("empty name")
	}
	// Create a message using a random format
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

func randomFormat() string {
	// A slice of message formats.
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}

	// Return one of the message formats selected at random.
	return formats[rand.Intn(len(formats))]
}
