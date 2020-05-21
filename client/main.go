package main

import (
	pb "hello-grpc/hello"
	"log"
	"time"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	addr string = "127.0.0.1:50051"
)

func main() {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	c := pb.NewHelloServiceClient(conn)
	for {
		r, err := c.Echo(context.Background(), &pb.StringMessage{"hello"})
		if err != nil {
			log.Print(err)
		} else {
			log.Printf("client: %s", r.Value)
		}
		time.Sleep(1 * time.Second)
	}
}
