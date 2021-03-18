package plan

import (
	"fmt"

	"github.com/pkg/errors"
)

type Step interface {
	fmt.Stringer
	Execute(opts *ExecuteOptions) error
}

// Plan stores the succession of operations to reach a desired state.
type Plan struct {
	Steps []Step
}

func New() *Plan {
	return &Plan{
		Steps: []Step{},
	}
}

func (p *Plan) Add(step Step) {
	p.Steps = append(p.Steps, step)
}

func (p *Plan) Execute(opts *ExecuteOptions) error {
	for _, step := range p.Steps {
		fmt.Println(step)
		if err := step.Execute(opts); err != nil {
			return errors.Wrapf(err, "Error while executing step %s", step)
		}
	}
	return nil
}

func (p *Plan) Preview(opts *PreviewOptions) error {
	for _, step := range p.Steps {
		fmt.Println(step)
	}
	return nil
}
