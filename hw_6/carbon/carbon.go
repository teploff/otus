package carbon

import (
	"errors"
	pb2 "github.com/cheggaaa/pb/v3"
	"github.com/teploff/otus/hw_6/utils"
	"io"
	"os"
)

var (
	// see https://eklitzke.org/efficient-file-copying-on-linux
	efficientBufferSize = 128 * utils.KILOBYTE
	errorInvalidValue   = errors.New("invalid value")
	errorHugeOffset     = errors.New("offset more source file size")
)

// Carbon is a file copy utility
//
// srcFile - File source
//
// destFile - File copy
//
// offset - offset from the file source
//
// limit - limit to copy from the file source
//
// pb - progress bar to show how many bytes to copy
type Carbon struct {
	srcFile, destFile *os.File
	offset, limit     int64
	pb                *pb2.ProgressBar
}

// NewCarbon gets Carbon instance
//
// srcFilePath - File source path
//
// destFilePath - File copy path
//
// offset - offset from the file source
//
// limit - limit to copy from the file source
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

// Copy copies bytes from source file to copy file
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

	c.pb = pb2.Full.Start64(srcLength)
	c.pb.Set(pb2.Bytes, true)

	buffer := make([]byte, efficientBufferSize)
	var offset int64
	for offset < c.limit {
		var read int
		if offset == 0 {
			c.pb.Add64(c.offset)
			read, err = c.srcFile.ReadAt(buffer, c.offset)
		} else {
			read, err = c.srcFile.Read(buffer)
		}
		offset += int64(read)
		if err == io.EOF {
			if _, err = c.destFile.Write(buffer[:read]); err != nil {
				return err
			}
			c.pb.Add(read)
			break
		}
		if err != nil {
			return err
		}

		if _, err = c.destFile.Write(buffer[:read]); err != nil {
			return err
		}
		c.pb.Add(read)
	}
	c.pb.Finish()

	return nil
}

// validate validates limit & offset passed params
func validate(limit, offset int64) error {
	if limit < 0 || offset < 0 {
		return errorInvalidValue
	}

	return nil
}

// filesInit initializes source and copy files
func filesInit(srcFilePath, destFilePath string) (src, dest *os.File, err error) {
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
