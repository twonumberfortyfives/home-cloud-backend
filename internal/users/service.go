package users

type UserService struct{}

func (s *UserService) GetCurrentUser() UserDto {
	return UserDto{
	}
}
