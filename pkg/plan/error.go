package plan

import (
	"fmt"

	"raw.tools/over/pkg/styles"
)

// Error wraps an errored step and optionaly its error.
type Error struct {
	Step   Step
	Reason string
	Error  error
}

func (e *Error) String() string {
	prefix := styles.Red.Sprintf("error(%s):", e.Reason)

	return fmt.Sprintf("%s %s", prefix, e.Step)
}

func (e *Error) Execute(opts *ExecuteOptions) error {
	return nil
}
