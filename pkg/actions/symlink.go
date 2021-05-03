package actions

import (
	"fmt"
	"os"

	"github.com/sami2020pro/gmoji"

	"raw.tools/over/pkg/plan"
	"raw.tools/over/pkg/styles"
)

// Symlink step creates symlinks
type Symlink struct {
	Source string
	Target string
}

func (s *Symlink) String() string {
	prefix := styles.White.Sprint("link:")
	symbol := styles.White.Sprint("->")
	return fmt.Sprintf("%s %s %s %s %s", gmoji.Link, prefix, s.Source, symbol, s.Target)
}

func (s *Symlink) Execute(opts *plan.ExecuteOptions) error {
	return os.Symlink(s.Source, s.Target)
}

func EnsureLink(from string, to string) plan.Step {
	step := &Symlink{
		Source: from,
		Target: to,
	}
	stat, err := os.Lstat(to)
	if os.IsNotExist(err) {
		return step
	}
	if stat.Mode()&os.ModeSymlink == 0 {
		return &plan.Error{
			Reason: "Target exists and is not a symlink",
			Step:   step,
			Error:  nil,
		}
	}
	source, err := os.Readlink(to)
	if err != nil {
		return &plan.Error{
			Reason: "Unable to read existing link",
			Step:   step,
			Error:  err,
		}
	}
	if source != from {
		return &plan.Error{
			Reason: "Existing link source differ",
			Step:   step,
			Error:  nil,
		}
	}

	return &plan.Skip{
		Step:   step,
		Reason: "exists",
	}
}
