package file

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestFileManager_Read_ShouldSucceed(t *testing.T) {
	expect := []byte("unittest")
	targetFile, err := ioutil.TempFile(os.TempDir(), "unittest-")
	if err != nil {
		t.Fatal(err)
	}
	if _, err := targetFile.Write(expect); err != nil {
		t.Fatal(err)
	}
	defer targetFile.Close()

	fileManager := NewManager(targetFile.Name())
	result, err := fileManager.Read()
	if err != nil {
		t.Fatalf("error should not be returned: %+v", err)
	}

	if string(result) != string(expect) {
		t.Fatalf("result %+v is not equal with expected %+v", result, expect)
	}
}

func TestFileManager_Read_ShouldFailedByFileNotExists(t *testing.T) {
	fileManager := NewManager("/not_exists_file")
	_, err := fileManager.Read()
	if err == nil {
		t.Fatalf("error should be returend")
	}
}

func TestFileManager_Write_ShouldSucceed(t *testing.T) {
	expect := []byte("unittest")

	targetFile, err := ioutil.TempFile(os.TempDir(), "unittest-")
	if err != nil {
		t.Fatal(err)
	}
	defer targetFile.Close()

	fileManager := NewManager(targetFile.Name())
	if err := fileManager.Write(expect); err != nil {
		t.Fatalf("error should not be returned: %+v", err)
	}

	result, err := ioutil.ReadFile(targetFile.Name())
	if err != nil {
		t.Fatal(err)
	}
	if string(result) != string(expect) {
		t.Fatalf("result %+v is not equal with expected %+v", result, expect)
	}
}

func TestFileManager_Write_ShouldFailedByMkdir(t *testing.T) {
	fileManager := NewManager("/sys/bus/mkdir_permission_error")
	if err := fileManager.Write([]byte("unittest")); err == nil {
		t.Fatalf("error should be returned")
	}
}
