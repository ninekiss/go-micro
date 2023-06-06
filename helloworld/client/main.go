package main

import (
	"context"
	"log"

	pb "github.com/ninekiss/go-micro/helloworld/proto"
	micro "go-micro.dev/v4"
)

func main() {
	serivce := micro.NewService(micro.Name("greeter.client"))

	serivce.Init()

	greeter := pb.NewGreeterService("greeter", serivce.Client())

	rsp, err := greeter.Hello(context.Background(), &pb.Request{Name: "Kean"})
	if err != nil {
		log.Fatal(err)
	}

	log.Println(rsp.Message)
}
