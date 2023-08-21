package resolver

import "go.mongodb.org/mongo-driver/mongo"

type DbRepository interface {
	FindUrl(field, url string) (*ShortUrl, error)
	InsertUrl(shortedUrl ShortUrl) (*mongo.InsertOneResult, error)
	DeleteUrlsByDate() (*mongo.DeleteResult, error)
	DropCollection()
}

type CacheRepository interface {
	FindUrl(url string) (*string, error)
	InsertShortUrl(shortedUrl ShortUrl) error
	InsertTargetUrl(targetUrl ShortUrl) error
	FlushAll() error
}
