package service

import (
	"context"

	pb "review-business/api/business/v1"
)

type BusinessService struct {
	pb.UnimplementedBusinessServer
}

func NewBusinessService() *BusinessService {
	return &BusinessService{}
}

func (s *BusinessService) ReplyReview(ctx context.Context, req *pb.ReplyReviewRequest) (*pb.ReplyReviewReply, error) {
	return &pb.ReplyReviewReply{}, nil
}
