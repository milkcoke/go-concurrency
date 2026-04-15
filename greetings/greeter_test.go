package greetings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGreeter_Hello(t *testing.T) {
	// given
	var greeter = &Greeter{Prefix: "Go"}
	// when
	var result, err = greeter.Hello("World")
	// then
	if assert.NoError(t, err) {
		assert.Equal(t, "Go, World, Welcome!", result)
	}
}
