package resolver

import (
	"time"

	"github.com/go-playground/validator"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Shorten struct {
	ID             primitive.ObjectID `bson:"_id,omitempty"`
	ExpirationDate time.Time          `json:"expiration_date,omitempty"`
	TargetUrl      string             `bson:"target_url,omitempty"`
	ShortUrl       string             `bson:"short_url,omitempty"`
	Clicks         int64              `bson:"clicks"`
	CreatedAt      time.Time          `bson:"created_at"`
	LastHitAt      time.Time          `bson:"last_hit_at,omitempty"`
}

type Response struct {
	ShortUrl       string        `json:"short_url"`
	ExpirationDate time.Duration `json:"expiration_date,omitempty"`
}

var validate = validator.New()

type ErrorResponse struct {
	Field string `json:"field"`
	Value string `json:"value,omitempty"`
	Tag   string `json:"tag"`
}

func ValidateStruct[T any](payload T) []*ErrorResponse {
	err := validate.Struct(payload)
	if err == nil {
		return nil
	}
	var errors []*ErrorResponse
	for _, err := range err.(validator.ValidationErrors) {
		errors = append(errors, &ErrorResponse{
			Field: err.StructNamespace(),
			Value: err.Value().(string),
			Tag:   err.Tag(),
		})
	}
	return errors
}
