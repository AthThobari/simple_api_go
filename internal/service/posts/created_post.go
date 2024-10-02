package posts

import (
	"context"
	"strconv"
	"strings"
	"time"

	"github.com/AthThobari/simple_api_go/internal/model/posts"
	"github.com/rs/zerolog/log"
)

func (s *service) CreatePost(ctx context.Context, userID int64, req posts.CreatePostRequest) error {
	postHashtag:= strings.Join(req.PostHastag, ",")
	now:= time.Now()
	model := posts.PostModel{
		UserID: userID,
		PostTitle: req.PostTitle,
		PostContent: req.PostContent,
		PostHastag: postHashtag,
		CreatedAt: now,
		UpdatedAt: now,
		CreatedBy: strconv.FormatInt(userID, 10),
		UpdatedBy: strconv.FormatInt(userID, 10),
	}
	err:= s.postRepo.CreatePost(ctx, model) 
	if err != nil{
		log.Error().Err(err).Msg("error create post to repository")
		return err
	}
	return nil
}