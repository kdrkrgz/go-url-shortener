package pkg

import (
	"time"

	"github.com/go-co-op/gocron"
	"github.com/kdrkrgz/go-url-shortener/conf"
	log "github.com/kdrkrgz/go-url-shortener/pkg/logger"
	"github.com/kdrkrgz/go-url-shortener/repository"
)

func RunTasks() {
	expiredUrlsDeleteHour := conf.Get("Tasks.ExpiredUrlsDeleteHour")
	repo := repository.New()
	s := gocron.NewScheduler(time.UTC)
	//tasks
	_, err := s.Every(1).Day().At(expiredUrlsDeleteHour).Do(func() {
		DeleteExpiredUrls(repo)
	})

	if err != nil {
		log.Logger().Error(err.Error())
	}
	s.StartAsync()
}

func DeleteExpiredUrls(repo *repository.Repository) error {
	deletedUrlsCount, err := repo.DeleteShortedUrlsByDate()
	if err != nil {
		return err
	}
	log.Logger().Sugar().Infof("%v Expired urls deleted!", *deletedUrlsCount)
	return nil
}
