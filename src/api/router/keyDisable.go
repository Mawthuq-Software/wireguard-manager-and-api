package router

import (
	"net/http"

	"gitlab.com/raspberry.tech/wireguard-manager-and-api/src/db"
	"gitlab.com/raspberry.tech/wireguard-manager-and-api/src/logger"
)

type keyDisableJSON struct {
	KeyID string `json:"keyID"`
}

func keyDisable(res http.ResponseWriter, req *http.Request) {
	var incomingJson keyDisableJSON
	combinedLogger := logger.GetCombinedLogger()

	jsonResponse := make(map[string]string)

	err := parseResponse(req, &incomingJson)
	if err != nil {
		combinedLogger.Error(err.Error())
		sentStandardRes(res, map[string]string{"response": err.Error()}, http.StatusBadRequest)
		return
	}

	if incomingJson.KeyID == "" {
		jsonResponse["response"] = "Bad Request, keyID needs to be filled"
		sentStandardRes(res, map[string]string{"response": "Bad Request, keyID needs to be filled"}, http.StatusBadRequest)
		return
	}

	keyID := incomingJson.KeyID

	boolRes, mapRes := db.DisableKey(keyID)
	if !boolRes {
		sentStandardRes(res, mapRes, http.StatusBadRequest)
	} else {
		sentStandardRes(res, mapRes, http.StatusAccepted)
	}
}
