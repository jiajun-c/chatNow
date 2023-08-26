package crud

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"personalService/dal"
	"testing"
)

const uri = "mongodb://localhost:27017"

func TestInsertUserInfo(t *testing.T) {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		//fmt.Errorf("failed to connect the client")
		t.Errorf("failed to connect the client")
	}
	connect := client.Database("user").Collection("user")
	err = InsertUserInfo(connect, dal.UserInfo{
		Name:     "root",
		Password: "123456",
	})
	if err != nil {
		t.Errorf(err.Error())
	}
}
