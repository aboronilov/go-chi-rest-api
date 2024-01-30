package db

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/stdlib"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/lib/pq"
)

type DB struct {
	DB *sql.DB
}

var DBconn = &DB{}

const (
	maxOpenDbConn = 10
	maxIdleDbConn = 5
	maxDbLifetime = 5 * time.Minute
	dbDriver      = "pgx"
)

func ConnectPostgres(dsn string) (*DB, error) {
	d, err := sql.Open(dbDriver, dsn)
	if err != nil {
		log.Fatalf("Unalble to connect to DB %s", err)
	}

	d.SetMaxOpenConns(maxOpenDbConn)
	d.SetMaxIdleConns(maxIdleDbConn)
	d.SetConnMaxLifetime(maxDbLifetime)

	err = testDB(d)
	if err != nil {
		return nil, err
	}

	DBconn.DB = d
	return DBconn, nil

}

func testDB(d *sql.DB) error {
	err := d.Ping()
	if err != nil {
		log.Fatalf("Unalble to ping DB %s", err)
		return err
	}

	fmt.Println("DB pinged succesfully")
	return nil
}
