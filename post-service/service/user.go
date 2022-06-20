package service

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	pb "github.com/venomuz/task-iman-uz/post-service/genproto"
	l "github.com/venomuz/task-iman-uz/post-service/pkg/logger"
	"github.com/venomuz/task-iman-uz/post-service/storage"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

//PostService ...
type PostService struct {
	storage storage.IStorage
	logger  l.Logger
}

//NewPostService ...
func NewPostService(db *sqlx.DB, log l.Logger) *PostService {
	return &PostService{
		storage: storage.NewStoragePg(db),
		logger:  log,
	}
}

func (s *PostService) GetList(ctx context.Context, req *pb.LimitRequest) (*pb.LimitResponse, error) {
	posts, err := s.storage.Post().GetList(req.Page, req.Limit)
	if err != nil {
		fmt.Println(err)
		s.logger.Error("Error while getting post info", l.Error(err))
		return nil, status.Error(codes.Internal, "Error insert post")
	}
	return posts, nil
}

func (s *PostService) GetById(ctx context.Context, req *pb.IdRequest) (*pb.Post, error) {
	post, err := s.storage.Post().GetById(req.UserId)
	if err != nil {
		fmt.Println(err)
		s.logger.Error("Error while getting post info", l.Error(err))
		return nil, status.Error(codes.Internal, "Error insert post")
	}
	return post, nil
}
