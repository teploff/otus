package main

import (
	"flag"
	"github.com/teploff/otus/hw_6/carbon"
	"log"
)

var (
	srsFilePath  = flag.String("src", "full/source/file/path", "full path to src file")
	destFilePath = flag.String("dest", "full/dest/file/path", "full path to dest file")
	offset       = flag.Int64("offset", 0, "offset count bytes from source file")
	limit        = flag.Int64("limit", 0, "count of bytes which should copied from source file")
)

func main() {
	flag.Parse()

	c, err := carbon.NewCarbon(*srsFilePath, *destFilePath, *offset, *limit)
	if err != nil {
		log.Fatalln(err)
	}

	if err = c.Copy(); err != nil {
		log.Fatalln(err)
	}
}
