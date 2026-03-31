package main

import (
	"net/http"
	// "structs"
)

func handlerReadiness(w http.ResponseWriter, r *http.Request){
	 
	respondWithJSON(w, 200, struct{}{})
}