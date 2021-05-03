package plan

import (
	"io"
)

type ExecuteOptions struct {
	Output io.Writer
}

type PreviewOptions struct {
	Output io.Writer
}
