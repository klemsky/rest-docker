package resthttp

import (
	"context"
	"rest-docker/services"
)

type RecordService interface {
	GetRecords(ctx context.Context, filter services.RecordFilter) ([]services.Record, services.ServiceError)
}
