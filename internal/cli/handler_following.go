package cli

import (
	"context"
	"errors"
	"fmt"

	"github.com/dey12956/gator/internal/database"
)

func HandlerFollowing(s *State, cmd Command, user database.User) error {
	if len(cmd.Args) != 0 {
		return errors.New("you should not enter any argument")
	}

	userID := user.ID

	feeds, err := s.DB.GetFeedFollowsForUser(context.Background(), userID)
	if err != nil {
		return err
	}

	fmt.Printf("%s is following:\n", user.Name)
	for _, feed := range feeds {
		fmt.Printf("- %s\n", feed.FeedName)
	}

	return nil
}
