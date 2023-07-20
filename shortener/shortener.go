package shortener

import "fmt"

type Request struct {
	TargetUrl string `json:"target_url"`
}

func GenerateShortUrl(targetUrl string) string {
	return fmt.Sprintf("http://localhost:8000/shorted --> %v", targetUrl)
}
