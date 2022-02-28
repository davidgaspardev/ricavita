package mongodb

import (
	"context"
	"ricavi-data/src/helpers"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type _Mongo struct {
	client   *mongo.Client
	context  *context.Context
	database *mongo.Database
	timeout  time.Duration
}

/// Create instance to MongoDB
func NewMongo(uri string, databaseName string) (helpers.DataBase, error) {
	// Connect with MongoDB server by URI
	ctx := context.Background()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}

	database := client.Database(databaseName)

	collectionNames, err := database.ListCollectionNames(ctx, nil)
	if err != nil {
		return nil, err
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
			return nil, err
		}
	}

	// Return mongodb
	return &_Mongo{
		client:   client,
		context:  &ctx,
		database: database,
	}, nil
}

func (io *_Mongo) Add(collectionName string, data interface{}) (id string, err error) {
	result, err := io.database.Collection(collectionName).InsertOne(*io.context, data)
	if err != nil {
		return id, err
	}

	id = result.InsertedID.(primitive.ObjectID).Hex()
	return id, err
}

func (io *_Mongo) GetAll(collectionName string) (data []map[string]interface{}, err error) {
	cursor, err := io.database.Collection(collectionName).Find(*io.context, bson.M{})
	if err != nil {
		return data, err
	}
	defer cursor.Close(*io.context)
	data = make([]map[string]interface{}, 0)

	for cursor.Next(*io.context) {
		var cursorData bson.M
		var err = cursor.Decode(&cursorData)
		if err != nil {
			return data, err
		}
		data = append(data, cursorData)
	}

	return data, err
}

func (io *_Mongo) Close() error {
	if err := io.client.Disconnect(*io.context); err != nil {
		return err
	}
	return nil
}
