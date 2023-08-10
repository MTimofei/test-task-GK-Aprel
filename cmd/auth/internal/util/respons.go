package util

import "net/http"

func Response(w http.ResponseWriter, statusCode int, msg []byte) {
	w.WriteHeader(statusCode)
	w.Write(msg)
}
