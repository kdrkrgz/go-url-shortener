package repository

import (
	"fmt"
	"os"

	"github.com/go-redis/redis"
	resolver "github.com/kdrkrgz/go-url-shortener/resolver"
)

type RedisRepository struct {
	RedisClient *redis.Client
}

func NewRedisRepository() *RedisRepository {
	return &RedisRepository{
		RedisClient: getRedisClient(os.Getenv("RedisHost"), os.Getenv("RedisPort")),
	}
}

func getRedisClient(host, port string) *redis.Client {
	fmt.Println("Connecting to Redis...")
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", host, port),
		Password: os.Getenv("RedisPassword"),
		DB:       0,
	})
	fmt.Println("Connection Success!")
	return client
}

func (repo *RedisRepository) FindUrl(url string) (*string, error) {
	cur, err := repo.RedisClient.Get(url).Bytes()
	if err != nil {
		return nil, err
	}
	res := string(cur)

	return &res, nil
}

func (repo *RedisRepository) InsertShortUrl(shortedUrl resolver.ShortUrl) error {
	if err := repo.RedisClient.Set(shortedUrl.ShortUrl, shortedUrl.TargetUrl, 0).Err(); err != nil {
		err = fmt.Errorf("an error occured insert url to redis - Error: %v", err)
		return err
	}

	return nil
}

func (repo *RedisRepository) InsertTargetUrl(targetUrl resolver.ShortUrl) error {
	if err := repo.RedisClient.Set(targetUrl.TargetUrl, targetUrl.ShortUrl, 0).Err(); err != nil {
		err = fmt.Errorf("an error occured insert url to redis - Error: %v", err)
		return err
	}

	return nil
}

func (repo *RedisRepository) FlushAll() error {
	if err := repo.RedisClient.FlushAll().Err(); err != nil {
		err = fmt.Errorf("an error occured flush all - Error: %v", err)
		return err
	}

	return nil
}
