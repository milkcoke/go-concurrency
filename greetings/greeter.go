package greetings

import (
	"errors"
	"fmt"
)

type Greeter struct {
	Prefix string
}

func (g *Greeter) Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}

	return fmt.Sprintf("%s, %s, Welcome!", g.Prefix, name), nil
}
