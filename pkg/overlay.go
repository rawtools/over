package pkg

import (
	"fmt"
	"io/fs"
	"path/filepath"

	"github.com/milad-abbasi/gonfig"
	"github.com/pkg/errors"

	"raw.tools/over/pkg/actions"
	"raw.tools/over/pkg/plan"
	"raw.tools/over/pkg/styles"
)

type OverlayConfig struct {
	Name    string
	Exclude []string
}

type Overlay struct {
	Name   string
	Root   string
	Config *OverlayConfig
}

// NewOverlay instianciate a new or existing overlay
func NewOverlay(name string, root string) *Overlay {
	return &Overlay{
		Name:   name,
		Root:   root,
		Config: &OverlayConfig{},
	}
}

func (o *Overlay) String() string {
	return o.Name
}

func (o *Overlay) Pretty() string {
	return styles.White.Sprint(o.Name)
}

func (o *Overlay) ParseConfig() error {
	path := filepath.Join(o.Root, "over.yaml")

	// Input argument must be a pointer to struct
	err := gonfig.Load().FromFile(path).Into(o.Config)
	if err != nil {
		return fmt.Errorf("error loading config from %s: \n%w", path, err)
	}
	return nil
}

// Init perform the overlay initialization.
// If the overlay does not exists, Init create the overlay on-disk layout
// and adds initial files to git management.
// If the overlay already exists, Init does nothing.
func (o *Overlay) Init(path string) error {
	return errors.New("Not implemented")
}

// Delete the overlay from the repository.
// If the overlay does not exists, Delete does not nothing.
// If the overlay exists, files are removed from git management and disk.
func (o *Overlay) Delete(path string) error {
	return errors.New("Not implemented")
}

// Add a file to the overlay
func (o *Overlay) Add(path string) error {
	return errors.New("Not implemented")
}

func (o *Overlay) PlanExecution(target string, opts *plan.ExecuteOptions) (*plan.Plan, error) {

	p := plan.New()

	err := filepath.Walk(o.Root, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return fmt.Errorf("prevent panic by handling failure accessing a path %q:\n%w", path, err)
		}
		// if info.IsDir() && info.Name() == subDirToSkip {
		// 	fmt.Printf("skipping a dir without errors: %+v \n", info.Name())
		// 	return filepath.SkipDir
		// }
		// fmt.Printf("visited file or dir: %q (%s)\n", path, info)
		// fmt.Printf("visited file or dir: %q\n", path)

		base, err := filepath.Rel(o.Root, path)
		if err != nil {
			return err
		}
		if base == "over.yaml" {
			return nil
		}
		targetLn := filepath.Join(target, base)

		if info.IsDir() {
			p.Add(actions.EnsureDirectory(targetLn, info.Mode()))
		} else {
			p.Add(actions.EnsureLink(path, targetLn))
		}
		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("error walking the path %q: %w", o.Root, err)
	}
	return p, nil
}

// Rm remove a file from the overlay
func (o *Overlay) Rm(path string) error {
	return errors.New("Not implemented")
}

func (o *Overlay) Apply(target string) error {
	plan, err := o.PlanExecution(target, nil)
	if err != nil {
		return fmt.Errorf("unable to plan execution to %q: %w", target, err)
	}

	return plan.Execute(nil)
}
