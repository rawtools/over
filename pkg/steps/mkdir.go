package steps

import (
	"fmt"
	"io/fs"
	"os"

	"raw.tools/over/pkg/plan"
	"raw.tools/over/pkg/styles"
)

type MkDir struct {
	Path string
	Mode fs.FileMode
}

func (m *MkDir) String() string {
	prefix := styles.White.Sprint("create directory:")
	return fmt.Sprintf("📂 %s %s (%s)", prefix, m.Path, m.Mode)
}

func (m *MkDir) Execute(opts *plan.ExecuteOptions) error {
	return os.MkdirAll(m.Path, m.Mode)
}
