package cli

import (
	"context"
	"errors"
)

func HandlerReset(s *State, cmd Command) error {
	if len(cmd.Args) != 0 {
		return errors.New("you should not enter arguments")
	}

	err := s.DB.DeleteAllUsers(context.Background())
	if err != nil {
		return err
	}

	return nil
}
