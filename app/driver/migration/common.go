package migration

import (
	database "app/app/driver/core/database"
)

func Common() {
	_, err := database.MainExec(
		// Added on 2025-09-03 11:00
		`CREATE EXTENSION IF NOT EXISTS "uuid-ossp"; ` +
		`CREATE SCHEMA IF NOT EXISTS jht_driver; `,
	)
	if err != nil {
		panic(err)
	}
}
