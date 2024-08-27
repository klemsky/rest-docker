package record

import (
	"time"

	"github.com/lib/pq"
)

type Record struct {
	ID        int64 `json:"id"`
	Mark      int64
	Marks     pq.Int64Array `json:"marks"`
	CreatedAt time.Time     `json:"createdAt"`
}
