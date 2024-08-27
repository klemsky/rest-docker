package resthttp

import (
	"context"
	"encoding/json"
	"net/http"
	"rest-docker/services"
	"rest-docker/utils"
)

type recordHandler struct {
	service RecordService
}

type RecordFilter struct {
	StartDate string `json:"startDate"`
	EndDate   string `json:"endDate"`
	MinCount  int    `json:"minCount"`
	MaxCount  int    `json:"maxCount"`
}

func newRecordHandler(service RecordService) *recordHandler {
	return &recordHandler{
		service: service,
	}
}

func (rh recordHandler) getRecords(w http.ResponseWriter, r *http.Request) {

	var filter RecordFilter
	err := json.NewDecoder(r.Body).Decode(&filter)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, 1, "Invalid request payload")
		return
	}

	records, svcErr := rh.service.GetRecords(context.Background(), services.RecordFilter{
		StartDate: filter.StartDate,
		EndDate:   filter.EndDate,
		MinCount:  filter.MinCount,
		MaxCount:  filter.MaxCount,
	})
	if svcErr.Err != nil && svcErr.IsInvalidErr() {
		utils.RespondWithError(w, http.StatusInternalServerError, svcErr.Code, "Invalid error: "+svcErr.Msg+": "+svcErr.Err.Error())
		return
	}
	if svcErr.Err != nil && svcErr.IsSystemErr() {
		utils.RespondWithError(w, http.StatusInternalServerError, svcErr.Code, "Internal system error: "+svcErr.Err.Error())
		return
	}

	response := map[string]interface{}{
		"code":    0,
		"msg":     "Success",
		"records": records,
	}
	utils.RespondWithJSON(w, http.StatusOK, response)
}
