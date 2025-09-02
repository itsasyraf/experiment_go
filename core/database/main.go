package database

import (
	fmt "fmt"
	context "context"
	log "log"
	os "os"

	pgxpool "github.com/jackc/pgx/v5/pgxpool"
)

var MainPool *pgxpool.Pool

func MainInit() {
	host := os.Getenv("POSTGRES_MAIN_HOST")
	port := os.Getenv("POSTGRES_MAIN_PORT")
	user := os.Getenv("POSTGRES_MAIN_USER")
	pass := os.Getenv("POSTGRES_MAIN_PASS_ENCODED")
	dbname := os.Getenv("POSTGRES_MAIN_DB")

	// construct DSN with fmt.Sprintf
	dsn := fmt.Sprintf(
		"postgresql://%s:%s@%s:%s/%s",
		user, pass, host, port, dbname,
	)

	var err error
	MainPool, err = pgxpool.New(context.Background(), dsn)
	if err != nil {
		log.Fatalf("⛔ Unable to connect to database: %v", err)
	}

	log.Println("✅ Connected to Postgres")
}

func MainClose() {
	if MainPool != nil {
		MainPool.Close()
		log.Println("✅ Connection to Postgres closed")
	}
}
