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

func (r *postRepo) GetList(page, limit int64) (*pb.LimitResponse, error) {
	offset := (page - 1) * limit
	var (
		posts pb.LimitResponse
		post  pb.Post
	)

	GetPost := `SELECT id, user_id, title, body FROM posts ORDER BY id OFFSET $1 LIMIT $2;`
	rows, err := r.db.Query(GetPost, offset, limit)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		err := rows.Scan(&post.Id, &post.UserId, &post.Title, &post.Body)
		if err != nil {
			return nil, err
		}
		posts.Posts = append(posts.Posts, &post)
	}

	var count int64
	CountUsersQuery := `SELECT count(*) FROM posts`
	err = r.db.QueryRow(CountUsersQuery).Scan(&count)
	if err != nil {
		return nil, err
	}
	posts.AllPosts = count
	return &posts, nil
}
