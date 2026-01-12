package handlers

import (
	"encoding/json"
	"github.com/dexterhere04/Learn_go/api"
	"github.com/dexterhere04/Learn_go/internal/tools"
	"github.com/gorilla/schema"
	"net/http"
)

func GetCoinBalance(w http.ResponseWriter, r *http.Request) {
	var params = api.CoinBalanceRequest{}
	var decoder *schema.Decoder = schema.NewDecoder()
	var err error
	err = decoder.Decode(&params, r.URL.Query())
	if err != nil {
		api.RequestInvalidError(w, err)
		return
	}

	var database tools.DatabaseInterface
	database, err = tools.NewDatabase()
	if err != nil {
		api.InternalErrorHandler(w, err)
		return
	}
	var tokenDetails *tools.CoinDetails
	tokenDetails = database.GetCoinBalance(params.Username)
	if tokenDetails == nil {
		api.RequestInvalidError(w, err)
		return
	}

	var resp = api.CoinBalanceResponse{
		Code:    http.StatusOK,
		Balance: tokenDetails.Balance,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}
