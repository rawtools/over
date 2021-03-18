package plan

import (
	"fmt"

	"raw.tools/over/pkg/styles"
)

type Skip struct {
	Step Step
}

func (s *Skip) String() string {
	prefix := styles.Cyan.Sprint("skip:")

	return fmt.Sprintf("%s %s", prefix, s.Step)
}

func (s *Skip) Execute(opts *ExecuteOptions) error {
	return nil
}
