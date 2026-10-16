package config

import (
	"fmt"
	"os/exec"
)

type app struct {
	info
	BundleID string // prefered
	Path     string
	Args     []string
}

func (a *app) Run() error {
	var cmd *exec.Cmd
	switch {
	case a.BundleID != "":
		rest := make([]string, 0, len(a.Args)+3)
		rest = append(rest, "-b", a.BundleID)
		if len(a.Args) != 0 {
			rest = append(rest, "--args")
			rest = append(rest, a.Args...)
		}

		cmd = exec.Command("open", rest...)
	case a.Path != "":
		rest := make([]string, 0, len(a.Args)+2)
		rest = append(rest, a.Path)
		if len(a.Args) != 0 {
			rest = append(rest, "--args")
			rest = append(rest, a.Args...)
		}

		cmd = exec.Command("open", rest...)
	default:
		return fmt.Errorf("provided app has no path/bundleid")
	}

	return cmd.Run()
}

func (a *app) Info() info {
	return a.info
}
