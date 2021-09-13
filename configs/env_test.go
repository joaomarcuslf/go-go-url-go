package configs_test

import (
	"testing"

	configs "github.com/joaomarcuslf/go-go-url-shortener/configs"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func init() {
}

func TestFromEnvNoEnvfile(t *testing.T) {
	config, _ := configs.FromEnv()
	assert.True(t, config.Redis.Port == "")
}

func TestFromEnvWitEnvfile(t *testing.T) {
	_ = godotenv.Load("../.env")
	config, _ := configs.FromEnv()
	assert.True(t, config.Redis.Port != "")
}
