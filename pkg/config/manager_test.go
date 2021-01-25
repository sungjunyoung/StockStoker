package config

import (
	"fmt"
	"testing"
)

type mockFileManager struct {
	content string
}

func (m *mockFileManager) Read() ([]byte, error) {
	return []byte(m.content), nil
}

func (m *mockFileManager) Write(data []byte) error {
	m.content = string(data)
	return nil
}

func TestManager_GetDataDir_ShouldSucceed(t *testing.T) {
	expect := "/unittest"
	mockContent := fmt.Sprintf(`dataDir: %s`, expect)
	manager := manager{fileManger: &mockFileManager{content: mockContent}}
	result, err := manager.GetDataDir()
	if err != nil {
		t.Fatalf("error should not be returned: %+v", err)
	}

	if result != expect {
		t.Fatalf("result %s is not equal with expected %s", result, expect)
	}
}

func TestManager_GetDataDir_ShouldFailedByUnmarshal(t *testing.T) {
	mockContent := `dataDir: a: 1`
	manager := manager{fileManger: &mockFileManager{content: mockContent}}
	_, err := manager.GetDataDir()
	if err == nil {
		t.Fatal("error should be returned")
	}
}

func TestManager_SetDataDir_ShouldSucceed(t *testing.T) {
	expect := "/unittest2"

	manager := manager{fileManger: &mockFileManager{}}
	if err := manager.SetDataDir("/unittest2"); err != nil {
		t.Fatalf("error should not be returned: %+v", err)
	}

	result, err := manager.GetDataDir()
	if err != nil {
		t.Fatal(err)
	}
	if result != expect {
		t.Fatalf("result %s is not equal with expected %s", result, expect)
	}
}
