package main

import (
	"net/http"
)

func handlerRead(w http.ResponseWriter, r *http.Request) {
	respondWithJson(w, http.StatusOK, struct{}{})
}
