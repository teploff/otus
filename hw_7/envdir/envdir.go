package envdir

import (
	"fmt"
	"github.com/imdario/mergo"
	"github.com/joho/godotenv"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
)

// according exit codes https://www.unix.com/man-page/debian/8/envdir/
const exitCode = 111

// ReadDir scans the specified directory and returns all environment variables defined in it
func ReadDir(dir string) (map[string]string, error) {
	environments := make(map[string]string, 0)

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	filesPath := make([]string, 0, len(files))
	for _, file := range files {
		if !file.IsDir() {
			filesPath = append(filesPath, filepath.Join(dir, file.Name()))
		}
	}

	for _, path := range filesPath {
		env, err := godotenv.Read(path)
		if err != nil {
			return nil, err
		}
		if err := mergo.Merge(&environments, env); err != nil {
			return nil, err
		}
	}

	return environments, nil
}

// RunCmd starts a program with arguments (cmd) with an overridden environment.
func RunCmd(cmd []string, env map[string]string) int {

	var command *exec.Cmd
	if len(cmd) == 1 {
		command = exec.Command(cmd[0])
	} else {
		command = exec.Command(cmd[0], cmd[1:]...)
	}

	sliceEnv := make([]string, 0, len(env))
	for name, value := range env {
		sliceEnv = append(sliceEnv, fmt.Sprintf("%s=%s", name, value))
	}
	command.Env = append(os.Environ(), sliceEnv...)
	command.Stdin = os.Stdin
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr

	if err := command.Run(); err != nil {
		return exitCode
	}

	return 0
}
