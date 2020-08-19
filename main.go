package main

import (
	"encoding/json"
	"net/http"
)

func main() {
	http.HandleFunc("/", GetIPHandler)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}

// GetIPHandler handles the get ip address request
func GetIPHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	resp, _ := json.Marshal(map[string]string{
		"ip": GetIP(r),
	})
	w.Write(resp)
}

// GetIP gets a requests IP address by reading off the x-real-ip
func GetIP(r *http.Request) string {
	real := r.Header.Get("X-REAL-IP")
	if real != "" {
		return real
	}
	return r.RemoteAddr
}