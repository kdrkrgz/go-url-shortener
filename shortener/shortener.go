package shortener

import (
	"math/rand"
	"os"
	"strconv"
	"time"

	log "github.com/kdrkrgz/go-url-shortener/pkg/logger"
	"github.com/skip2/go-qrcode"
)

type Request struct {
	TargetUrl string `json:"target_url"`
}

type Response struct {
	ShortUrl string `json:"short_url"`
	QrCode   []byte `json:"qr_code"`
}

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

// create random number
var randomInt = func(min, max int) int {
	return min + rand.Intn(max-min)
}

func getMinAndMaxLen() (int, int) {
	min, err := strconv.Atoi(os.Getenv("ShortUrlMinLen"))
	if err != nil {
		panic(err)
	}
	max, err := strconv.Atoi(os.Getenv("ShortUrlMaxLen"))
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

func GenerateQrCode(target_url string) []byte {
	var png []byte
	png, err := qrcode.Encode(target_url, qrcode.Medium, 256)
	if err != nil {
		log.Logger().Error("Qr Generation Failed!")
	}
	return png
}
