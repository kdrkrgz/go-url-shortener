package repository

import (
	"context"
	"fmt"

	"github.com/kdrkrgz/go-url-shortener/conf"
	resolver "github.com/kdrkrgz/go-url-shortener/resolver"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Repository struct {
	UrlCollection *mongo.Collection
}

func New() *Repository {
	return &Repository{
		UrlCollection: getUrlCollection(conf.Get("DbName"), conf.Get("CollectionName")),
	}
}

// CollectionApi api a collection api interface
type CollectionApi interface {
	FindOne(ctx context.Context, filter interface{},
		opts ...*options.FindOptions) (cur *mongo.Cursor, err error)
	InsertOne(ctx context.Context, document interface{},
		opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error)
	DeleteOne(ctx context.Context, filter interface{},
		opts ...*options.DeleteOptions) (*mongo.DeleteResult, error)
}

func getUrlCollection(db, collection string) *mongo.Collection {
	fmt.Println("Connecting to MongoDB...")
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(fmt.Sprintf("%s:%s", conf.Get("DbUri"), conf.Get("DbPort"))))
	if err != nil {
		panic(err)
	}
	fmt.Println("Connection Success!")
	return client.Database(db).Collection(collection)
}

// Get Shorted Url
func (repo *Repository) FindShortedUrl(targetUrl string) (resolver.Shorten, error) {
	var shortedUrl resolver.Shorten
	ctx := context.Background()
	cur := repo.UrlCollection.FindOne(ctx, bson.M{"target_url": targetUrl})
	print(cur)
	if err := cur.Decode(&shortedUrl); err != nil {
		fmt.Printf("Shorted Url Not Found - Error: %v \n", err)
	}

	return shortedUrl, nil
}

// Insert Shorted Url
func (repo *Repository) InsertShortedUrl(shortedUrl resolver.Shorten) (*mongo.InsertOneResult, error) {
	res, err := repo.UrlCollection.InsertOne(context.Background(), shortedUrl)
	if err != nil {
		panic(err)
	}
	return res, nil
}

// Delete Shorted Url
func (repo *Repository) DeleteShortedUrl(collection CollectionApi, shortedUrl resolver.Shorten) (*mongo.DeleteResult, error) {
	res, err := collection.DeleteOne(context.Background(), bson.M{"short_url": shortedUrl.ShortUrl})
	if err != nil {
		return nil, err
	}
	return res, nil
}
