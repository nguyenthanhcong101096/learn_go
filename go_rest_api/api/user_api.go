package api

import (
	"encoding/json"
	repo "go_rest_api/repository"
	"net/http"
)

func FindUser(response http.ResponseWriter, request *http.Request) {
	ids, ok := request.URL.Query()["id"]

	if !ok || len(ids) < 1 {
		responseWithError(response, http.StatusBadRequest, "Url Param id is missing")
		return
	}

	user, err := repo.FindUser(ids[0])
	if err != nil {
		responseWithError(response, http.StatusBadRequest, err.Error())
		return
	}
	responseWithJSON(response, http.StatusOK, user)
}

func responseWithError(response http.ResponseWriter, statusCode int, msg string) {
	responseWithJSON(response, statusCode, map[string]string{
		"error": msg,
	})
}

func responseWithJSON(response http.ResponseWriter, statusCode int, data interface{}) {
	result, _ := json.Marshal(data)
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(statusCode)
	response.Write(result)
}
