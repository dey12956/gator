package cli

import (
	"context"
	"errors"
	"fmt"
)

func HandlerFeeds(s *State, cmd Command) error {
	if len(cmd.Args) != 0 {
		return errors.New("you should not enter arguments")
	}

	feeds, err := s.DB.GetFeeds(context.Background())
	if err != nil {
		return err
	}
	for _, feed := range feeds {
		name := feed.Name
		url := feed.Url
		userName, err := s.DB.GetUserName(context.Background(), feed.UserID)
		if err != nil {
			return err
		}
		fmt.Printf("Feed Name: %s\n", name)
		fmt.Printf("Feed URL: %s\n", url)
		fmt.Printf("Created by: %s\n", userName)
		fmt.Println()
	}

	return nil
}
