package cli

import (
	"context"
	"errors"
	"fmt"

	"github.com/dey12956/gator/internal/database"
)

func HandlerUnfollow(s *State, cmd Command, user database.User) error {
	if len(cmd.Args) != 1 {
		return errors.New("you need to enter a feed url")
	}

	userID := user.ID

	feed, err := s.DB.GetFeedByUrl(context.Background(), cmd.Args[0])
	if err != nil {
		return err
	}
	feedID := feed.ID

	err = s.DB.DeleteFeedFollowByUserIDAndFeedID(context.Background(), database.DeleteFeedFollowByUserIDAndFeedIDParams{
		UserID: userID,
		FeedID: feedID,
	})
	if err != nil {
		return err
	}

	fmt.Printf("%s unfollowed %s\n", user.Name, feed.Name)

	return nil
}
