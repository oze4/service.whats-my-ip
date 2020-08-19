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
	resp, _ := GetIP(r)
	w.Write(resp)
}

// GetIP gets a requests IP address by reading off the x-real-ip
func GetIP(r *http.Request) ([]byte, error) {
	return json.Marshal(map[string]string{
		"realIp":        r.Header.Get("x-forwarded-for"),
		"forwardedFor":  r.Header.Get("x-real-ip"),
		"remoteAddress": r.RemoteAddr,
	})
}
