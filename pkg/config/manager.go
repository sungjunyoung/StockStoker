package config

import (
	"github.com/sungjunyoung/StockStoker/pkg/file"
	"gopkg.in/yaml.v2"
)

type Manager interface {
	GetDataDir() (string, error)
	SetDataDir(dataDir string) error
}

type manager struct {
	fileManger file.Manager
}

func NewManager(configFile string) *manager {
	return &manager{
		fileManger: file.NewManager(configFile),
	}
}

func (m *manager) GetDataDir() (string, error) {
	conf, err := m.readConfig()
	if err != nil {
		return "", err
	}

	return conf.DataDir, nil
}

func (m *manager) SetDataDir(dataDir string) error {
	conf, err := m.readConfig()
	if err != nil {
		return err
	}

	conf.DataDir = dataDir
	return m.writeConfig(conf)
}

func (m *manager) readConfig() (*Config, error) {
	data, err := m.fileManger.Read()
	if err != nil {
		return nil, err
	}

	conf := Config{}
	if err := yaml.Unmarshal(data, &conf); err != nil {
		return nil, err
	}

	return &conf, nil
}

func (m *manager) writeConfig(conf *Config) error {
	data, err := yaml.Marshal(conf)
	if err != nil {
		return err
	}

	return m.fileManger.Write(data)
}
