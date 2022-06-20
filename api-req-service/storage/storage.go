package storage

import (
	"github.com/jmoiron/sqlx"
	"github.com/venomuz/task-iman-uz/api-req-service/storage/postgres"
	"github.com/venomuz/task-iman-uz/api-req-service/storage/repo"
)

//IStorage ...
type IStorage interface {
	Req() repo.ReqStorageI
}

type storagePg struct {
	db      *sqlx.DB
	reqRepo repo.ReqStorageI
}

//NewStoragePg ...
func NewStoragePg(db *sqlx.DB) *storagePg {
	return &storagePg{
		db:      db,
		reqRepo: postgres.NewReqRepo(db),
	}
}

func (s storagePg) Req() repo.ReqStorageI {
	return s.reqRepo
}
