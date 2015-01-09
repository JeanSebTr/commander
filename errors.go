package commander

import (
	"fmt"
)

type MissingParamError struct {
	*Param
	*Cmd
}

func (e MissingParamError) Error() string {
	return fmt.Sprintf("Missing param %s from command %s", e.Param.Name, e.Cmd.Name)
}
