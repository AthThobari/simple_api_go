package memberships

import (
	"context"
	"errors"

	"github.com/AthThobari/simple_api_go/internal/model/memberships"
	"github.com/AthThobari/simple_api_go/pkg/jwt"
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
)

func (s *service) Login(ctx context.Context, req memberships.LoginRequest) (string, error) {
	user, err := s.membershipRepo.GetUser(ctx, req.Email, "")
	if err != nil {
		log.Error().Err(err).Msg("failed to get user")
		return "", err
	}

	if user == nil {
		return "", errors.New("email not exist")
	}

	log.Info().Msgf("Stored password: %s", user.Password)
	log.Info().Msgf("Input password: %s", req.Password)

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		log.Print("Password comparison failed:", err)
	} else {
		log.Print("Password comparison succeeded!")
	}
	
	token, err := jwt.CreateToken(user.Id, user.Username, s.cfg.Service.SecretJWT)
	if err != nil {
		return "", nil
	}
	return token, nil
}
