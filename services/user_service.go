package services

import (
	"context"
	"errors"
	"time"

	"github.com/Jellicle-Cats/user-service/models"
	"github.com/Jellicle-Cats/user-service/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserService struct {
	Collection *mongo.Collection
	Ctx        context.Context
}

func NewUserService(collection *mongo.Collection, ctx context.Context) *UserService {
	return &UserService{
		Collection: collection,
		Ctx:        ctx,
	}
}

func (userService UserService) FindUserById(id string) (*models.DBResponse, error) {
	oid, _ := primitive.ObjectIDFromHex(id)

	var user *models.DBResponse

	query := bson.M{"_id": oid}
	err := userService.Collection.FindOne(userService.Ctx, query).Decode(&user)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return &models.DBResponse{}, err
		}
		return nil, err
	}

	return user, nil
}

func (userService UserService) FindUserByEmail(email string) (*models.DBResponse, error) {
	var user *models.DBResponse

	query := bson.M{"email": email}
	err := userService.Collection.FindOne(userService.Ctx, query).Decode(&user)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return &models.DBResponse{}, err
		}
		return nil, err
	}

	return user, nil
}

func (userService UserService) UpsertUser(email string, data *models.User) (*models.DBResponse, error) {
	doc, err := utils.ToDoc(data)
	if err != nil {
		return nil, err
	}

	opts := options.FindOneAndUpdate().SetUpsert(true).SetReturnDocument(1)
	query := bson.D{{Key: "email", Value: email}}
	update := bson.D{{Key: "$set", Value: doc}, {Key: "$setOnInsert", Value: bson.D{{Key: "created_at", Value: time.Now()}}}}
	res := userService.Collection.FindOneAndUpdate(userService.Ctx, query, update, opts)

	var updatedPost *models.DBResponse

	if err := res.Decode(&updatedPost); err != nil {
		return nil, errors.New("no post with that Id exists")
	}

	return updatedPost, nil
}
