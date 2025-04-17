package user

import (
	"context"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type Repository interface {
	GetUserById(ctx context.Context, id string) (*User, error)
	GetUserByWalletAddress(ctx context.Context, walletAddress string) (*User, error)
	DeleteUserById(ctx context.Context, id string) error
	DeleteUserByWalletAddress(ctx context.Context, walletAddress string) error

	Create(ctx context.Context, user *User) (*User, error)
}

type repositoryImpl struct {
	collection *mongo.Collection
}

func NewRepository(db *mongo.Database) Repository {
	return &repositoryImpl{
		collection: db.Collection("users"),
	}
}

func (r *repositoryImpl) GetUserById(ctx context.Context, id string) (*User, error) {

	var u User
	err := r.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&u)
	if err != nil {
		return nil, err
	}
	return &u, nil
}
func (r *repositoryImpl) GetUserByWalletAddress(ctx context.Context, walletAddress string) (*User, error) {

	var u User
	err := r.collection.FindOne(ctx, bson.M{"wallet_address": walletAddress}).Decode(&u)
	if err != nil {
		return nil, err
	}
	return &u, nil
}
func (r *repositoryImpl) DeleteUserById(ctx context.Context, id string) error {
	_, err := r.collection.DeleteOne(ctx, bson.M{"_id": id})
	return err
}

func (r *repositoryImpl) DeleteUserByWalletAddress(ctx context.Context, walletAddress string) error {
	_, err := r.collection.DeleteOne(ctx, bson.M{"wallet_address": walletAddress})
	return err
}

func (r *repositoryImpl) Create(ctx context.Context, user *User) (*User, error) {
	_, err := r.collection.InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}
	return user, nil
}
