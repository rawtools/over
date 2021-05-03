package plan

import (
	"fmt"

	"raw.tools/over/pkg/styles"
)

// Skip ignore the wrapped step for a given reason (without error)
type Skip struct {
	Step   Step
	Reason string
}

func (s *Skip) String() string {
	prefix := styles.Cyan.Sprintf("skip(%s):", s.Reason)

	return fmt.Sprintf("%s %s", prefix, s.Step)
}

func (s *Skip) Execute(opts *ExecuteOptions) error {
	return nil
}
