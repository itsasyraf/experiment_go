package database

import (
	fmt "fmt"
	context "context"
	log "log"
	url "net/url"

	pgxpool "github.com/jackc/pgx/v5/pgxpool"
	environment "app/app/driver/core/environment"
)

var MainPool *pgxpool.Pool

func MainInit() {
	host := environment.Mandatory("POSTGRES_MAIN_HOST")
	port := environment.Mandatory("POSTGRES_MAIN_PORT")
	user := environment.Mandatory("POSTGRES_MAIN_USER")
	pass := url.PathEscape(environment.Mandatory("POSTGRES_MAIN_PASS")) // encodes for path segments
	dbname := environment.Mandatory("POSTGRES_MAIN_DB")

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

	log.Println("✅ Connected to Postgres [Main]")
}

func MainClose() {
	if MainPool != nil {
		MainPool.Close()
		log.Println("✅ Connection to Postgres closed [Main]")
	}
}
