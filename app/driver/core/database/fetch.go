package database

import (
	context "context"
	time "time"
	log "log"
)

func MainFetch(query string, args ...any) ([]map[string]interface{}, error) {
	MainInit()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	rows, err := MainPool.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	log.Printf("✔️ Query fetch Postgres [Main]: %s", query)

	// get column names
	fieldDescriptions := rows.FieldDescriptions()
	var result []map[string]interface{}

	for rows.Next() {
		values, err := rows.Values()
		if err != nil {
			return nil, err
		}

		rowMap := make(map[string]interface{})
		for i, fd := range fieldDescriptions {
			rowMap[string(fd.Name)] = values[i]
		}

		result = append(result, rowMap)
	}
	MainClose()
	return result, nil
}
