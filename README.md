# gator

gator is a small RSS feed aggregator CLI backed by Postgres.

## Prerequisites

- Go (see `go.mod` for the declared Go version)
- Postgres (to store users, feeds, follows, and posts)

## Install

Install the `gator` CLI with `go install`:

```bash
go install github.com/dey12956/gator@latest
```

Make sure your Go bin directory is on your `PATH`:

```bash
export PATH="$(go env GOPATH)/bin:$PATH"
```

## Database setup

1. Start Postgres.
2. Create a database for gator.
3. Apply the SQL migrations in `sql/schema`.

The migration files include `-- +goose` directives, but you can apply them with whatever migration tool you prefer.
If you want to use `goose`, one simple approach is:

```bash
go install github.com/pressly/goose/v3/cmd/goose@latest
goose -dir sql/schema postgres "$DB_URL" up
```

## Configuration

gator reads its config from `~/.gatorconfig.json`.

Create the file with at least a `db_url`:

```json
{
  "db_url": "postgres://postgres:postgres@localhost:5432/gator?sslmode=disable",
  "current_user_name": ""
}
```

- `db_url` is used to connect to Postgres.
- `current_user_name` is managed by `gator login` / `gator register`.

If `~/.gatorconfig.json` does not exist, the CLI will fail on startup.

## Run

After installing and configuring:

```bash
gator <command> [args]
```

For local development (from the repo root):

```bash
go run . <command> [args]
```

## Commands

gator is a simple command-based CLI (no flags).

- `register <username>`: create a user and set it as the current user
- `login <username>`: set the current user (must already exist)
- `users`: list all users (marks the current user)
- `reset`: delete all users (cascades to related data via foreign keys)

- `addfeed <name> <url>`: add a feed and automatically follow it (requires login)
- `feeds`: list all feeds
- `follow <url>`: follow an existing feed by URL (requires login)
- `unfollow <url>`: unfollow a feed by URL (requires login)
- `following`: list feeds the current user is following (requires login)

- `agg <time_between_reqs>`: continuously fetch feeds and store new posts
  - Example: `gator agg 1m`
- `browse [limit]`: show recent posts from feeds you follow (default limit: 2, requires login)
  - Example: `gator browse 10`
