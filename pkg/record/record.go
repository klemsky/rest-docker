package record

import (
	"context"
	"database/sql"
	"time"
)

type module struct {
	primary primary
}

type RecordFilter struct {
	StartDate time.Time
	EndDate   time.Time
	MinCount  int
	MaxCount  int
}

type IResource interface {
	GetRecords(ctx context.Context, filter RecordFilter) ([]Record, error)
}

func New(db *sql.DB) *module {
	return &module{
		primary: newPrimary(db),
	}
}

func (m *module) GetRecords(ctx context.Context, filter RecordFilter) ([]Record, error) {
	return m.primary.getRecords(ctx, filter)
}
