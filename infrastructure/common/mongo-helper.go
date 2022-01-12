package common

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"sync"
	"time"
)

var once sync.Once
var instance *MongoHelper

type MongoHelper struct {
	db *mongo.Database
}

func NewMongoHelper(uri, database string, timeout int) *MongoHelper {
	once.Do(func() {
		ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
		defer cancel()

		if client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri)); err != nil {
			panic(err)
		} else {
			instance = &MongoHelper{
				db: client.Database(database),
			}
		}
	})
	return instance
}

func (m *MongoHelper) GetCollection(name string) *mongo.Collection {
	return m.db.Collection(name)
}
