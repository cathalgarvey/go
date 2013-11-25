package gist7576154

import (
	"io"
	"os/exec"
)

// CmdTemplater is the interface for command templates.
type CmdTemplater interface {
	NewCommand() *exec.Cmd
}

// CmdTemplate is a command template.
type CmdTemplate struct {
	NameArgs []string
	Dir      string
	Stdin    func() io.Reader
}

// NewCmdTemplate returns a CmdTemplate.
func NewCmdTemplate(name string, arg ...string) CmdTemplate {
	return CmdTemplate{
		NameArgs: append([]string{name}, arg...),
	}
}

// NewCommand generates a new *exec.Cmd from the template.
func (ct CmdTemplate) NewCommand() *exec.Cmd {
	cmd := exec.Command(ct.NameArgs[0], ct.NameArgs[1:]...)
	cmd.Dir = ct.Dir
	if ct.Stdin != nil {
		cmd.Stdin = ct.Stdin()
	}
	return cmd
}
