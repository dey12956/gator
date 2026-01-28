package cli

import (
	"context"
	"errors"
	"fmt"
)

func HandlerUsers(s *State, cmd Command) error {
	if len(cmd.Args) != 0 {
		return errors.New("you should not enter arguments here")
	}

	users, err := s.DB.GetUsers(context.Background())
	if err != nil {
		return err
	}

	for _, user := range users {
		if s.C.CurrentUserName == user.Name {
			fmt.Printf("* %s (current)\n", user)
		} else {
			fmt.Printf("* %s\n", user)
		}
	}

	return nil
}
