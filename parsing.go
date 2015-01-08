package commander

import (
	"fmt"
	"github.com/JeanSebTr/recoverer"
	"strings"
)

var _ = fmt.Sprintf("")

type context struct {
	cmd    *Cmd
	data   interface{}
	params map[*Param]ParamValue
}

func (c *context) Command() *Cmd {
	return c.cmd
}

func (c *context) Data() interface{} {
	return c.data
}

func (c *context) Run() (err error) {
	defer recoverer.Catch(&err)
	c.cmd.Action(c)
	return err
}

func (c *context) Param(p *Param) ParamValue {
	return c.params[p]
}

const trimChars = " \t\r\n"

func (c *Cmd) ParseStr(line string, data interface{}) (Context, error) {
	ctx := &context{
		data:   data,
		params: make(map[*Param]ParamValue),
	}
	line = strings.Trim(line, trimChars)

	if err := c.parseStr(line, ctx); err != nil {
		return nil, err
	}
	return ctx, nil
}

func (c *Cmd) parseStr(line string, ctx *context) error {
	ctx.cmd = c
	l := len(line)
	if l == 0 {
		return nil
	}
	for i, last, l, pLen := 0, 0, len(line), len(ctx.params); i <= l; i++ {
		if i < l && line[i] != ' ' {
			continue
		}
		part := line[last:i]
		if subCmd, found := c.GetCommand(part); found {
			fmt.Printf("Cmd: %s\n", subCmd.Name)
			return subCmd.parseStr(strings.TrimLeft(line[i:], trimChars), ctx)
		}
		// TODO : check flags before parameters
		if pLen < len(c.params) {
			param := c.params[pLen]
			ctx.params[param] = ParamValue{part}
			pLen++
		}
	}
	return nil
}
