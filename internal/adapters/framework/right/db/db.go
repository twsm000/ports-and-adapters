package db

import (
	"database/sql"
	sq "github.com/Masterminds/squirrel"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

type Adapter struct {
	db *sql.DB
}

func NewAdapter(driverName, dataSourceName string) (*Adapter, error) {
	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		log.Fatalf("DB connection failure: %v", err)
	}

	// test db connection
	err = db.Ping()
	if err != nil {
		log.Fatalf("DB ping failure: %v", err)
	}

	return &Adapter{db: db}, nil
}

func (adp *Adapter) Close() {
	log.Println("db.Close: Closing DB connection")
	err := adp.db.Close()
	if err != nil {
		log.Fatalf("DB close failure: %v", err)
	}
}

func (adp *Adapter) AddHistory(answer int32, operation string) error {
	query, args, err := sq.Insert("arith_history").Columns("date", "answer", "operation").
		Values(time.Now(), answer, operation).
		ToSql()
	if err != nil {
		return err
	}

	_, err = adp.db.Exec(query, args...)
	if err != nil {
		return err
	}

	return nil
}
