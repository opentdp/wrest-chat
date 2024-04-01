package args

import (
	"os"

	"github.com/opentdp/go-helper/filer"
	"gopkg.in/yaml.v3"
)

type Config struct {
	File string `yaml:"File"`
	Data any    `yaml:"Data"`
}

func (c *Config) Load() error {

	if !filer.Exists(c.File) {
		return nil
	}

	bytes, err := os.ReadFile(c.File)
	if err != nil {
		return err
	}

	return yaml.Unmarshal(bytes, c.Data)

}

func (c *Config) Save() error {

	bytes, err := yaml.Marshal(c.Data)
	if err != nil {
		return err
	}

	return os.WriteFile(c.File, bytes, 0644)

}
