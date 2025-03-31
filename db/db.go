package db

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"kvm-manager/types"
)

var conn *sql.DB

func Init() {
	var err error
	conn, err = sql.Open("sqlite3", "file:kvm_manager.db?cache=shared&mode=rwc")
	if err != nil {
		log.Fatal("DB connection failed:", err)
	}

	createJobs := `CREATE TABLE IF NOT EXISTS job_history (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		in_progress BOOLEAN,
		last_step TEXT,
		started_at DATETIME,
		ended_at DATETIME,
		error TEXT
	)`

	createCustomers := `CREATE TABLE IF NOT EXISTS customers (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT,
		wallet_address TEXT,
		contract TEXT,
		hashrate TEXT,
		start_date DATETIME,
		end_date DATETIME,
		email TEXT
	)`

	if _, err = conn.Exec(createJobs); err != nil {
		log.Fatal("Failed creating job_history table:", err)
	}
	if _, err = conn.Exec(createCustomers); err != nil {
		log.Fatal("Failed creating customers table:", err)
	}
}

func SaveJobStatus(status *types.JobStatus) {
	stmt, _ := conn.Prepare("INSERT INTO job_history (in_progress, last_step, started_at, ended_at, error) VALUES (?, ?, ?, ?, ?)")
	_, err := stmt.Exec(status.InProgress, status.LastStep, status.StartedAt, status.EndedAt, status.Error)
	if err != nil {
		log.Printf("Failed to save job status: %v", err)
	}
}

func FetchJobHistory(limit int) []types.JobStatus {
	rows, err := conn.Query("SELECT in_progress, last_step, started_at, ended_at, error FROM job_history ORDER BY started_at DESC LIMIT ?", limit)
	if err != nil {
		log.Println("Failed fetching job history:", err)
		return nil
	}
	defer rows.Close()

	var history []types.JobStatus
	for rows.Next() {
		var js types.JobStatus
		if err := rows.Scan(&js.InProgress, &js.LastStep, &js.StartedAt, &js.EndedAt, &js.Error); err == nil {
			history = append(history, js)
		}
	}
	return history
}

func AddCustomer(c types.Customer) error {
	stmt, err := conn.Prepare(`INSERT INTO customers (name, wallet_address, contract, hashrate, start_date, end_date, email) VALUES (?, ?, ?, ?, ?, ?, ?)`)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(c.Name, c.WalletAddress, c.Contract, c.Hashrate, c.StartDate, c.EndDate, c.Email)
	return err
}
