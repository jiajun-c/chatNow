package main

import (
	"fmt"
	"go-micro.dev/v4"
	"personalService/handler"
	pb "personalService/pb/proto"
)

func main() {
	service := micro.NewService(micro.Name("personService"))
	service.Init()
	err := pb.RegisterUserHandler(service.Server(), &handler.User{})
	if err != nil {
		return
	}
	if err := service.Run(); err != nil {
		fmt.Printf(err.Error())
	}
}
