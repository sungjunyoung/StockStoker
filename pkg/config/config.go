package config

import (
	"github.com/ghodss/yaml"
	"github.com/sungjunyoung/StockStoker/pkg/market"
	"io/ioutil"
)

type Configer interface {
	Load() error
	Save(markets []market.Market) error
}

type config struct {
}

type Config struct {
	path string
	*config
}

func New(path string) *Config {
	return &Config{path: path}
}

func (c *Config) Load() error {
	file, err := ioutil.ReadFile(c.path)
	if err != nil {
		return err
	}

	config := &config{}
	if err = yaml.Unmarshal(file, config); err != nil {
		return err
	}

	c.config = config
	return nil
}

func (c *Config) Save() error {
	config := &config{}

	data, err := yaml.Marshal(config)
	if err != nil {
		return err
	}

	if err := ioutil.WriteFile(c.path, data, 0644); err != nil {
		return err
	}

	c.config = config
	return nil
}
