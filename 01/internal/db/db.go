package db

import (
	"context"
	"database/sql"
	"log"
	"time"
)

func New(addr string, maxOpenCons, maxIdleCons int, maxIdleTime time.Duration) (*sql.DB, error) {
	db, err := sql.Open("postgres", addr)
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	db.SetMaxOpenConns(maxOpenCons)
	db.SetConnMaxIdleTime(maxIdleTime)
	db.SetConnMaxLifetime(time.Duration(maxIdleCons))

	if err := db.PingContext(ctx); err != nil {
		return nil, err
	}
	log.Println("connected to postgres db.")
	return db, nil
}
