package gUtils

import "database/sql"

func Query(db *sql.DB, s string) (error, [] string, []map[string]string) {
	records := make([]map[string]string, 0)
	rows, err := db.Query(s)
	if err != nil {
		return err, nil, nil
	}
	columns, _ := rows.Columns()
	scanArgs := make([]interface{}, len(columns))
	values := make([]interface{}, len(columns))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	fields := make([]string, 0)

	for _, e := range columns {
		fields = append(fields, e)
	}

	for rows.Next() {
		err = rows.Scan(scanArgs...)
		record := make(map[string]string, 0)
		for i, col := range values {
			if col != nil {
				record[columns[i]] = string(col.([]byte))
			}
		}
		records = append(records, record)
	}
	return err, fields, records
}
