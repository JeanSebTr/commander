package commander

import (
	"fmt"
)

type Param struct {
	Name    string
	options paramOptions
	//Completion Completable
}

type ParamValue struct {
	value interface{}
}

func (p ParamValue) String() string {
	if str, ok := p.value.(string); ok {
		return str
	}
	return fmt.Sprint(p.value)
}

type paramOptions uint8

const (
	paramRequired paramOptions = 1 << iota
	paramEllipsis
)

func (p *Param) Required() *Param {
	p.options |= paramRequired
	return p
}

func (p *Param) MultiWords() *Param {
	p.options |= paramEllipsis
	return p
}
