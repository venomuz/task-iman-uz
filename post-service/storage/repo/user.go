package repo

import (
	pb "github.com/venomuz/task-iman-uz/post-service/genproto"
)

//PostStorageI ...
type PostStorageI interface {
	Create(*pb.Post) (*pb.Post, error)
}
