package encoders_test

import (
	"testing"

	encoders "github.com/joaomarcuslf/go-go-url-shortener/encoders"

	"github.com/stretchr/testify/assert"
)

func init() {
}

func TestBase62Encode(t *testing.T) {
	key := uint64(5577006791947779410)

	result := encoders.Encode(key)

	assert.True(t, result == "OTv0FdGU8Ng")
}

func TestBase62Decode01(t *testing.T) {
	key := "OTv0FdGU8Ng"

	result, err := encoders.Decode(key)

	assert.True(t, result == uint64(5577006791947779410))
	assert.True(t, err == nil)
}

func TestBase62Decode02(t *testing.T) {
	key := "OTv0FdGdfsU8Ng"

	result, err := encoders.Decode(key)

	assert.True(t, result != uint64(5577006791947779410))
	assert.True(t, err == nil)
}
