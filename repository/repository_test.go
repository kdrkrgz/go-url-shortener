package repository

import (
	"fmt"
	"os"
	"testing"

	"github.com/kdrkrgz/go-url-shortener/resolver"
	"gopkg.in/go-playground/assert.v1"
)

var (
	appRepository   *AppRepository
	dbRepository    resolver.DbRepository
	cacheRepository resolver.CacheRepository
)

func setup() {
	setupEnvs()
	dbRepository = NewMongoRepository()
	cacheRepository = NewRedisRepository()
	appRepository = &AppRepository{
		DbRepository:    dbRepository,
		CacheRepository: cacheRepository,
	}
}

func cleanup() {
	fmt.Println("Cleanup Test Database...")

	dbRepository.DropCollection()
	cacheRepository.FlushAll()
}

func TestMain(t *testing.T) {
	setup()
	InsertUrl(t)
	FindUrl(t)
	defer cleanup()
}

func setupEnvs() {
	os.Setenv("DbName", "GoUrlShortenerTests")
	os.Setenv("CollectionName", "Urls")
	os.Setenv("DbUri", "mongodb://localhost")
	os.Setenv("DbPort", "27017")
	os.Setenv("RedisHost", "localhost")
	os.Setenv("RedisPort", "6379")
}

func InsertUrl(t *testing.T) {
	fmt.Println("TestInsertUrl")
	t.Run("InsertUrl", func(t *testing.T) {
		shortUrl := resolver.ShortUrl{
			ShortUrl:  "test",
			TargetUrl: "test",
		}
		_, err := appRepository.DbRepository.InsertUrl(shortUrl)
		if err != nil {
			t.Errorf("InsertUrl() error = %v", err)
			return
		}
		if err := appRepository.CacheRepository.InsertShortUrl(shortUrl); err != nil {
			t.Errorf("InsertUrl() error = %v", err)
			return
		}

	})
}

func FindUrl(t *testing.T) {
	t.Run("FindUrl", func(t *testing.T) {
		dbResult, err := appRepository.DbRepository.FindUrl("short_url", "test")
		if err != nil {
			t.Errorf("FindUrl() error = %v", err)
			return
		}
		cacheResult, err := appRepository.CacheRepository.FindUrl("test")
		if err != nil {
			t.Errorf("FindUrl() error = %v", err)
			return
		}
		assert.Equal(t, dbResult.ShortUrl, "test")
		assert.Equal(t, cacheResult, "test")
	})
}
