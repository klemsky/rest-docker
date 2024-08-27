package record

import (
	"context"
	"database/sql"
)

type primary interface {
	getRecords(ctx context.Context, filter RecordFilter) ([]Record, error)
}

type psql struct {
	db *sql.DB
}

func newPrimary(db *sql.DB) primary {
	return &psql{
		db: db,
	}
}

func (p *psql) getRecords(ctx context.Context, filter RecordFilter) ([]Record, error) {

	// Select query using filter range of date
	query := `SELECT id, marks, created_at FROM records WHERE created_at BETWEEN $1 AND $2`

	rows, err := p.db.Query(query, filter.StartDate, filter.EndDate)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var records []Record

	// Loop through all the data, if the data is not between the min & max total mark, then do not fetch the data to the service
	for rows.Next() {
		var record Record
		if err := rows.Scan(&record.ID, &record.Marks, &record.CreatedAt); err != nil {
			return nil, err
		}

		// Sum every score mark on a single student
		mark := sum(record.Marks)
		if mark >= int64(filter.MinCount) && mark <= int64(filter.MaxCount) {
			record.Mark = mark
			records = append(records, record)
		}
	}

	return records, nil
}

// Helper function to sum the total of marks the student get
func sum(marks []int64) int64 {
	var total int64
	for _, mark := range marks {
		total += mark
	}
	return total
}
