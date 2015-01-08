package commander

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
	Param(*Param) ParamValue
}

func (c *Cmd) AddSubCommand(sub *Cmd) {
	c.subCommands = append(c.subCommands, sub)
}

func (c *Cmd) AddFlag(short, long string) {
	//
}

func (c *Cmd) Param(param *Param) *Cmd {
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
