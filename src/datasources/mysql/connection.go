package db

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type DbPort interface {
	CloseConnection()
	GetDbConnection() *sql.DB
}

type DbAdapter struct {
	db *sql.DB
}

func NewDbAdapter(driverName, dbaseConnectionString string) *DbAdapter {
	db, err := sql.Open(driverName, dbaseConnectionString)
	if err != nil {
		log.Fatalf("could not connect to database server: %v", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("could not ping to database server: %v", err)
	}

	return &DbAdapter{db}
}

func (dba DbAdapter) CloseConnection() {
	err := dba.db.Close()
	if err != nil {
		log.Fatalf("couldnot close database connection: %v", err)
	}
}

func (dba DbAdapter) GetDbConnection() *sql.DB {
	return dba.db
}
