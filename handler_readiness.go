package main

import "net/http"

func handlerReadiness(w http.ResponseWriter, r *http.Request) {
	respondwithJOSN(w,200, struct{}{})
}