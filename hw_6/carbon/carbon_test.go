package carbon

import (
	"fmt"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
	"github.com/teploff/otus/hw_6/utils"
	"io/ioutil"
	"os"
	"testing"
)

var workDirectoryPath, _ = os.Getwd()

func TestIncorrectInput(t *testing.T) {
	srcUUID := uuid.NewV4()
	destUUID := uuid.NewV4()
	srcFilePath := fmt.Sprintf(workDirectoryPath+"/%s", srcUUID)
	destFilePath := fmt.Sprintf(workDirectoryPath+"/%s", destUUID)

	inputData := []struct {
		src, dest     string
		offset, limit int64
	}{
		{srcFilePath, destFilePath, 0, -5},
		{srcFilePath, destFilePath, -5, 0},
	}

	for _, data := range inputData {
		c := NewCarbon(data.src, data.dest, data.offset, data.limit)
		err := c.Copy()
		assert.Error(t, err)
		assert.Equal(t, errorInvalidValue, err)
	}
}

func TestSrsFileNotFound(t *testing.T) {
	srcUUID := uuid.NewV4()
	destUUID := uuid.NewV4()
	srcFilePath := fmt.Sprintf(workDirectoryPath+"/%s", srcUUID)
	destFilePath := fmt.Sprintf(workDirectoryPath+"/%s", destUUID)
	offset := int64(0)
	limit := int64(10)

	c := NewCarbon(srcFilePath, destFilePath, offset, limit)
	err := c.Copy()
	assert.Error(t, err)
}

func TestCopyAllPayload(t *testing.T) {
	srcUUID := uuid.NewV4()
	destUUID := uuid.NewV4()
	srcFilePath := fmt.Sprintf(workDirectoryPath+"/%s", srcUUID)
	destFilePath := fmt.Sprintf(workDirectoryPath+"/%s", destUUID)
	srcPayload := []byte("some bytes")

	err := utils.WriteFile(srcFilePath, srcPayload)
	assert.NoError(t, err)

	c := NewCarbon(srcFilePath, destFilePath, 0, 0)
	err = c.Copy()

	copyPayload, err := ioutil.ReadFile(destFilePath) // прочитать весь файл по имени

	assert.NoError(t, err)
	assert.Equal(t, srcPayload, copyPayload)

	_ = os.Remove(srcFilePath)
	_ = os.Remove(destFilePath)
}

func TestCopyAllPayloadFromEmptyFile(t *testing.T) {
	srcUUID := uuid.NewV4()
	destUUID := uuid.NewV4()
	srcFilePath := fmt.Sprintf(workDirectoryPath+"/%s", srcUUID)
	destFilePath := fmt.Sprintf(workDirectoryPath+"/%s", destUUID)
	srcPayload := []byte("")

	_ = utils.WriteFile(srcFilePath, srcPayload)

	c := NewCarbon(srcFilePath, destFilePath, 0, 5)
	err := c.Copy()

	copyPayload, err := ioutil.ReadFile(destFilePath) // прочитать весь файл по имени

	assert.NoError(t, err)
	assert.Equal(t, srcPayload, copyPayload)

	_ = os.Remove(srcFilePath)
	_ = os.Remove(destFilePath)
}

func TestCopyWithOffset(t *testing.T) {
	srcUUID := uuid.NewV4()
	destUUID := uuid.NewV4()
	srcFilePath := fmt.Sprintf(workDirectoryPath+"/%s", srcUUID)
	destFilePath := fmt.Sprintf(workDirectoryPath+"/%s", destUUID)
	srcPayload := []byte("1234567890")

	_ = utils.WriteFile(srcFilePath, srcPayload)

	c := NewCarbon(srcFilePath, destFilePath, 1, 10)
	err := c.Copy()

	copyPayload, err := ioutil.ReadFile(destFilePath) // прочитать весь файл по имени

	assert.NoError(t, err)
	assert.Equal(t, []byte("234567890"), copyPayload)

	_ = os.Remove(srcFilePath)
	_ = os.Remove(destFilePath)
}

func Test1CopyWithOffset(t *testing.T) {
	srcUUID := uuid.NewV4()
	destUUID := uuid.NewV4()
	srcFilePath := fmt.Sprintf(workDirectoryPath+"/%s", srcUUID)
	destFilePath := fmt.Sprintf(workDirectoryPath+"/%s", destUUID)
	srcPayload := []byte("12345678901234567890123456789012345678901234567890")

	_ = utils.WriteFile(srcFilePath, srcPayload)

	c := NewCarbon(srcFilePath, destFilePath, 10, 0)
	err := c.Copy()

	copyPayload, err := ioutil.ReadFile(destFilePath) // прочитать весь файл по имени

	assert.NoError(t, err)
	assert.Equal(t, []byte("1234567890123456789012345678901234567890"), copyPayload)

	_ = os.Remove(srcFilePath)
	_ = os.Remove(destFilePath)
}
