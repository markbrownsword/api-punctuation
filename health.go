package main

import (
	"encoding/json"
	"net/http"
)

type appInfo struct {
	Version   string `json:"version"`
	IsHealthy bool   `json:"ishealthy"`
}

func ping(w http.ResponseWriter, r *http.Request) {
	info := appInfo{Version: "1.0", IsHealthy: true}
	healthBytes, err := json.Marshal(info)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(healthBytes)
}
