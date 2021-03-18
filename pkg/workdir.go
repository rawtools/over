package pkg

// Workdir represent a git working directory used by git-over
type Workdir struct {
	Root    string
	Overlay *Overlay
}
