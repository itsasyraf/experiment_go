package database

import (
	log "log"
	context "context"
	time "time"
)

func MainExec(query string, args ...any) (int64, error) {
	MainInit()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cmd, err := MainPool.Exec(ctx, query, args...)
	if err != nil {
		return 0, err
	}
	log.Printf("✔️ Query exec Postgres [Main]: %s", query)

	MainClose()
	return cmd.RowsAffected(), nil
}
