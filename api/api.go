package api

import (
	"encoding/json"
	"net/http"
)

// Coin Balance Params
type CoinBalanceParamas struct {
	Username string
}

// Coin Balance Response
type CoinBalanceResponse struct {
	// success Code, Usally 200
	Code int

	//Account Balance
	Balance int64
}

// Error response
type Error struct {
	// error Code
	code int

	//Error message
	message string
}

func writeError(w http.ResponseWriter, message string, code int) {
	resp := Error{
		code:    code,
		message: message,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	json.NewEncoder(w).Encode(resp)
}

var (
	RequestErrorHandler = func(w http.ResponseWriter, r *http.Request) {
		writeError(w, http.StatusBadRequest)
	}
	InternalErrorHandler = func(w http.ResponseWriter, r *http.Request) {
		writeError(w, "An Unexpected Error Occured", http.StatusInternalServerError)
	}
)
