package config

import "os/exec"

type command struct {
	info
	Command string
	Args    []string
}

func (c *command) Run() error {
	return exec.Command(c.Command, c.Args...).Run()
}

func (c *command) Info() info {
	return c.info
}
