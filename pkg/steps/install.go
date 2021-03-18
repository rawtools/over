package steps

import (
	"os/exec"
)

type DependencyList struct {
	system string
	// dependencies []string
}

func (dp *DependencyList) Install() error {
	var cmd *exec.Cmd
	switch dp.system {
	case "arch":
		cmd = dp.installOnArch()
	case "debian":
		cmd = dp.installOnDebian()
	}

	return cmd.Run()
}

func (dp *DependencyList) installOnArch() *exec.Cmd {
	return nil
}

func (dp *DependencyList) installOnDebian() *exec.Cmd {
	return nil
}
