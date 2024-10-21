package memberships

import (
	"context"
	"errors"
	"strconv"
	"time"

	"github.com/AthThobari/simple_api_go/internal/model/memberships"
	"github.com/AthThobari/simple_api_go/pkg/jwt"
	tokenUtil "github.com/AthThobari/simple_api_go/pkg/token"
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
)

func (s *service) Login(ctx context.Context, req memberships.LoginRequest) (string, string, error) {
	user, err := s.membershipRepo.GetUser(ctx, req.Email, "", 0)
	if err != nil {
		log.Error().Err(err).Msg("failed to get user")
		return "","", err
	}

	if user == nil {
		return "","", errors.New("email not exist")
	}

	log.Info().Msgf("Stored password: %s", user.Password)
	log.Info().Msgf("Input password: %s", req.Password)

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		log.Print("Password comparison failed:", err)
		return "","", errors.New("email not exist")
	} else {
		log.Print("Password comparison succeeded!")
	}
	
	token, err := jwt.CreateToken(user.Id, user.Username, s.cfg.Service.SecretJWT)
	if err != nil {
		return "","", nil
	}

	existingRefreshToken, err := s.membershipRepo.GetRefreshToken(ctx, user.Id, time.Now())
	if err != nil {
		log.Error().Err(err).Msg("error get latest token from database")
		return "","", err
	}

	if existingRefreshToken != nil {
		return token, existingRefreshToken.RefreshToken, nil
	}

	refreshToken := tokenUtil.GenerateRefreshToken()
	if refreshToken == "" {
		return token, "", errors.New("failed to generate refresh token")
	}

	err = s.membershipRepo.InsertRefreshToken(ctx, memberships.RefreshTokenModel{
		UserID: user.Id,
		RefreshToken: refreshToken,
		Expired_at: time.Now().Add(10 * 24 * time.Hour),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		CreatedBy: strconv.FormatInt(user.Id, 10),
		UpdatedBy : strconv.FormatInt(user.Id, 10),

	})

	if err != nil {
		log.Error().Err(err).Msg("error inserting refresh token to database")
		return token, refreshToken, err
	}

	return token,refreshToken, nil
}
