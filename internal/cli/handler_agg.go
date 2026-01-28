package cli

import (
	"context"
	"fmt"

	"github.com/dey12956/gator/internal/rss"
)

func HandlerAgg(s *State, cmd Command) error {
	feed, err := rss.FetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")
	if err != nil {
		return err
	}

	fmt.Print(*feed)
	return nil
}
