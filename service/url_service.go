package service

import (
	"fmt"

	log "github.com/kdrkrgz/go-url-shortener/pkg/logger"
	"github.com/kdrkrgz/go-url-shortener/resolver"
)

type UrlServiceApi interface {
	FindUrl(url string) (*string, error)
	InsertUrl(shortedUrl resolver.ShortUrl) error
}

type UrlService struct {
	DbRepository    resolver.DbRepository
	CacheRepository resolver.CacheRepository
}

func (service *UrlService) InsertUrl(shortedUrl resolver.ShortUrl) error {
	_, err := service.DbRepository.InsertUrl(shortedUrl)
	if err != nil {
		log.Logger().Sugar().Errorf("An error occured insert url to db - Error: %v", err)
		return err
	}
	if err := service.CacheRepository.InsertTargetUrl(shortedUrl); err != nil {
		log.Logger().Sugar().Errorf("An error occured insert url to redis - Error: %v", err)
		return err
	}

	return nil
}

func (service *UrlService) FindTargetUrl(url string) (*string, error) {
	resFromCache, err := service.CacheRepository.FindUrl(url)
	if err != nil {
		log.Logger().Sugar().Infof("Url Not Found on Cache - Url: %s: err: %s", url, err)
	}
	if resFromCache != nil {
		return resFromCache, nil
	}
	shortUrl, err := service.DbRepository.FindUrl("short_url", url)
	if err != nil {
		return nil, err
	}
	fmt.Println("resFromDb: ", shortUrl.TargetUrl)
	service.CacheRepository.InsertShortUrl(*shortUrl)
	return &shortUrl.TargetUrl, nil
}

func (service *UrlService) FindShortUrl(url string) (*string, error) {
	resFromCache, err := service.CacheRepository.FindUrl(url)
	if err != nil {
		log.Logger().Sugar().Infof("Url Not Found on Cache - Url: %s: err: %s", url, err)
	}
	if resFromCache != nil {
		fmt.Println("resFromCache: ", *resFromCache)
		return resFromCache, nil
	}
	resFromDb, err := service.DbRepository.FindUrl("target_url", url)
	if err != nil {
		return nil, err
	}
	fmt.Println("resFromDb: ", resFromDb.TargetUrl)
	// request to db for this url exists Bloom Filter coming here
	service.CacheRepository.InsertTargetUrl(*resFromDb)
	return &resFromDb.ShortUrl, nil
}
