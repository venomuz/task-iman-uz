package postgres

import (
	"github.com/jmoiron/sqlx"
	pb "github.com/venomuz/task-iman-uz/post-service/genproto"
)

type postRepo struct {
	db *sqlx.DB
}

//NewPostRepo ...
func NewPostRepo(db *sqlx.DB) *postRepo {
	return &postRepo{db: db}
}

func (r *postRepo) Create(post *pb.Post) (*pb.Post, error) {
	return nil, nil
}
