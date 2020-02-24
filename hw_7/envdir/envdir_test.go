package envdir

import (
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

func TestRunCmd(t *testing.T) {
	// /Users/aleksandrteplov/Desktop/otus/hw_7/env sh -c "echo $a"
	cmd := exec.Command("/bin/bash", "build.sh")
	cmd.Stdout = os.Stdout
	cmd.Run()

}
