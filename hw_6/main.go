package main

import (
	"fmt"
	uuid "github.com/satori/go.uuid"
	"github.com/teploff/otus/hw_6/carbon"
	"github.com/teploff/otus/hw_6/utils"
	"log"
	"os"
)

func main() {
	workDirectoryPath, err := os.Getwd()
	if err != nil {
		log.Fatalln(err)
	}

	srcUUID := uuid.NewV4()
	destUUID := uuid.NewV4()
	srcFilePath := fmt.Sprintf(workDirectoryPath+"/%s", srcUUID)
	destFilePath := fmt.Sprintf(workDirectoryPath+"/%s", destUUID)
	srcPayload := make([]byte, utils.GIGABYTE*5) // заполнен нулями

	if err := utils.WriteFile(srcFilePath, srcPayload); err != nil {
		log.Fatal(err)
	}
	c, err := carbon.NewCarbon(srcFilePath, destFilePath, 0, 0)
	if err != nil {
		log.Fatalln(err)
	}

	if err = c.Copy(); err != nil {
		log.Fatalln(err)
	}

	if err = os.Remove(srcFilePath); err != nil {
		log.Fatalln(err)
	}
	if err = os.Remove(destFilePath); err != nil {
		log.Fatalln(err)
	}
}
