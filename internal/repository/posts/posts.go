package posts

import (
	"context"
	"strings"

	"github.com/AthThobari/simple_api_go/internal/model/posts"
)

func (r *repository) CreatePost(ctx context.Context, model posts.PostModel) error {
	query := `INSERT INTO posts(user_id, post_title, post_content, post_hastag, created_at, updated_at, updated_by, created_by) VALUES (?, ?, ?, ?, ?, ?, ?, ?)`
	_, err := r.db.ExecContext(ctx, query, model.UserID, model.PostTitle, model.PostContent, model.PostHastag,
		model.CreatedAt, model.UpdatedAt, model.CreatedBy, model.UpdatedBy)

	if err != nil {
		return err
	}

	return nil
}

func (r *repository) GetAllPost(ctx context.Context, limit, offset int) (posts.GetAllPostResponse, error) {
	query := `SELECT p.id, p.user_id, u.username, p.post_title, p.post_content, p.post_hastag FROM posts p JOIN users u ON p.user_id = p.user_id = u.id  ORDER BY p.updated_at DESC LIMIT ? OFFSET ?`

	response := posts.GetAllPostResponse{}
	rows, err := r.db.QueryContext(ctx, query, limit, offset)
	if err != nil {
		return response, err
	}
	defer rows.Close()
	data := make([]posts.Post, 0)
	for rows.Next() {
		var (
			model    posts.PostModel
			username string
		)
		err = rows.Scan(&model.ID, &model.UserID, &username, &model.PostTitle, &model.PostContent, &model.PostHastag)
		if err != nil {
			return response, err
		}
		data = append(data, posts.Post{
			ID:          model.ID,
			UserID:      model.UserID,
			Username:    username,
			PostTitle:   model.PostTitle,
			PostContent: model.PostContent,
			PostHastag:  strings.Split(model.PostHastag, ","),
		})
	}
	response.Data =data
	response.Pagination = posts.Pagination{
		Limit : limit,
		Offset : offset,
	}
	return response, nil
}

func (r *repository) GetPostById(ctx context.Context, id int64) (*posts.Post, error) {
	query := `SELECT p.id, p.user_id, u.username, p.post_title, p.post_content, p.post_hastag, uv.is_liked FROM posts p JOIN users u ON p.user_id = u.id JOIN user_activities uv ON uv.post_id = p.id 
	WHERE p.id = ?`

	var (
		model posts.PostModel
		username string
		isLiked bool
	)
	row:=r.db.QueryRowContext(ctx, query, id)
	err:=row.Scan(&model.ID, &model.UserID, &username, &model.PostTitle, &model.PostContent, &model.PostHastag, &isLiked)
	if err != nil {
		return nil, err
	}
	return &posts.Post{
		ID: model.ID,
		UserID: model.UserID,
		Username: username,
		PostTitle: model.PostTitle,
		PostContent: model.PostContent,
		PostHastag:  strings.Split(model.PostHastag, ","),
		IsLiked: isLiked,
	}, nil
	
}