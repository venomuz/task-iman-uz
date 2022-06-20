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

	GetPostQuery := `SELECT id, user_id, title, body FROM posts ORDER BY id OFFSET $1 LIMIT $2;`
	rows, err := r.db.Query(GetPostQuery, offset, limit)
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
func (r *postRepo) GetById(UserId int64) (*pb.Post, error) {
	post := pb.Post{}
	GetByIdQuery := `SELECT id, user_id, title, body FROM posts WHERE user_id = $1'`
	err := r.db.QueryRow(GetByIdQuery, UserId).Scan(&post.Id, &post.UserId, &post.Title, &post.Body)
	if err != nil {
		return nil, err
	}
	return &post, nil
}
func (r *postRepo) DeleteById(UserId int64) (*pb.Ok, error) {
	DeleteQuery := `DELETE FROM posts WHERE user_id = $1`
	_, err := r.db.Exec(DeleteQuery, UserId)
	if err != nil {
		return nil, err
	}
	return &pb.Ok{Status: true}, nil
}

func (r *postRepo) UpdateById(post *pb.Post) (*pb.Ok, error) {
	UpdateQuery := `UPDATE posts SET title = $1, body = $2 WHERE user_id = $3`
	_, err := r.db.Exec(UpdateQuery, post.Title, post.Body, post.UserId)
	if err != nil {
		return nil, err
	}
	return &pb.Ok{Status: true}, nil
}
