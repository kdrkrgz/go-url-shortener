package repository

import (
	"github.com/kdrkrgz/go-url-shortener/resolver"
)

type AppRepository struct {
	DbRepository    resolver.DbRepository
	CacheRepository resolver.CacheRepository
}
