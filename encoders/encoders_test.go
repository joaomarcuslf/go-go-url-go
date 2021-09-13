package encoders_test

import (
	"testing"

	encoders "github.com/joaomarcuslf/go-go-url-shortener/encoders"

	"github.com/stretchr/testify/assert"
)

func init() {
}

func TestEncode01(t *testing.T) {
	key := "joaomarcuslf"

	result := encoders.Encode(key)

	assert.True(t, result == "Ru4GBDi3")
}

func TestEncode02(t *testing.T) {
	key := "joaomarcuslf1"

	result := encoders.Encode(key)

	assert.True(t, result != "Ru4GBDi3")
}
