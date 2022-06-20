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
	limitResponse, err := s.storage.Post().GetList(req.Page, req.Limit)
	if err != nil {
		fmt.Println(err)
		s.logger.Error("Error while Getting list post from Db", l.Error(err))
		return nil, status.Error(codes.Internal, "Error getting list post")
	}
	return limitResponse, nil
}

func (s *PostService) GetById(ctx context.Context, req *pb.IdRequest) (*pb.Post, error) {
	post, err := s.storage.Post().GetById(req.UserId)
	if err != nil {
		fmt.Println(err)
		s.logger.Error("Error while getting post by id from db", l.Error(err))
		return nil, status.Error(codes.Internal, "Error getting post")
	}
	return post, nil
}
func (s *PostService) DeleteById(ctx context.Context, req *pb.IdRequest) (*pb.Ok, error) {
	stat, err := s.storage.Post().DeleteById(req.UserId)
	if err != nil {
		fmt.Println(err)
		s.logger.Error("Error while deleting post by id from db", l.Error(err))
		return nil, status.Error(codes.Internal, "Error deleting post")
	}
	return stat, nil
}

func (s *PostService) UpdateById(ctx context.Context, post *pb.Post) (*pb.Ok, error) {
	ok, err := s.storage.Post().UpdateById(post)
	if err != nil {
		fmt.Println(err)
		s.logger.Error("Error while updating post by id from db", l.Error(err))
		return nil, status.Error(codes.Internal, "Error updating post")
	}
	return ok, nil
}
