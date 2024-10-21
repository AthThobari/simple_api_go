package memberships

import (
	"context"
	"database/sql"
	"time"

	"github.com/AthThobari/simple_api_go/internal/model/memberships"
)

func (r *repository) InsertRefreshToken(ctx context.Context, model memberships.RefreshTokenModel) error {
	query := `INSERT INTO refresh_tokens (user_id, refresh_token, expired_at, created_at, updated_at, created_by, updated_by) VALUES (?, ?, ?, ?, ?, ?, ?)`
	_, err := r.db.ExecContext(ctx, query, &model.UserID, &model.RefreshToken, &model.Expired_at, &model.CreatedAt, &model.UpdatedAt, &model.CreatedBy, &model.UpdatedBy)

	if err != nil {
		return err
	}
	return nil
}

func (r *repository) GetRefreshToken(ctx context.Context, userID int64, now time.Time) (*memberships.RefreshTokenModel, error){
	query := `SELECT id, user_id, refresh_token, expired_at, created_at, updated_at, created_by, updated_by FROM refresh_tokens WHERE user_id = ? AND expired_at >= ?`
	var resp memberships.RefreshTokenModel
	row:=r.db.QueryRowContext(ctx, query, userID, now)
	err:=row.Scan(&resp.ID, &resp.UserID, &resp.RefreshToken, &resp.Expired_at, &resp.CreatedAt, &resp.UpdatedAt, &resp.CreatedBy, &resp.UpdatedBy)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &resp, nil
}
