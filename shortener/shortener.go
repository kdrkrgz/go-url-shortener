package shortener

import (
	"math/rand"
	"strconv"
	"time"

	"github.com/kdrkrgz/go-url-shortener/conf"
)

type Request struct {
	TargetUrl string `json:"target_url"`
}

var shortUrlMinLen = conf.Get("App.ShortUrlMinLen")
var shortUrlMaxLen = conf.Get("App.ShortUrlMaxLen")

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

// create random number
var randomInt = func(min, max int) int {
	return min + rand.Intn(max-min)
}

func getMinAndMaxLen() (int, int) {
	min, err := strconv.Atoi(shortUrlMinLen)
	if err != nil {
		panic(err)
	}
	max, err := strconv.Atoi(shortUrlMaxLen)
	if err != nil {
		panic(err)
	}
	return min, max
}

func GenerateShortUrl() string {
	b := make([]byte, randomInt(getMinAndMaxLen()))
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}
