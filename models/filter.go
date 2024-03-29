package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//BaseFilter filter
type BaseFilter struct {
	UserID    primitive.ObjectID `json:"userID" bson:"userID"`
	ID        primitive.ObjectID `json:"id" bson:"_id"`
	Sort      string             `json:"sort" query:"sort" bson:"sort"`
	LastIndex string             `json:"lastIndex" query:"lastIndex" bson:"lastIndex"`
	Page      int64              `json:"page" query:"page" bson:"page"`
	Limit     int64              `json:"limit" query:"limit" bson:"limit"`
	NextPage  int64              `json:"nextPage" query:"nextPage" bson:"nextPage"`
}

//AdvertiseFilter
type AdvertiseFilter struct {
	BaseFilter ",inline"
	Tags       []*Tag `json:"tags" query:"tags" bson:"tags"`
}
