package database

import mongodb "ricavi-data/src/database/mongo"

func GetDatabase() Database {
	return mongodb.GetMongo()
}

type Database interface {
	Add(name string, data interface{}) (id string, err error)
	GetAll(name string) (data []map[string]interface{}, err error)
}
