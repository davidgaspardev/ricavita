package mongodb

import "go.mongodb.org/mongo-driver/bson"

var collections = []string{
	recipeCollectionName,
}

const recipeCollectionName = "recipe"

func getSchemaByCollectionName(collectionName string) bson.M {
	var schema bson.M
	switch collectionName {
	case recipeCollectionName:
		schema = recipeSchema
	default:
		panic("schema don't exists")
	}

	return bson.M{
		"$jsonSchema": schema,
	}
}

var recipeSchema = bson.M{
	"bsonType":             "object",
	"required":             []string{"name", "cover", "category", "ingredients", "instructions"},
	"additionalProperties": false,
	"properties": bson.M{
		"_id": bson.M{
			"bsonType": "objectId",
		},
		"name": bson.M{
			"bsonType": "string",
		},
		"cover": bson.M{
			"bsonType": "string",
		},
		"category": bson.M{
			"enum": []string{"massa", "sobremesa", "bolo", "torta", "vegetariano", "vegano", "carnes", "pratico", "Drinks"},
		},
		"ingredients": bson.M{
			"bsonType":    []string{"array"},
			"uniqueItems": true,
			"items":       ingredientSchema,
		},
		"instructions": bson.M{
			"bsonType":    []string{"array"},
			"uniqueItems": true,
			"items":       instructionSchema,
		},
	},
}

var ingredientSchema = bson.M{
	"bsonType":             "object",
	"required":             []string{"name", "quantity", "unit"},
	"additionalProperties": false,
	"properties": bson.M{
		"name": bson.M{
			"bsonType": "string",
		},
		"quantity": bson.M{
			"bsonType": "double",
		},
		"unit": bson.M{
			"bsonType": "string",
		},
	},
}

var instructionSchema = bson.M{
	"bsonType":            "object",
	"required":            []string{"order", "description"},
	"additinalProperties": false,
	"properties": bson.M{
		"order": bson.M{
			"bsonType": "int",
		},
		"description": bson.M{
			"bsonType": "string",
		},
	},
}
