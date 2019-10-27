package helpers

import (
	"context"
	"time"

	"github.com/amiraliio/advertiselocator/configs"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//TODO improvement must be just interface not function

//Mongo build
func Mongo() MongoInterface {
	return new(mongoService)
}

//MongoInterface interface
type MongoInterface interface {
	InsertOne(collectionName string, object interface{}) (primitive.ObjectID, error)
}

type mongoService struct{}

//InsertOne document in mongo
func (service *mongoService) InsertOne(collectionName string, object interface{}) (primitive.ObjectID, error) {
	db := configs.DB().Collection(collectionName)
	context, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	entityModel, err := Flatten(object)
	if err != nil {
		return primitive.NilObjectID, err
	}
	result, err := db.InsertOne(context, entityModel)
	if err != nil {
		return primitive.NilObjectID, err
	}
	return result.InsertedID.(primitive.ObjectID), nil
}
