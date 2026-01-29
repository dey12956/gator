package cli

import (
	"context"
	"errors"
	"fmt"
	"strconv"
)

func HandlerBrowse(s *State, cmd Command) error {
	limit := 2
	if len(cmd.Args) == 1 {
		var err error
		limit, err = strconv.Atoi(cmd.Args[0])
		if err != nil {
			return err
		}
	} else if len(cmd.Args) > 1 {
		return errors.New("you should only enter a limit number")
	}

	posts, err := s.DB.GetPostsForUser(context.Background(), int32(limit))
	if err != nil {
		return err
	}
	for _, post := range posts {
		fmt.Print(post.Title)
		fmt.Println()
		fmt.Print(post.Description)
		fmt.Println()
	}

	return nil
}
