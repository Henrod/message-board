package main

import (
	"context"
	"log"

	pb "github.com/Henrod/message-board/protogen"

	"google.golang.org/grpc"
)

func main() {
	dial, err := grpc.Dial("localhost:8888", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	client := pb.NewBoardAPIClient(dial)

	_, err = client.CreateMessage(context.Background(), &pb.CreateMessageRequest{
		Message: &pb.Message{Body: "msg1"},
		UserId:  "1",
	})
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.ListMessages(context.Background(), &pb.ListMessagesRequest{
		UserId: "1",
	})
	if err != nil {
		log.Fatal(err)
	}

	log.Print(resp)
}
