package posts

import (
	"context"

	"github.com/AthThobari/simple_api_go/internal/model/posts"
)

func (r *repository) CreatePost(ctx context.Context, model posts.PostModel) error {
	query:= `INSERT INTO posts(user_id, post_title, post_content, post_hastag, created_at, updated_at, updated_by, created_by) VALUES (?, ?, ?, ?, ?, ?, ?, ?)`
	_,err:=r.db.ExecContext(ctx, query, model.UserID, model.PostTitle, model.PostContent, model.PostHastag,
	model.CreatedAt, model.UpdatedAt, model.CreatedBy, model.UpdatedBy)

	if err != nil {
		return err
	}

	return nil
}