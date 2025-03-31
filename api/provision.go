package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync/atomic"
	"time"

	"kvm-manager/db"
	"kvm-manager/internal"
	"kvm-manager/types"
)

var isProvisioning atomic.Bool
var jobStatus = &types.JobStatus{}

func ProvisionHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	auth := r.Header.Get("Authorization")
	if auth != fmt.Sprintf("Bearer %s", internal.Config.AuthToken) {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	if isProvisioning.Load() {
		http.Error(w, "Provisioning already in progress", http.StatusConflict)
		return
	}

	isProvisioning.Store(true)
	jobStatus = &types.JobStatus{
		InProgress: true,
		StartedAt:  time.Now(),
		LastStep:   "received",
		Error:      "",
	}

	go func() {
		defer func() {
			jobStatus.InProgress = false
			jobStatus.EndedAt = time.Now()
			db.SaveJobStatus(jobStatus)
			isProvisioning.Store(false)
		}()

		err := internal.OrchestrateProvisioning(jobStatus)
		if err != nil {
			jobStatus.Error = err.Error()
		}
	}()

	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(map[string]string{"message": "Provisioning started"})
}

func StatusHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(jobStatus)
}

func HistoryHandler(w http.ResponseWriter, r *http.Request) {
	history := db.FetchJobHistory(20)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(history)
}
