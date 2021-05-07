package pkg

import (
	"fmt"

	"github.com/spf13/viper"
)

const (
	// defaultCfg = "~/.config/over.yml"
	EnvPrefix = "OVER_"
)

// Config stores the un marshalled parameters
type Config struct {
	Home  string
	Debug bool
}

func (cfg Config) String() string {
	return fmt.Sprintf("Home: %s, DEBUG: %t", cfg.Home, cfg.Debug)
}

// type Parser *viper.Viper

// NewParser instanciate a new Viper config parser configured for over
func NewParser() *viper.Viper {
	parser := viper.New()
	parser.SetConfigName("over")
	parser.SetEnvPrefix(EnvPrefix) // All environment variables that match this prefix will be loaded and can be referenced in the code
	parser.AllowEmptyEnv(true)     // Allows an environment variable to not exist and not blow up, I suggest using switch statements to handle these though
	parser.AutomaticEnv()          // Do the darn thing :D

	return parser
}

// NewConfig extracts configuration parameters from execution parameters
// and configuration file
func NewConfig(parser *viper.Viper) (*Config, error) {
	var config Config

	if err := parser.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("unable to read config: %w", err)
	}

	if err := parser.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("config unmarshalling failed: %w", err)
	}

	return &config, nil
}
