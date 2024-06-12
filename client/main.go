package main

import (
	"github.com/NishimuraTakuya-nt/n-tech-go-api/grpc_sample"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	c := grpc_sample.NewSampleServiceClient(conn)

	response, err := c.GetData(context.Background(), &grpc_sample.Message{Body: "送信データ"})
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}
	log.Print(response.Body)

	defer conn.Close()
}
