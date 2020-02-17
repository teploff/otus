package utils

import (
	"bufio"
	"os"
)

const (
	_ = 1 << (10 * iota)
	// KILOBYTE has 1024 bytes
	KILOBYTE
	_
	// GIGABYTE has 1024 * 1024 * 1024 bytes
	GIGABYTE
)

// WriteFile helps to save bytes payload to the file with filePath path.
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
