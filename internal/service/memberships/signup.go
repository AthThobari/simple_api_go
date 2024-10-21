package memberships

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"time"

	"github.com/AthThobari/simple_api_go/internal/model/memberships"
	"golang.org/x/crypto/bcrypt"
)

func (s *service) SignUp(ctx context.Context, req memberships.SignUpRequest) error {
	user, err := s.membershipRepo.GetUser(ctx, req.Email, req.Username, 0)
	if err != nil && err != sql.ErrNoRows {
		return err
	}
	

	if user != nil {
		return errors.New("username or email already exist")
	}

	pass, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	now := time.Now()
	model := memberships.UserModel{
		Email:     req.Email,
		Password:  string(pass),
		Username:  req.Username,
		CreatedAt: now,
		UpdatedAt: now,
		CreatedBy: req.Email,
		UpdatedBy: req.Email,
	}
	err = s.membershipRepo.CreateUser(ctx, model)
if err != nil {
    log.Println("Error inserting user to database:", err)
    return err
}
return nil

	
}

