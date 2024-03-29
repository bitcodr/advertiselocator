package repositories

import (
	"errors"

	"github.com/amiraliio/advertiselocator/helpers"
	"github.com/amiraliio/advertiselocator/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//TODO transaction in mongodb for personRegister

//AuthRepository interface
type AuthInterface interface {
	PersonRegister(person *models.Person, auth *models.Auth, client *models.Client) (*models.Client, error)
	GetAuthData(authValue string) (*models.Auth, error)
	InsertClient(client *models.Client) (primitive.ObjectID, error)
}

//AuthRepository repository
type AuthRepository struct{}

//TODO this function must be ACID
//PersonRegister method
func (service *AuthRepository) PersonRegister(person *models.Person, auth *models.Auth, client *models.Client) (*models.Client, error) {
	userExist, err := checkUserExistOrNot(auth)
	if err != nil {
		return nil, err
	}
	if userExist {
		return nil, errors.New("user with the requested email exist")
	}
	//insert auth
	_, err = helpers.Mongo().InsertOne(models.AuthCollection, auth)
	if err != nil {
		return nil, err
	}
	//insert person
	_, err = helpers.Mongo().InsertOne(models.PersonCollection, person)
	if err != nil {
		return nil, err
	}
	_, err = service.InsertClient(client)
	if err != nil {
		return nil, err
	}
	return client, nil
}

//InsertClient func
func (service *AuthRepository) InsertClient(client *models.Client) (primitive.ObjectID, error) {
	filter := bson.D{
		bson.E{
			Key:   "userID",
			Value: client.UserID,
		},
		bson.E{
			Key:   "status",
			Value: models.ActiveStatus,
		},
	}
	update := bson.D{
		bson.E{
			Key: "$set",
			Value: bson.D{
				bson.E{
					Key:   "status",
					Value: models.InactiveStatus,
				},
			},
		},
	}
	_ = helpers.Mongo().FindOneAndUpdate(models.ClientCollection, filter, update)
	insertedID, err := helpers.Mongo().InsertOne(models.ClientCollection, client)
	if err != nil {
		return primitive.NilObjectID, err
	}
	return insertedID, nil
}

func checkUserExistOrNot(auth *models.Auth) (bool, error) {
	query := bson.M{"value": auth.Value, "status": models.ActiveStatus, "userType": auth.UserType, "type": auth.Type}
	var result *models.Auth
	_ = helpers.Mongo().FindOne(models.AuthCollection, query).Decode(&result)
	if result != nil && result.Value == auth.Value {
		return true, nil
	}
	return false, nil
}

//GetAuthData with auth value
func (service *AuthRepository) GetAuthData(authValue string) (auth *models.Auth, err error) {
	query := bson.M{"value": authValue, "status": models.ActiveStatus}
	if err = helpers.Mongo().FindOne(models.AuthCollection, query).Decode(&auth); err != nil {
		return nil, err
	}
	return auth, nil
}
