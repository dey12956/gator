package cli

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/dey12956/gator/internal/database"
	"github.com/dey12956/gator/internal/rss"
	"github.com/google/uuid"
	"github.com/lib/pq"
)

func HandlerAgg(s *State, cmd Command) error {
	if len(cmd.Args) != 1 {
		return errors.New("incorrect time_between_reqs")
	}

	timeBetweenReqs, err := time.ParseDuration(cmd.Args[0])
	if err != nil {
		return err
	}

	fmt.Printf("Collecting feeds every %s\n", cmd.Args[0])

	ticker := time.NewTicker(timeBetweenReqs)
	for ; ; <-ticker.C {
		scrapeFeeds(s)
	}
}

func scrapeFeeds(s *State) {
	feed, err := s.DB.GetNextFeedToFetch(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	err = s.DB.MarkFeedFetched(context.Background(), feed.ID)
	if err != nil {
		log.Fatal(err)
	}

	rssFeed, err := rss.FetchFeed(context.Background(), feed.Url)
	if err != nil {
		log.Fatal(err)
	}

	for _, post := range rssFeed.Channel.Item {
		description := post.Description
		if strings.TrimSpace(post.Content) != "" {
			description = post.Content
		}

		pubDate := parsePostPubDate(post.PubDate)
		_, err := s.DB.CreatePost(context.Background(), database.CreatePostParams{
			ID:          uuid.New(),
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			Title:       post.Title,
			Url:         post.Link,
			Description: description,
			PublishedAt: pubDate,
			FeedID:      feed.ID,
		})
		if err != nil {
			var pqErr *pq.Error
			if errors.As(err, &pqErr) && string(pqErr.Code) != "23505" {
				log.Fatal(err)
			}
		}
	}
}

func parsePostPubDate(raw string) time.Time {
	s := strings.TrimSpace(raw)

	layouts := []string{
		time.RFC3339Nano,               // 2006-01-02T15:04:05.999999999Z07:00
		time.RFC3339,                   // 2006-01-02T15:04:05Z07:00
		"2006-01-02T15:04:05Z0700",     // 2006-01-02T15:04:05-0700 (no colon)
		"2006-01-02T15:04:05.000Z0700", // with millis + no colon tz
		time.RFC1123Z,                  // Mon, 02 Jan 2006 15:04:05 -0700
		time.RFC1123,                   // Mon, 02 Jan 2006 15:04:05 MST
		time.RFC850,                    // Monday, 02-Jan-06 15:04:05 MST
		time.RFC822Z,                   // 02 Jan 06 15:04 -0700
		time.RFC822,                    // 02 Jan 06 15:04 MST
		time.RubyDate,                  // Mon Jan 02 15:04:05 -0700 2006
		time.UnixDate,                  // Mon Jan 02 15:04:05 MST 2006
		time.ANSIC,                     // Mon Jan _2 15:04:05 2006
		"2006-01-02 15:04:05",          // common “no tz” format
		"2006-01-02",
	}

	for _, layout := range layouts {
		if t, err := time.Parse(layout, s); err == nil {
			return t
		}
	}

	return time.Time{}
}
