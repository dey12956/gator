package cli

type Commands struct {
	MapOfCommands map[string]func(*State, Command) error
}

func (c *Commands) Run(s *State, cmd Command) error {
	err := c.MapOfCommands[cmd.Name](s, cmd)
	if err != nil {
		return err
	}

	return nil
}

func (c *Commands) Register(name string, f func(*State, Command) error) {
	c.MapOfCommands[name] = f
}
