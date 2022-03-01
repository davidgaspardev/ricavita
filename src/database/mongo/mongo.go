package mongodb

import (
	"context"
	"fmt"
	"net/url"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type _Mongo struct {
	client   *mongo.Client
	context  *context.Context
	database *mongo.Database
	timeout  time.Duration
}

var _instance _Mongo

/// Create instance to MongoDB
func init() {
	connectMongo()
}

func connectMongo() {
	attemps := 0

CONNECTION:
	for attemps < 5 {
		uri := createUri()
		ctx := context.Background()

		client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
		if err != nil {
			continue
		}

		database := client.Database(getDatabaseName())

		collectionNames, err := database.ListCollectionNames(ctx, nil)
		if err != nil {
			client.Disconnect(ctx)
			continue
		}

		for i := 0; i < len(collections); i++ {
			collectionAlreadyExists := false
			for _, collectionName := range collectionNames {
				if collectionName == collections[i] {
					collectionAlreadyExists = true
					break
				}
			}

			if collectionAlreadyExists {
				continue
			}

			options := options.CreateCollection().SetValidator(getSchemaByCollectionName(collections[i]))
			if err := database.CreateCollection(ctx, collections[i], options); err != nil {
				continue CONNECTION
			}
		}

		_instance = _Mongo{
			client:   client,
			context:  &ctx,
			database: database,
		}
		return
	}
}

func createUri() string {
	// Getting environment variables
	username := os.Getenv("MONGO_USER")
	password := os.Getenv("MONGO_PASS")
	hostname := os.Getenv("MONGO_HOST")
	if hostname == "" {
		hostname = "localhost"
	}

	// Mounting url to mongo database
	var mongoURI string
	if username == "" && password == "" {
		mongoURI = fmt.Sprintf("mongodb://%s", hostname)
	} else {
		mongoURI = fmt.Sprintf("mongodb+srv://%s:%s@%s", username, url.QueryEscape(password), hostname)
	}

	return mongoURI
}

func getDatabaseName() string {
	database := os.Getenv("MONGO_DB")
	if database == "" {
		database = "ricavi"
	}
	return database
}
