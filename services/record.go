package services

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"rest-docker/pkg/record"
	"time"

	goValidator "gopkg.in/validator.v2"
)

type recordService struct {
	db     *sql.DB
	record RecordResource
}

type RecordFilter struct {
	StartDate string
	EndDate   string
	MinCount  int
	MaxCount  int
}

type RecordService interface {
	GetRecords(ctx context.Context, req RecordFilter) ([]Record, ServiceError)
}

func NewRecordService(dep RecordDependencies) (RecordService, error) {
	// Validate the dependency(ies)
	if err := goValidator.Validate(dep); err != nil {
		return nil, err
	}

	return &recordService{
		db:     dep.DB,
		record: dep.Record,
	}, nil
}

func (rs *recordService) GetRecords(ctx context.Context, filter RecordFilter) ([]Record, ServiceError) {

	// Validate the start and end date
	startDate, err := time.Parse("2006-01-02", filter.StartDate)
	if err != nil {
		log.Println("Invalid start date: ", err.Error())
		return nil, setInvalidReturn(err, "Invalid start date")
	}
	endDate, err := time.Parse("2006-01-02", filter.EndDate)
	if err != nil {
		log.Println("Invalid start date: ", err.Error())
		return nil, setInvalidReturn(err, "Invalid end date")
	}

	// Validate the total min & max mark count, the value should not be less than 0
	if filter.MinCount < 0 || filter.MaxCount < 0 {
		log.Println("Invalid mark count: both minimal & maximal count must be more than 0")
		return nil, setInvalidReturn(errors.New("both minimal & maximal count must be more than 0"), "Invalid mark count")
	}

	res, err := rs.record.GetRecords(ctx, record.RecordFilter{
		StartDate: startDate,
		EndDate:   endDate,
		MinCount:  filter.MinCount,
		MaxCount:  filter.MaxCount,
	})
	if err != nil {
		log.Println("Internal system error:", err.Error())
		return nil, setSystemErrReturn(err)
	}

	var records []Record
	for _, rec := range res {
		records = append(records, Record{
			ID:        rec.ID,
			Marks:     rec.Mark,
			CreatedAt: rec.CreatedAt,
		})
	}

	return records, ServiceError{}
}
