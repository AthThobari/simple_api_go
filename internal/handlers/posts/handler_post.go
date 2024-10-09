package posts

import (
	"context"

	"github.com/AthThobari/simple_api_go/internal/middleware"
	"github.com/AthThobari/simple_api_go/internal/model/posts"
	"github.com/gin-gonic/gin"
)

type postService interface {
	CreatePost(ctx context.Context, userID int64, req posts.CreatePostRequest) error
	CreateComment(ctx context.Context, postID, userID int64, request posts.CreateCommentRequest) error
	UpsertUserActivity(ctx context.Context, postID, userID int64, request posts.UserActivityRequest) error
	GetAllPost(ctx context.Context, pageSize, pageIndex int) (posts.GetAllPostResponse, error)

	GetPostById(ctx context.Context, postID int64) (*posts.GetPostResponse, error)
}
type Handler struct {
	*gin.Engine

	postSvc postService
}

func NewHandler(api *gin.Engine, postSvc postService) *Handler {
	return &Handler{
		Engine:  api,
		postSvc: postSvc,
	}
}

func (h *Handler) RegisterRoute() {
	route:= h.Group("posts")
	route.Use(middleware.AuthMiddleware())

	route.POST("/create", h.CreatePost)
	route.POST("/comment/:postID", h.CreateComment)
	route.PUT("/user_activity/:postID", h.UpsertUserActivity)
	route.GET("/", h.GetAllPost)
	route.GET("/:postID", h.GetPostById)
}
