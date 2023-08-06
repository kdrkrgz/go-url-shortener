package pkg

import (
	"os"
	"time"

	"github.com/go-co-op/gocron"
	log "github.com/kdrkrgz/go-url-shortener/pkg/logger"
	"github.com/kdrkrgz/go-url-shortener/resolver"
)

func RunTasks(repo resolver.DbRepository) {
	expiredUrlsDeleteHour := os.Getenv("ExpiredUrlsDeleteHour")
	s := gocron.NewScheduler(time.UTC)
	//tasks
	_, err := s.Every(1).Day().At(expiredUrlsDeleteHour).Do(func() {
		DeleteExpiredUrlsFromDb(repo)
	})

	if err != nil {
		log.Logger().Error(err.Error())
	}
	s.StartAsync()
}

func DeleteExpiredUrlsFromDb(repo resolver.DbRepository) error {
	deletedUrlsCount, err := repo.DeleteUrlsByDate()
	if err != nil {
		return err
	}
	log.Logger().Sugar().Infof("%v Expired urls deleted!", *deletedUrlsCount)
	return nil
}
