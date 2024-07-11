package main

import "net/http"

func handlerhandlerErrReadiness(w http.ResponseWriter, r *http.Request) {
	respondWithError(w, 400, "Something went wrong")
}
