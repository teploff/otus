package utils

import (
	"bufio"
	"os"
)

func WriteFile(filePath string, payload []byte) error {
	src, err := os.Create(filePath)
	if err != nil {
		return err
	}

	bw := bufio.NewWriter(src)
	if _, err = bw.Write(payload); err != nil {
		return err
	}

	if err = bw.Flush(); err != nil {
		return err
	}
	return nil
}
