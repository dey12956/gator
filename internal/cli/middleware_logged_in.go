package cli

import (
	"context"

	"github.com/dey12956/gator/internal/database"
)

func MiddlewareLoggedIn(handler func(s *State, cmd Command, user database.User) error) func(*State, Command) error {
	return func(s *State, cmd Command) error {
		user, err := s.DB.GetUser(context.Background(), s.C.CurrentUserName)
		if err != nil {
			return err
		}
		return handler(s, cmd, user)
	}
}
