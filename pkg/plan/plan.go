package plan

import (
	"fmt"
	"strings"

	"github.com/pkg/errors"
)

type Step interface {
	fmt.Stringer
	Execute(opts *ExecuteOptions) error
}

// Steps aggregate multiple steps
type Steps struct {
	steps []Step
}

func NewSteps() *Steps {
	return &Steps{
		steps: []Step{},
	}
}

func (s *Steps) Add(step Step) {
	if step != nil {
		s.steps = append(s.steps, step)
	}
}

func (s *Steps) Execute(opts *ExecuteOptions) error {
	for _, step := range s.steps {
		fmt.Println(step)
		if err := step.Execute(opts); err != nil {
			return errors.Wrapf(err, "Error while executing step %s", step)
		}
	}
	return nil
}

func (s *Steps) String() string {

	// It's ready to use from the get-go.
	// You don't need to initialize it.
	var sb strings.Builder

	for _, step := range s.steps {
		sb.WriteString(step.String())
	}
	return sb.String()
}

// Plan stores the succession of operations to reach a desired state.
type Plan struct {
	Steps *Steps
}

func New() *Plan {
	return &Plan{
		Steps: NewSteps(),
	}
}

func (p *Plan) Add(step Step) {
	p.Steps.Add(step)
}

func (p *Plan) Execute(opts *ExecuteOptions) error {
	return p.Steps.Execute(opts)
}

func (p *Plan) Preview(opts *PreviewOptions) error {
	// return p.Steps.Preview(opts)
	fmt.Println(p.Steps.String())
	return nil
}
