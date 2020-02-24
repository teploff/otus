package envdir

import (
	"github.com/stretchr/testify/assert"
	"os"
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
		"a":        "1",
		"b":        "2",
		"c":        "3",
		"d":        "4",
		"LOCALDIR": ".",
	}

	assert.Equal(t, expected, env)
}
