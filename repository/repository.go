package repository

import (
	"github.com/kdrkrgz/go-url-shortener/resolver"
)

type AppRepository struct {
	DbRepository    resolver.DbRepository
	CacheRepository resolver.CacheRepository
}

// // func New() *Repository {
// // 	col := &Repository{
// // 		UrlCollection: getUrlCollection(os.Getenv("DbName"), os.Getenv("CollectionName")),
// // 		RedisCache:    getUrlCache(os.Getenv("RedisHost"), os.Getenv("RedisPort")),
// // 	}
// // 	// create unique index for shorted url
// // 	col.UrlCollection.Indexes().CreateOne(context.Background(), mongo.IndexModel{
// // 		Keys: bson.M{
// // 			"short_url": 1,
// // 		},
// // 		Options: options.Index().SetUnique(true),
// // 	})
// // 	return col
// // }

// // CollectionApi api a collection api interface
// type CollectionApi interface {
// 	FindOne(ctx context.Context, filter interface{},
// 		opts ...*options.FindOptions) (cur *mongo.Cursor, err error)
// 	InsertOne(ctx context.Context, document interface{},
// 		opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error)
// 	DeleteOne(ctx context.Context, filter interface{},
// 		opts ...*options.DeleteOptions) (*mongo.DeleteResult, error)
// }

// // func getUrlCollection(db, collection string) *mongo.Collection {
// // 	fmt.Println("Connecting to MongoDB...")
// // 	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(fmt.Sprintf("%s:%s", os.Getenv("DbUri"), os.Getenv("DbPort"))))
// // 	if err != nil {
// // 		panic(err)
// // 	}
// // 	fmt.Println("Connection Success!")
// // 	return client.Database(db).Collection(collection)
// // }

// // func getUrlCache(host, port string) *redis.Client {
// // 	fmt.Println("Connecting to Redis...")
// // 	client := redis.NewClient(&redis.Options{
// // 		Addr:     fmt.Sprintf("%s:%s", host, port),
// // 		Password: os.Getenv("RedisPassword"),
// // 		DB:       0,
// // 	})
// // 	fmt.Println("Connection Success!")
// // 	return client
// // }

// func (repo *Repository) FindUrl(field, url string) (*string, error) {
// 	resFromCache, err := repo.FindUrlFromCache(url)
// 	if err != nil {
// 		log.Logger().Sugar().Infof("Url Not Found on Cache - Url: %s", url)
// 	}
// 	if resFromCache != nil {
// 		return resFromCache, nil
// 	}
// 	resFromDb, err := repo.FindUrlFromDb(field, url)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &resFromDb.TargetUrl, nil
// }

// // Get Shorted or Target Url From Db
// func (repo *Repository) FindUrlFromDb(field, url string) (*resolver.ShortUrl, error) {
// 	var shortUrl = new(resolver.ShortUrl)
// 	ctx := context.Background()

// 	cur := repo.UrlCollection.FindOne(ctx, bson.M{field: url})

// 	if err := cur.Decode(&shortUrl); err != nil {
// 		return nil, err
// 	}
// 	if err := repo.setUrlToRedis(shortUrl.ShortUrl, shortUrl.TargetUrl); err != nil {
// 		log.Logger().Sugar().Errorf("An error occured insert url to redis - Error: %v", err)
// 	}
// 	return shortUrl, nil
// }

// // Insert Shorted Url
// func (repo *Repository) InsertShortedUrl(shortedUrl resolver.ShortUrl) (*mongo.InsertOneResult, error) {
// 	res, err := repo.UrlCollection.InsertOne(context.Background(), shortedUrl)
// 	if err != nil {
// 		log.Logger().Sugar().Errorf("An error occured insert url to db - Error: %v", err)
// 		return nil, err
// 	}
// 	if err := repo.setUrlToRedis(shortedUrl.TargetUrl, shortedUrl.ShortUrl); err != nil {
// 		log.Logger().Sugar().Errorf("An error occured insert url to redis - Error: %v", err)
// 		return nil, err
// 	}

// 	return res, nil
// }

// // Get Shorted or Target Url From Cache
// // func (repo *Repository) FindUrlFromCache(url string) (*string, error) {
// // 	cur, err := repo.RedisCache.Get(url).Bytes()
// // 	if err != nil {
// // 		return nil, err
// // 	}
// // 	res := string(cur)
// // 	return &res, nil
// // }

// func (repo *Repository) setUrlToRedis(key, val string) error {
// 	cacheUrlExpiration, err := strconv.Atoi(os.Getenv("CacheUrlExpiration"))
// 	if err != nil {
// 		return err
// 	}
// 	repo.RedisCache.Set(key, val, time.Duration(cacheUrlExpiration)*time.Minute)
// 	return nil
// }

// // Delete Shorted Urls With Given Duration
// func (repo *Repository) DeleteShortedUrlsByDate() (*mongo.DeleteResult, error) {
// 	expireTime, _ := strconv.Atoi(os.Getenv("UrlExpirationTime"))
// 	minutesAgo := time.Now().UTC().Add(-time.Duration(expireTime) * time.Minute)
// 	fmt.Printf("Time Now: %v\n", time.Now())
// 	res, err := repo.UrlCollection.DeleteMany(context.Background(), bson.M{"created_at": bson.M{"$lte": minutesAgo}})
// 	if err != nil {
// 		return nil, err
// 	}
// 	return res, nil
// }
