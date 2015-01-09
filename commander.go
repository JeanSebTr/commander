package commander

import (
	"fmt"
)

type Cmd struct {
	Name string

	Action func(Context)
	//Completion  Completable
	subCommands []*Cmd
	params      []*Param
	// flags
}

type Context interface {
	Command() *Cmd
	Data() interface{}
	Run() error
	Param(*Param) (ParamValue, bool)
}

func (c *Cmd) AddSubCommand(sub *Cmd) {
	c.subCommands = append(c.subCommands, sub)
}

func (c *Cmd) AddFlag(short, long string) {
	//
}

func (c *Cmd) Param(param *Param) *Cmd {
	i := len(c.params) - 1
	if i >= 0 && param.options&paramRequired == paramRequired {
		p := c.params[i]
		if p.options&paramRequired == 0 {
			panic(fmt.Errorf("Cannot add required param %s after optional param %s on command %s", param.Name, p.Name, c.Name))
		}
	}
	c.params = append(c.params, param)
	return c
}

func (c *Cmd) GetCommand(name string) (*Cmd, bool) {
	for _, cmd := range c.subCommands {
		if cmd.Name == name {
			return cmd, true
		}
	}
	return nil, false
}
