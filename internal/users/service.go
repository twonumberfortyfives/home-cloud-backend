package users

import (
	"context"

	"github.com/twonumberfortyfives/home-cloud-backend/internal/utils"
)

type UserService struct{}

func (s *UserService) GetCurrentUser(ctx context.Context) UserDto {
	user, _ := ctx.Value(utils.UserContextKey).(UserDto)
	return user
}
