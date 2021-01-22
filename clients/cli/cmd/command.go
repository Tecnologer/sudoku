package cmd

import (
	"fmt"
	"strings"
)

type command struct {
	cmd    string
	action func(...string)
	about  string
	alias  []string
}

func (c *command) String() string {
	alias := ""
	if len(c.alias) > 0 {
		alias = "\n\t   * Alias: " + strings.Join(c.alias, ", ")
	}
	return fmt.Sprintf("- %s: %s%s", c.cmd, c.about, alias)
}
