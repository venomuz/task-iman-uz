package storage

import (
	"github.com/jmoiron/sqlx"
	"github.com/venomuz/task-iman-uz/post-service/storage/postgres"
	"github.com/venomuz/task-iman-uz/post-service/storage/repo"
)

//IStorage ...
type IStorage interface {
	Post() repo.PostStorageI
}

type storagePg struct {
	db       *sqlx.DB
	postRepo repo.PostStorageI
}

//NewStoragePg ...
func NewStoragePg(db *sqlx.DB) *storagePg {
	return &storagePg{
		db:       db,
		postRepo: postgres.NewPostRepo(db),
	}
}

func (s storagePg) Post() repo.PostStorageI {
	return s.postRepo
}
