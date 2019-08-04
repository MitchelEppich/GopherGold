package models

import (
	"testing"

	bandname "github.com/davidbanham/bandname_go"
	"github.com/stretchr/testify/assert"
)

func TestRandom(t *testing.T) {
	t.Parallel()

	joke := Joke{}
	assert.Nil(t, joke.Random())
}

func TestSave(t *testing.T) {
	t.Parallel()

	joke := Joke{
		Name: bandname.Bandname(),
		Body: bandname.Bandname(),
	}
	assert.Nil(t, joke.Save())
}
