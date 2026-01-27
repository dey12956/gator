package cli

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/dey12956/gator/internal/database"
	"github.com/google/uuid"
)

func HandlerRegister(s *State, cmd Command) error {
	if len(cmd.Args) == 0 {
		return errors.New("you need to enter the username argument")
	}

	_, err := s.DB.CreateUser(context.Background(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      cmd.Args[0],
	})
	if err != nil {
		return err
	}

	err = s.C.SetUser(cmd.Args[0])
	if err != nil {
		return err
	}

	fmt.Printf("User %s is created\n", cmd.Args[0])

	return nil
}
