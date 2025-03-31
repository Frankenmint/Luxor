// This is the entry point for the Luxor KVM Manager CLI/API
package main

import (
	"log"
	"net/http"

	"kvm-manager/api"
	"kvm-manager/db"
	"kvm-manager/internal"
)

func main() {
	db.Init()
	internal.InitProvisioning()

	http.HandleFunc("/provision", api.ProvisionHandler)
	http.HandleFunc("/status", api.StatusHandler)
	http.HandleFunc("/history", api.HistoryHandler)

	log.Println("Luxor KVM Manager API running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
