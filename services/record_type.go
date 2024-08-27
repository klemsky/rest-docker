package services

import (
	"database/sql"
	"time"
)

type RecordDependencies struct {
	DB     *sql.DB
	Record RecordResource
}

type Record struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	Marks     int64     `json:"totalMarks"`
}
