// Package cli is for handling CLI commands
package cli

import (
	"github.com/dey12956/gator/internal/config"
	"github.com/dey12956/gator/internal/database"
)

type State struct {
	DB *database.Queries
	C  *config.Config
}
