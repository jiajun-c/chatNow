package handler

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"personalService/crud"
	"personalService/dal"
	pb "personalService/pb/proto"
)

const uri = "mongodb://localhost:27017"

type User struct {
}

func (g *User) Register(ctx context.Context, req *pb.RegisterRequest, rsp *pb.RegisterResponse) error {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)
	client, _ := mongo.Connect(context.TODO(), opts)
	connect := client.Database("user").Collection("user")

	err := crud.InsertUserInfo(connect, dal.UserInfo{
		UID:      0,
		Name:     req.Name,
		Password: req.Password,
	})
	if err != nil {
		return err
	}
	return nil
}
