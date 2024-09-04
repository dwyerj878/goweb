package db

import (
	"context"
	"hello/config"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type USER struct {
	UserName          string `json:"user_name"`
	Password          string `json:"password"`
	FullName          string `json:"full_name"`
	PlainTextPassword string `json:"plain_text_password"`
}

func Get_user(username string) {
	mongoUrl := config.Get_config().MongoHost
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(mongoUrl))

	if err != nil {
		panic(err)
	}

	defer client.Disconnect(context.TODO())
	var result bson.M
	col := client.Database("go_app").Collection("users")
	col.FindOne(context.TODO(), bson.D{{Key: "username", Value: username}}).Decode(&result)
}

func Create_user(user *USER) {
	log.Println("Creating user ", user.UserName)
	mongoUrl := config.Get_config().MongoHost
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(mongoUrl))

	if err != nil {
		panic(err)
	}
	// TODO - encrypt plain text password int password and remove
	defer client.Disconnect(context.TODO())

	col := client.Database("go_app").Collection("users")
	result, err := col.InsertOne(
		context.TODO(),
		bson.D{
			{Key: "user_name", Value: user.UserName},
			{Key: "password", Value: user.Password},
			{Key: "full_name", Value: user.FullName},
			{Key: "plain_text_password", Value: user.PlainTextPassword}})

	if err != nil {
		log.Println(err.Error())
	}
	log.Println("Created user ", user.UserName, result.InsertedID)
}
