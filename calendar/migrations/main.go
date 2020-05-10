package main

import (
	"database/sql"
	"flag"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/pressly/goose"
	"log"
)

var (
	dir      = flag.String("dir", ".", "directory with migration files")
	host     = flag.String("host", "127.0.0.1", "host of postgreSQL")
	port     = flag.Int("port", 5438, "port of postgreSQL")
	user     = flag.String("user", "postgres", "user of postgreSQL")
	password = flag.String("password", "password", "password of postgreSQL")
	dbname   = flag.String("dbname", "otus", "database name of postgreSQL")
	sslmode  = flag.String("sslmode", "disable", "sslmode of postgreSQL")
)

// "port=5438 host=127.0.0.1 user=postgres password=password dbname=otus sslmode=disable"
func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) < 1 {
		flag.Usage()
		return
	}

	connString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s", *host, *port, *user,
		*password, *dbname, *sslmode)

	command := args[0]
	db, err := sql.Open("postgres", connString)
	if err != nil {
		log.Fatalf("goose: failed to open DB: %v\n", err)
	}

	defer func() {
		if err = db.Close(); err != nil {
			log.Fatalf("goose: failed to close DB: %v\n", err)
		}
	}()

	var arguments []string
	if len(args) > 3 {
		arguments = append(arguments, args[3:]...)
	}

	if err = goose.Run(command, db, *dir, arguments...); err != nil {
		log.Fatalf("goose %v: %v", command, err)
	}
}
