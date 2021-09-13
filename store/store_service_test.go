package store_test

import (
	"fmt"
	"testing"

	configs "github.com/joaomarcuslf/go-go-url-go/configs"
	store "github.com/joaomarcuslf/go-go-url-go/store"
	"github.com/stretchr/testify/assert"

	"github.com/joho/godotenv"
)

var envError = godotenv.Load("../.env")
var testStoreService = &store.StorageService{}
var config, _ = configs.FromEnv()

func init() {
	if envError != nil {
		panic(envError)
	}

	tss, err := store.InitializeStore(&config.Redis)

	if err != nil {
		panic(err)
	}

	tss.RedisClient.Del("testingKeyJoaomarcuslf01", "testingKeyJoaomarcuslf02", "testingKeyJoaomarcuslf03", "testingKeyJoaomarcuslf04", "testingKeyJoaomarcuslfadfsfg05")

	testStoreService = tss
}

func TestStoreInit(t *testing.T) {
	fmt.Println(testStoreService)
	assert.True(t, testStoreService.RedisClient != nil)
}

func TestSaveUrlMappingOnce(t *testing.T) {
	initialLink := "http://joaomarcuslf.com/"
	shortURL := "testingKeyJoaomarcuslf01"

	err := store.SaveUrlMapping(shortURL, initialLink)

	assert.True(t, err == nil)
}

func TestSaveUrlMappingTwice(t *testing.T) {
	initialLink := "http://joaomarcuslf.com/"
	shortURL := "testingKeyJoaomarcuslf02"

	store.SaveUrlMapping(shortURL, initialLink)

	err := store.SaveUrlMapping(shortURL, initialLink)

	assert.True(t, err != nil)
}

func TestRetrieveInitialUrl(t *testing.T) {
	initialLink := "http://joaomarcuslf.com/"
	shortURL := "testingKeyJoaomarcuslf03"

	store.SaveUrlMapping(shortURL, initialLink)

	result, _ := store.RetrieveInitialUrl(shortURL)

	assert.Equal(t, initialLink, result)
}

func TestSaveUrlMappingAndRetrieve(t *testing.T) {
	initialLink := "http://joaomarcuslf.com/"
	shortURL := "testingKeyJoaomarcuslf04"

	store.SaveUrlMapping(shortURL, initialLink)

	result, err := store.RetrieveInitialUrl(shortURL)

	assert.True(t, err == nil)
	assert.Equal(t, result, initialLink)
}

func TestRetrieveNotExists(t *testing.T) {
	shortURL := "testingKeyJoaomarcuslfadfsfg05"

	result, _ := store.RetrieveInitialUrl(shortURL)

	assert.True(t, result == "")
}
