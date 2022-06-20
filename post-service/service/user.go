package service

import (
	"context"
	"github.com/jmoiron/sqlx"
	pb "github.com/venomuz/task-iman-uz/post-service/genproto"
	l "github.com/venomuz/task-iman-uz/post-service/pkg/logger"
	"github.com/venomuz/task-iman-uz/post-service/storage"
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

func (s *PostService) GetList(ctx context.Context, request *pb.LimitRequest) (*pb.LimitResponse, error) {
	//TODO implement me
	panic("implement me")
}
