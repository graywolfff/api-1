package main

import (
	"database/sql"
	"log"
	"time"
	"web_study/01/internal/db"
	"web_study/01/internal/env"
	"web_study/01/internal/store"
)

func main() {
	cfg := &config{
		addr: env.GetString("ADDR", ":8080"),
		db: dbConfig{
			addr:        env.GetString("DB_ADDR", "postgresql://postgres:pass@localhost:5432/social?sslmode=disable"),
			maxOpenCons: env.GetInt("DB_MAX_OPEN_CONS", 30),
			maxIdleCons: env.GetInt("DB_MAX_IDLE_CONS", 30),
			maxIdleTime: time.Duration(env.GetInt("DB_MAX_IDLE_TIME", 30)),
		},
	}
	dbConnection, err := db.New(
		cfg.db.addr,
		cfg.db.maxOpenCons,
		cfg.db.maxIdleCons,
		cfg.db.maxIdleTime,
	)
	if err != nil {
		log.Panic(err)
	}
	defer func(dbConnection *sql.DB) {
		err := dbConnection.Close()
		if err != nil {
			log.Fatal("failed to close db connection")
		}
	}(dbConnection)
	dbStore := store.NewStore(dbConnection)
	app := &application{
		config: *cfg,
		store:  dbStore,
	}
	mux := app.mount()

	log.Fatal(app.run(mux))
}
