package services

import (
	"context"
	"rest-docker/pkg/record"
)

type RecordResource interface {
	GetRecords(ctx context.Context, filter record.RecordFilter) ([]record.Record, error)
}
