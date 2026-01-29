package cli

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/dey12956/gator/internal/database"
	"github.com/google/uuid"
)

func HandlerFollow(s *State, cmd Command, user database.User) error {
	if len(cmd.Args) != 1 {
		return errors.New("you need to enter a feed url")
	}

	userID := user.ID

	feed, err := s.DB.GetFeedByUrl(context.Background(), cmd.Args[0])
	if err != nil {
		return err
	}
	feedID := feed.ID

	feedFollow, err := s.DB.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    userID,
		FeedID:    feedID,
	})
	if err != nil {
		return err
	}

	fmt.Printf("%s started following %s\n", feedFollow.UserName, feedFollow.FeedName)

	return nil
}
