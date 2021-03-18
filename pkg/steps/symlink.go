package steps

import (
	"fmt"
	"os"

	"raw.tools/over/pkg/plan"
	"raw.tools/over/pkg/styles"
)

type Symlink struct {
	Source string
	Target string
}

func (s *Symlink) String() string {
	prefix := styles.White.Sprint("link:")
	symbol := styles.White.Sprint("->")
	return fmt.Sprintf("🔗 %s %s %s %s", prefix, s.Source, symbol, s.Target)
}

func (s *Symlink) Execute(opts *plan.ExecuteOptions) error {
	return os.Symlink(s.Source, s.Target)
}
