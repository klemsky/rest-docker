package resthttp

import (
	"net/http"

	"github.com/gorilla/mux"
)

type RouterDependencies struct {
	RS RecordService
}

func NewRouter(rd RouterDependencies) *mux.Router {
	r := mux.NewRouter()

	rh := newRecordHandler(rd.RS)

	r.HandleFunc("/api/records", rh.getRecords).Methods(http.MethodPost)

	return r
}
