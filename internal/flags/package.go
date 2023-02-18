package flags

import (
	"fmt"
)

type Package string

func (f *Package) String() string {
	return string(*f)
}

func (f *Package) Set(str string) error {
	if str == "" {
		str = f.Default()
	}
	*f = Package(str)
	return nil
}

func (f *Package) Type() string {
	return "package"
}

func (f *Package) Default() string {
	return "main"
}

func (f *Package) Shorthand() string {
	return "p"
}

func (f *Package) Description() string {
	return fmt.Sprintf("The name of the module's package used for building the templates. (default \"%s\")", f.Default())
}