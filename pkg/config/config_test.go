package config

import (
	"fmt"
	"github.com/ghodss/yaml"
	"io/ioutil"
	"math/rand"
	"os"
	"testing"
)

func TestConfig_Load(t *testing.T) {
	data := `markets:
- kospi
- nasdaq`
	path := fmt.Sprintf("/tmp/%d", rand.Intn(10000))
	if err := ioutil.WriteFile(path, []byte(data), 0644); err != nil {
		t.Fatalf("setup failed: %+v", err)
	}

	c := New(path)
	if err := c.Load(); err != nil {
		t.Fatalf("load must not be failed: %+v", err)
	}
}

func TestConfig_LoadNotExists(t *testing.T) {
	path := fmt.Sprintf("/tmp/%d", rand.Intn(10000))
	c := New(path)
	if err := c.Load(); err == nil {
		t.Fatal("load must be failed")
	}
}

func TestConfig_LoadUnmarshalFail(t *testing.T) {
	data := `unittest1`
	path := fmt.Sprintf("/tmp/%d", rand.Intn(10000))
	if err := ioutil.WriteFile(path, []byte(data), 0644); err != nil {
		t.Fatalf("setup failed: %+v", err)
	}
	defer os.Remove(path)

	c := New(path)
	if err := c.Load(); err == nil {
		t.Fatal("load must be failed")
	}
}

func TestConfig_Save(t *testing.T) {
	path := fmt.Sprintf("/tmp/%d", rand.Intn(10000))
	defer os.Remove(path)

	c := New(path)
	if err := c.Save(); err != nil {
		t.Fatal("save must not be failed")
	}

	data, err := ioutil.ReadFile(path)
	if err != nil {
		t.Fatal("file should be read")
	}
	config := &config{}
	if err := yaml.Unmarshal(data, config); err != nil {
		t.Fatal("loaded data should be unmarshal")
	}
}
