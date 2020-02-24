package envdir

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

var workDirectoryPath, _ = os.Getwd()

// TestCase try to get environments from absent folder
func TestReadUnknownDir(t *testing.T) {
	unknownPath := filepath.Join(workDirectoryPath, "unknown_folder/")

	env, err := ReadDir(unknownPath)
	assert.Nil(t, env)
	assert.Error(t, err)
}

// TestCase getting all environments from several files
func TestReadDir(t *testing.T) {
	pathToEnv, err := filepath.Abs("../env")
	assert.NoError(t, err)
	assert.NotEmpty(t, pathToEnv)

	env, err := ReadDir(pathToEnv)
	assert.NoError(t, err)
	assert.NotNil(t, env)

	expected := map[string]string{
		"A": "1",
		"B": "2",
		"C": "3",
		"D": "4",
	}

	assert.Equal(t, expected, env)
}

// TestCase running cmd with reading environments from the directory. Cmd and directory path indicated in build.sh
func TestRunCmd(t *testing.T) {
	buf := new(bytes.Buffer)

	commandA := exec.Command("/bin/bash", "build.sh")
	commandA.Stdout = buf
	assert.NoError(t, commandA.Run())
	expected := "1 2 3 4\n"
	assert.Equal(t, expected, buf.String())

}
