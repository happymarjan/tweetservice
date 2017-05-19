package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

//----------------------------------------------

const (
	host          = "localhost"
	postgres_port = 5432
	user          = "postgres"
	password      = "postgres"
	dbname        = "tweetsdb"
)

var (
	dbObject *DB
)

//----------------------------------------------

func main() {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, postgres_port, user, password, dbname)
	db := getDB()
	if err := db.Open(connStr); err != nil {
		log.Fatalf("DB open failed: %v.\n", err)
	}
	defer db.Close()

	db.InitTables(false)
	//db.populateTable()

	errs := make(chan error)

	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	go restServerMain(db)
	fmt.Println("main rest started")
	duration := time.Duration(8) * time.Second // Pause for 10 seconds
	time.Sleep(duration)
	fmt.Println("main wake up")
	go grpcMain(db)
	fmt.Println("main GRPC started")
	duration = time.Duration(8) * time.Second // Pause for 10 seconds
	time.Sleep(duration)
	log.Fatalf("%s", <-errs)
}
