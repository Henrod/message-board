package main

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/Henrod/message-board/protogen"
)

type BoardAPI struct {
	boards map[string][]*pb.Message
}

func NewBoardAPI() *BoardAPI {
	return &BoardAPI{boards: map[string][]*pb.Message{}}
}

func (b *BoardAPI) CreateMessage(
	_ context.Context,
	req *pb.CreateMessageRequest,
) (*pb.CreateMessageResponse, error) {
	userID := req.GetUserId()

	board, ok := b.boards[userID]
	if !ok {
		board = []*pb.Message{}
	}

	b.boards[userID] = append(board, req.GetMessage())

	return &pb.CreateMessageResponse{
		Message: req.GetMessage(),
	}, nil
}

func (b *BoardAPI) ListMessages(
	_ context.Context,
	req *pb.ListMessagesRequest,
) (*pb.ListMessagesResponse, error) {
	userID := req.GetUserId()
	board, ok := b.boards[userID]
	if !ok {
		return nil, status.Errorf(codes.NotFound, "user not found: %s", userID)
	}

	return &pb.ListMessagesResponse{
		Messages: board,
	}, nil
}
