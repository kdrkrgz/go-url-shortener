package repository

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"time"

	resolver "github.com/kdrkrgz/go-url-shortener/resolver"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoRepository struct {
	UrlCollection *mongo.Collection
}

func NewMongoRepository() *MongoRepository {
	col := &MongoRepository{
		UrlCollection: getUrlCollection(os.Getenv("DbName"), os.Getenv("CollectionName")),
	}
	// create unique index for shorted url
	col.UrlCollection.Indexes().CreateOne(context.TODO(), mongo.IndexModel{
		Keys: bson.M{
			"short_url": 1,
		},
		Options: options.Index().SetUnique(true),
	})
	return col
}

func getUrlCollection(db, collection string) *mongo.Collection {
	fmt.Println("Connecting to MongoDB...")
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(fmt.Sprintf("%s:%s", os.Getenv("DbUri"), os.Getenv("DbPort"))))
	if err != nil {
		panic(err)
	}
	fmt.Println("Connection Success!")
	return client.Database(db).Collection(collection)
}

func (repo *MongoRepository) FindUrl(field, url string) (*resolver.ShortUrl, error) {

	var shortUrl = new(resolver.ShortUrl)
	ctx := context.Background()

	cur := repo.UrlCollection.FindOne(ctx, bson.M{field: url})
	if err := cur.Decode(&shortUrl); err != nil {

		return nil, err

	}
	return shortUrl, nil
}

func (repo *MongoRepository) InsertUrl(shortedUrl resolver.ShortUrl) (*mongo.InsertOneResult, error) {
	res, err := repo.UrlCollection.InsertOne(context.Background(), shortedUrl)
	if err != nil {
		err = fmt.Errorf("an error occured insert url to db - Error: %v", err)
		return nil, err
	}
	return res, nil
}

func (repo *MongoRepository) DeleteUrlsByDate() (*mongo.DeleteResult, error) {
	expireTime, _ := strconv.Atoi(os.Getenv("UrlExpirationTime"))
	minutesAgo := time.Now().UTC().Add(-time.Duration(expireTime) * time.Minute)
	fmt.Printf("Time Now: %v\n", time.Now())
	res, err := repo.UrlCollection.DeleteMany(context.Background(), bson.M{"created_at": bson.M{"$lte": minutesAgo}})
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (repo *MongoRepository) DropCollection() {
	repo.UrlCollection.DeleteMany(context.Background(), bson.M{})

}
