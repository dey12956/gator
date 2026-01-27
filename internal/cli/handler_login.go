package cli

import (
	"context"
	"errors"
	"fmt"
)

func HandlerLogin(s *State, cmd Command) error {
	if len(cmd.Args) == 0 {
		return errors.New("you need to enter the username argument")
	}

	_, err := s.DB.GetUser(context.Background(), cmd.Args[0])
	if err != nil {
		return err
	}

	err = s.C.SetUser(cmd.Args[0])
	if err != nil {
		return err
	}

	fmt.Printf("Current username is set to %s\n", cmd.Args[0])

	return nil
}
