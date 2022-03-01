package mongodb

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetMongo() *_Mongo {
	return &_instance
}

func (mongo *_Mongo) Add(collectionName string, data interface{}) (id string, err error) {
	result, err := mongo.database.Collection(collectionName).InsertOne(*mongo.context, data)
	if err != nil {
		return id, err
	}

	id = result.InsertedID.(primitive.ObjectID).Hex()
	return id, err
}

func (mongo *_Mongo) GetAll(collectionName string) (data []map[string]interface{}, err error) {
	cursor, err := mongo.database.Collection(collectionName).Find(*mongo.context, bson.M{})
	if err != nil {
		return data, err
	}
	defer cursor.Close(*mongo.context)
	data = make([]map[string]interface{}, 0)

	for cursor.Next(*mongo.context) {
		var cursorData bson.M
		var err = cursor.Decode(&cursorData)
		if err != nil {
			return data, err
		}
		data = append(data, cursorData)
	}

	return data, err
}

func (mongo *_Mongo) Close() error {
	if err := mongo.client.Disconnect(*mongo.context); err != nil {
		return err
	}
	return nil
}
