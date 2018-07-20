package main // import "github.com/andrebq/validator-sample"

type (
	Project struct {
		Ops Ops
		Dev Dev
	}

	Ops struct {
		CommandPrefix string `yaml:"command_prefix",js:"command_prefix"`
		DepsCommand   string `yaml:"deps_command",js:"deps_command"`
	}

	Dev struct {
		// some stuff here
	}
)

func (p Project) Validate() Violation {
	return CombineValidation(p.Ops, p.Dev).Validate()
}

func (o Ops) Validate() Violation {
	return CombineValidation(
		MissingCommandPrefix(o.CommandPrefix),
		OnlyTwoCommandPrefix(o.CommandPrefix),
		MissingDepsCommand(o.DepsCommand)).Validate()
}

func (o Dev) Validate() Violation {
	return AlwaysValid().Validate() // dev is always too easy
}
