package actions

import (
	"fmt"
	"io/fs"
	"os"

	"github.com/sami2020pro/gmoji"

	"raw.tools/over/pkg/plan"
	"raw.tools/over/pkg/styles"
)

type MkDir struct {
	Path string
	Mode fs.FileMode
}

func (m *MkDir) String() string {
	prefix := styles.White.Sprint("create directory:")
	return fmt.Sprintf("%s %s %s (%s)", gmoji.Folder, prefix, m.Path, m.Mode)
}

func (m *MkDir) Execute(opts *plan.ExecuteOptions) error {
	return os.MkdirAll(m.Path, m.Mode)
}

func EnsureDirectory(path string, mode fs.FileMode) plan.Step {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return &MkDir{
			Path: path,
			Mode: mode,
		}
		// p.Add()
	}
	return nil
}
