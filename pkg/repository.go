package pkg

import (
	"errors"
	"fmt"
	"path/filepath"

	"github.com/bmatcuk/doublestar/v2"
)

type Repository struct {
	Cfg      *Config
	Root     string
	Overlays map[string]*Overlay
}

func NewRepository(cfg *Config) *Repository {
	return &Repository{
		Cfg:  cfg,
		Root: cfg.Home,
	}
}

func (r *Repository) String() string {
	return fmt.Sprintf("Repository %s", r.Root)
}

// Init create the repository on disk if not existant
func (r *Repository) Init() error {
	return errors.New("Not implemented")
}

// List all overlays into the repository
func (r *Repository) List() []*Overlay {
	if r.Overlays == nil {
		if err := r.parseOverlays(); err != nil {
			return nil
		}
	}
	overlays := []*Overlay{}
	for _, name := range SortedStringKeys(r.Overlays) {
		overlays = append(overlays, r.Overlays[name])
	}
	// for _, overlay := range r.Overlays {
	// 	overlays = append(overlays, overlay)
	// }
	return overlays
}

func (r *Repository) Get(overlay string) *Overlay {
	if r.Overlays == nil {
		if err := r.parseOverlays(); err != nil {
			return nil
		}
	}
	return r.Overlays[overlay]
}

// New create a new overlay into the repository
func (r *Repository) New() *Overlay {
	return nil
}

// Rm removes an overlay from the repository
func (r *Repository) Rm() error {
	return errors.New("Not implemented")
}

func (r *Repository) parseOverlays() error {
	matches, err := doublestar.Glob(filepath.Join(r.Root, "**/over.yaml"))
	if err != nil {
		return fmt.Errorf("unable to list overlays: %w", err)
	}
	r.Overlays = map[string]*Overlay{}
	for _, v := range matches {
		path := filepath.Dir(v)
		name, err := filepath.Rel(r.Root, path)
		if err != nil {
			return fmt.Errorf("unable to compute overlay name: %w", err)
		}
		r.Overlays[name] = NewOverlay(name, path)
	}
	return nil
}
