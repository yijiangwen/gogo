package util

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/mgutz/ansi"
)

var phosphorize = ansi.ColorFunc("gray+h")

type Command struct {
	dir    string
	env    []string
	status int
	cmd    *exec.Cmd
}

func NewCommand(path string, args ...string) (*Command, error) {
	var fullPath = ""
	fullPath, err := exec.LookPath(path)
	if err != nil {
		return nil, err
	}
	fmt.Println(phosphorize("exec:     " + fullPath + " " + strings.Join(args, " ")))
	return &Command{
		cmd: exec.Command(fullPath, args...),
		env: os.Environ(),
	}, nil
}

func (c *Command) SetEnv(name string, value string) {
	c.env = append(filterEnv(c.env, name), fmt.Sprintf("%s=%s", name, value))
}

func (c *Command) SetEnvLine(line string) {
	c.env = append(c.env, line)
}

func (c *Command) SetDir(dir string) {
	c.dir = dir
}

func (c *Command) Run() {
	c.cmd.Env = c.env
	c.cmd.Stdout = os.Stdout
	c.cmd.Stderr = os.Stderr
	c.cmd.Stdin = os.Stdin
	c.cmd.Dir = c.dir
	if err := c.cmd.Run(); err != nil {
		fmt.Println(phosphorize("exec: " + err.Error()))
	}
}

func filterEnv(env []string, removeName string) []string {
	removeName = removeName + "="
	ret := []string{}
	for _, v := range env {
		if strings.Index(v, removeName) != 0 {
			ret = append(ret, v)
		}
	}
	return ret
}
