package cli

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/dey12956/gator/internal/database"
	"github.com/google/uuid"
)

func HandlerAddFeed(s *State, cmd Command, user database.User) error {
	if len(cmd.Args) != 2 {
		return errors.New("you need to enter the name and the url of a new feed")
	}

	userID := user.ID

	_, err := s.DB.CreateFeed(context.Background(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      cmd.Args[0],
		Url:       cmd.Args[1],
		UserID:    userID,
	})
	if err != nil {
		return err
	}

	fmt.Printf("Feed %s is added by %s\n", cmd.Args[0], s.C.CurrentUserName)

	feed, err := s.DB.GetFeedByUrl(context.Background(), cmd.Args[1])
	if err != nil {
		return err
	}
	feedID := feed.ID

	_, err = s.DB.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    userID,
		FeedID:    feedID,
	})
	if err != nil {
		return err
	}

	fmt.Printf("%s started following %s\n", user.Name, feed.Name)

	return nil
}
