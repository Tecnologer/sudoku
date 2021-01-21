package cmd

import "fmt"

type command struct {
	cmd    string
	action func()
	about  string
	alias  []string
}

func (c *command) String() string {
	return fmt.Sprintf("- %s: %s\n\t* Alias: %v", c.cmd, c.about, c.alias)
}
