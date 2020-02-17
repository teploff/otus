package carbon

import (
	"errors"
	"io"
	"os"
)

var (
	errorInvalidValue = errors.New("invalid value")
	errorHugeOffset   = errors.New("offset more source file size")
)

type Carbon struct {
	srcFilePath, destFilePath string
	offset, limit             int64
}

func NewCarbon(src, dest string, offset, limit int64) Carbon {
	return Carbon{
		srcFilePath:  src,
		destFilePath: dest,
		offset:       offset,
		limit:        limit,
	}
}

func (c Carbon) Copy() error {
	if err := c.validate(); err != nil {
		return err
	}

	srcFile, err := os.Open(c.srcFilePath)
	if err != nil {
		return err
	}
	defer srcFile.Close()
	srcInfo, err := srcFile.Stat()
	if err != nil {
		return err
	}
	srcLength := srcInfo.Size()

	destFile, err := os.Open(c.destFilePath)
	if err != nil {
		destFile, err = os.Create(c.destFilePath)
		if err != nil {
			return err
		}
	}
	defer destFile.Close()

	if srcLength == 0 {
		return nil
	}

	if c.offset > srcLength {
		return errorHugeOffset
	}

	if c.limit == 0 {
		c.limit = srcLength
	}

	buffer := make([]byte, c.limit)
	var offset int64
	for offset < c.limit {
		var read int
		if offset == 0 {
			read, err = srcFile.ReadAt(buffer, c.offset)
		} else {
			read, err = srcFile.Read(buffer)
		}
		offset += int64(read)
		if err == io.EOF {
			if _, err = destFile.Write(buffer[:read]); err != nil {
				return err
			}
			break
		}
		if err != nil {
			return err
		}

		if _, err = destFile.Write(buffer[:read]); err != nil {
			return err
		}
	}

	return nil
}

func (c Carbon) validate() error {
	if c.limit < 0 || c.offset < 0 {
		return errorInvalidValue
	}

	return nil
}
