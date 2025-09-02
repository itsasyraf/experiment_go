package core

import (
	json "encoding/json"
	// log "log"
	sql "database/sql"

	// _ "github.com/marcboeker/go-duckdb"
)

func InitDB() {
	return
}
// func DBExec(db *sql.DB, query string, args ...interface`{}) (error) {
// 	result, err = db.Exec(query, args...)
// 	if err != nil {
// 		return "", err
// 	}
// 	return result, nil
// }

// DBFetch executes SELECT and returns JSON string
func DBFetch(db *sql.DB, query string, args ...interface{}) (string, error) {
	rows, err := db.Query(query, args...)
	if err != nil {
		return "", err
	}
	defer rows.Close()

	cols, err := rows.Columns()
	if err != nil {
		return "", err
	}

	var results []map[string]interface{}

	for rows.Next() {
		values := make([]interface{}, len(cols))
		valuePtrs := make([]interface{}, len(cols))
		for i := range values {
			valuePtrs[i] = &values[i]
		}

		if err := rows.Scan(valuePtrs...); err != nil {
			return "", err
		}

		rowMap := make(map[string]interface{})
		for i, col := range cols {
			val := values[i]
			if b, ok := val.([]byte); ok {
				rowMap[col] = string(b)
			} else {
				rowMap[col] = val
			}
		}
		results = append(results, rowMap)
	}

	if err := rows.Err(); err != nil {
		return "", err
	}

	jsonBytes, err := json.MarshalIndent(results, "", "  ")
	if err != nil {
		return "", err
	}

	return string(jsonBytes), nil
}
