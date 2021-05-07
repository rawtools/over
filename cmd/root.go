package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/adrg/xdg"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	over "raw.tools/over/pkg"
)

func InitConfig() {

}

var (
	parser = over.NewParser()
	config over.Config
)

func RootCommand(name string) *cobra.Command {
	cmd := &cobra.Command{
		Use:     name,
		Short:   "git-based overlays",
		Long:    `over allows you to version your configuration files and workspaces settings`,
		Version: over.Version,
	}

	// Adds common sub commands
	cmd.AddCommand(AddCmd)
	cmd.AddCommand(ReleaseCmd)
	cmd.AddCommand(ApplyCmd)
	cmd.AddCommand(StatusCmd)

	// parser := over.NewParser()

	pflags := cmd.PersistentFlags()

	// add config flag/env early parsing
	defaultHome := filepath.Join(xdg.ConfigHome, "over")
	if homeEnv, ok := os.LookupEnv(fmt.Sprintf("%sHOME", over.EnvPrefix)); ok {
		defaultHome = homeEnv
	}
	pflags.StringVarP(&config.Home, "home", "H", defaultHome, "Path to config file")

	pflags.BoolP("debug", "D", false, "Toggle debugging")
	if err := parser.BindPFlag("debug", pflags.Lookup("debug")); err != nil {
		fmt.Printf("unable to create flag: %s\n", err)
	}

	return cmd
}

// NewParser instanciate a new Viper config parser configured for over
func Parse() (*over.Config, error) {
	parser.AddConfigPath(config.Home)
	parser.AutomaticEnv() // Do the darn thing :D

	if err := parser.ReadInConfig(); err == nil {
		fmt.Printf("loading config from: %s\n", parser.ConfigFileUsed())
	} else if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
		// Config file not found; ignore error if desired
		fmt.Printf("unable to load config: %s\n", err)
	}

	if err := parser.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("config unmarshalling failed: %w", err)
	}

	return &config, nil
}
