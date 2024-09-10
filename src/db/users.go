package db

import (
	"context"
	"errors"
	"hello/config"
	"log"

	"golang.org/x/crypto/bcrypt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type USER struct {
	UserName          string `json:"user_name" bson:"user_name"`
	Password          string `json:"-" bson:"password"`
	FullName          string `json:"full_name" bson:"full_name"`
	PlainTextPassword string `json:"plain_text_password" bson:"-"`
}

func Get_user(username string) (USER, error) {
	mongoUrl := config.Get_config().MongoHost
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(mongoUrl))

	if err != nil {
		panic(err)
	}

	defer client.Disconnect(context.TODO())
	var user USER
	col := client.Database("go_app").Collection("users")
	err = col.FindOne(context.TODO(), bson.D{{Key: "user_name", Value: username}}).Decode(&user)
	return user, err
}

func Create_user(user *USER) (USER, error) {
	log.Println("Creating user ", user.UserName)
	mongoUrl := config.Get_config().MongoHost
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(mongoUrl))
	if err != nil {
		panic(err)
	}

	col := client.Database("go_app").Collection("users")

	var found USER
	err = col.FindOne(context.TODO(), bson.D{{Key: "user_name", Value: user.UserName}}).Decode(&found)
	if err == nil {
		return *user, errors.New("user exists")
	}

	bytes, err := bcrypt.GenerateFromPassword([]byte(user.PlainTextPassword), 14)
	if err != nil {
		panic(err)
	}

	defer client.Disconnect(context.TODO())

	result, err := col.InsertOne(
		context.TODO(),
		bson.D{
			{Key: "user_name", Value: user.UserName},
			{Key: "full_name", Value: user.FullName},
			{Key: "password", Value: bytes}})

	if err != nil {
		log.Println(err.Error())
	}
	log.Println("Created user ", user.UserName, result.InsertedID)
	return *user, nil
}
