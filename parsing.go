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

func (c *context) Param(p *Param) (v ParamValue, ok bool) {
	v, ok = c.params[p]
	return
}

const trimChars = " \t\r\n"

func (c *Cmd) ParseStr(line string, data interface{}) (Ctx Context, err error) {
	ctx := &context{
		data:   data,
		params: make(map[*Param]ParamValue),
	}
	line = strings.Trim(line, trimChars)

	err = c.parseStr(line, ctx)
	return ctx, err
}

func (c *Cmd) parseStr(line string, ctx *context) error {
	ctx.cmd = c
	for i, last, l, pLen := 0, 0, len(line), len(ctx.params); i <= l; i++ {
		if i < l && line[i] != ' ' {
			continue
		}
		part := strings.Trim(line[last:i], trimChars)
		if part == "" {
			continue
		}
		if pLen == 0 {
			// check sub command only if no param yet
			if subCmd, found := c.GetCommand(part); found {
				return subCmd.parseStr(strings.TrimLeft(line[i:], trimChars), ctx)
			}
		}
		// TODO : check flags before parameters
		if pLen < len(c.params) {
			param := c.params[pLen]
			ctx.params[param] = ParamValue{part}
			pLen++
		}
	}
	for _, p := range c.params {
		if _, ok := ctx.Param(p); p.options&paramRequired == paramRequired && !ok {
			return MissingParamError{p, c}
		}
	}
	return nil
}
