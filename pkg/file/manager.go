package file

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

type Writer interface {
	Write(data []byte) error
}

type Reader interface {
	Read() ([]byte, error)
}

type Manager interface {
	Reader
	Writer
}

type manager struct {
	targetFile string
}

func NewManager(targetFile string) *manager {
	return &manager{
		targetFile: targetFile,
	}
}

func (m *manager) Read() ([]byte, error) {
	if !m.isPathExists() {
		return nil, fmt.Errorf("target file %s not exists", m.targetFile)
	}

	return ioutil.ReadFile(m.targetFile)
}

func (m *manager) Write(data []byte) error {
	dir := filepath.Dir(m.targetFile)
	if !m.isDirExists(dir) {
		if err := m.createDir(dir); err != nil {
			return err
		}
	}

	return ioutil.WriteFile(m.targetFile, data, 0644)
}

func (m *manager) isPathExists() bool {
	if _, err := os.Stat(m.targetFile); os.IsNotExist(err) {
		return false
	}
	return true
}

func (m *manager) isDirExists(dir string) bool {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		return false
	}
	return true
}

func (m *manager) createDir(dir string) error {
	return os.MkdirAll(dir, 0644)
}
