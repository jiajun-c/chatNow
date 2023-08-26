package crud

import (
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"personalService/dal"
)

var Client mongo.Client

func InsertUserInfo(collection *mongo.Collection, userInfo dal.UserInfo) error {
	// check if the user already exist
	filter := bson.D{{"name", userInfo.Name}}
	cursor, err := collection.Find(context.TODO(), filter)
	var results []dal.UserInfo
	cursor.All(context.TODO(), &results)
	if len(results) > 0 {
		return errors.New("the user already exist")
	}
	// insert user info
	_, err = collection.InsertOne(context.TODO(), userInfo)
	if err != nil {
		fmt.Errorf(err.Error())
		return err
	}
	return nil
}
