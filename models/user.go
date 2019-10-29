package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)


//TODO order struct field from high to low field type size

//////////// user collection ///////////////

//PersonCollection collection
const PersonCollection string = "persons"

//////////// user types ///////////////

//PersonUserType user type
const PersonUserType string = "PERSON"

//AdminUserType user type
const AdminUserType string = "ADMIN"


//BaseUser model
type BaseUser struct {
	Base     ",inline"
	UserID   primitive.ObjectID `json:"userID" bson:"userID"`
	UserType string             `json:"userType" bson:"userType"`
	IP       string             `json:"ip" bson:"ip"`
}


//Person model
type Person struct {
	BaseUser  ",inline"
	Location  Location `json:"location" bson:"location"`
	Avatar    Image    `json:"avatar" bson:"avatar"`
	FirstName string   `json:"firstName" bson:"firstName"`
	LastName  string   `json:"lastName" bson:"lastName"`
	CellPhone string   `json:"cellPhone" bson:"cellPhone"`
	Email     string   `json:"email" bson:"email"`
	Radius    uint16   `json:"radius" bson:"radius"`
}
