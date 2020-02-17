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
	srcFile, destFile *os.File
	offset, limit     int64
}

func NewCarbon(srcFilePath, destFilePath string, offset, limit int64) (Carbon, error) {
	if err := validate(offset, limit); err != nil {
		return Carbon{}, err
	}

	src, dest, err := filesInit(srcFilePath, destFilePath)
	if err != nil {
		return Carbon{}, err
	}

	return Carbon{
		srcFile:  src,
		destFile: dest,
		offset:   offset,
		limit:    limit,
	}, nil
}

func (c *Carbon) Copy() error {
	defer c.srcFile.Close()
	defer c.destFile.Close()

	srcInfo, err := c.srcFile.Stat()
	if err != nil {
		return err
	}
	srcLength := srcInfo.Size()
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
			read, err = c.srcFile.ReadAt(buffer, c.offset)
		} else {
			read, err = c.srcFile.Read(buffer)
		}
		offset += int64(read)
		if err == io.EOF {
			if _, err = c.destFile.Write(buffer[:read]); err != nil {
				return err
			}
			break
		}
		if err != nil {
			return err
		}

		if _, err = c.destFile.Write(buffer[:read]); err != nil {
			return err
		}
	}

	return nil
}

func validate(limit, offset int64) error {
	if limit < 0 || offset < 0 {
		return errorInvalidValue
	}

	return nil
}

func filesInit(srcFilePath, destFilePath string) (src *os.File, dest *os.File, err error) {
	src, err = os.Open(srcFilePath)
	if err != nil {
		return nil, nil, err
	}

	dest, err = os.Open(destFilePath)
	if err != nil {
		dest, err = os.Create(destFilePath)
		if err != nil {
			return nil, nil, err
		}
	}

	return src, dest, nil
}
