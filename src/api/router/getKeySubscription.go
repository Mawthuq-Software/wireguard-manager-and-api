package router

import (
	"net/http"

	"gitlab.com/raspberry.tech/wireguard-manager-and-api/src/db"
	"gitlab.com/raspberry.tech/wireguard-manager-and-api/src/logger"
)

type userSubscription struct {
	KeyID string `json:"keyID"`
}

func getKeySub(res http.ResponseWriter, req *http.Request) {
	var incomingJson userSubscription
	combinedLogger := logger.GetCombinedLogger()

	err := parseResponse(req, &incomingJson) //parse JSON
	if err != nil {
		combinedLogger.Error("Parsing request " + err.Error())
		sentStandardRes(res, map[string]string{"response": err.Error()}, http.StatusBadRequest)
		return
	}

	if incomingJson.KeyID == "" {
		sentStandardRes(res, map[string]string{"response": "Bad Request, keyID needs to be filled"}, http.StatusBadRequest)
		return
	}

	boolRes, mapRes := db.GetUserSubscription(incomingJson.KeyID) //get key from db
	if !boolRes {
		sentStandardRes(res, mapRes, http.StatusBadRequest)
	} else {
		sentStandardRes(res, mapRes, http.StatusAccepted)
	}
}
