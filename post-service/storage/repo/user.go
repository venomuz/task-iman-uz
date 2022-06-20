package repo

import (
	pb "github.com/venomuz/task-iman-uz/post-service/genproto"
)

//PostStorageI ...
type PostStorageI interface {
	GetList(page, limit int64) (*pb.LimitResponse, error)
	GetById(UserId int64) (*pb.Post, error)
	DeleteById(UserId int64) (*pb.Ok, error)
	UpdateById(post *pb.Post) (*pb.Ok, error)
}
