package commander

import (
	"github.com/JeanSebTr/recoverer"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSimpleParsing(t *testing.T) {
	app := &Cmd{}

	cmd1 := &Cmd{
		Name: "cmd1",
	}

	app.AddSubCommand(cmd1)

	for i, line := range []string{"cmd1", "cmd1 ", "\ncmd1  \t ", "cmd1 arg1 \targ2"} {
		ctx, err := app.ParseStr(line, i)
		assert.NoError(t, err, "Index %d", ctx.Data().(int))
		assert.Equal(t, cmd1, ctx.Command(), "Index %d", ctx.Data().(int))
	}
}

func TestAddingRequiredParamAfterOptionalShouldPanic(t *testing.T) {
	var err error
	defer recoverer.Catch(&err)

	cmd1 := &Cmd{
		Name: "cmd1",
	}
	cmd1.Param(&Param{
		Name: "optional",
	}).Param((&Param{
		Name: "required",
	}).Required())

	t.Error("Test did not panic")
}

func TestRequiredParamMissing(t *testing.T) {
	app := &Cmd{}

	cmd1 := &Cmd{
		Name: "cmd1",
	}

	arg1 := (&Param{}).Required()
	cmd1.Param(arg1)

	app.AddSubCommand(cmd1)

	cases := []string{"cmd1", "cmd1  \t"}
	for i, line := range cases {
		ctx, err := app.ParseStr(line, i)
		assert.Error(t, err, "Index %d", ctx.Data().(int))
	}
}
