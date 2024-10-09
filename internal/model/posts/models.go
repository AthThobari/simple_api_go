package posts

import "time"

type (
	CreatePostRequest struct {
		PostTitle   string   `json:"postTitle"`
		PostContent string   `json:"postContent"`
		PostHastag  []string `json:"postHastag"`
	}
)

type (
	PostModel struct {
		ID          int64     `db:"id"`
		UserID      int64     `db:"user_id"`
		PostTitle   string    `db:"post_title"`
		PostContent string    `db:"post_content"`
		PostHastag  string    `db:"post_hastag"`
		CreatedAt   time.Time `db:"created_at"`
		UpdatedAt   time.Time `db:"updated_at"`
		CreatedBy   string    `db:"created_by"`
		UpdatedBy   string    `db:"updated_by"`
	}
)

type (
	GetAllPostResponse struct {
		Data       []Post     `json:"data"`
		Pagination Pagination `json:"pagination"`
	}

	Post struct {
		ID          int64    `json:"id"`
		UserID      int64    `json:"userID"`
		Username    string   `json:"username"`
		PostTitle   string   `json:"postTitle"`
		PostContent string   `json:"postContent"`
		PostHastag  []string `json:"postHastag"`
		IsLiked bool `json:"isLiked"`
	}

	Pagination struct {
		Limit int `json:"limit"`
		Offset int `json:"offset"`
	}

	GetPostResponse struct {
		PostDetail Post `json:"postDetail"`
		LikeCount int `json:"likeCount"`
		Comments []Comment `json:"comments"`
	}

	Comment struct {
		ID int64 `json:"id"`
		UserID int64 `json:"user_id"`
		Username string `json:"username"`
		CommentContent string `json:"commentContent"`
	}
)

