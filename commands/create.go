package commands

import (
	"context"
	"flag"
	"fmt"

	"github.com/google/subcommands"
)

// CreateCommand creates a command.
func CreateCommand() subcommands.Command {
	return new(createCmd)
}

type createCmd struct {
}

func (*createCmd) Name() string     { return "create" }
func (*createCmd) Synopsis() string { return "Creates a new project." }
func (*createCmd) Usage() string {
	return `create <something>: 
	Creates a project
`
}

func (p *createCmd) SetFlags(f *flag.FlagSet) {
	//  f.BoolVar(&p.capitalize, "capitalize", false, "capitalize output")
}

func (*createCmd) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	fmt.Println("I would create something now")
	return subcommands.ExitSuccess
}
