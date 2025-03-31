package types

import "time"

type Config struct {
	Region       string `json:"region"`
	S3BucketName string `json:"s3_bucket_name"`
	SalesTeamNet string `json:"sales_team_net"`
	EmailDomain  string `json:"email_domain"`
	AuthToken    string `json:"auth_token"`
}

type JobStatus struct {
	InProgress bool      `json:"in_progress"`
	LastStep   string    `json:"last_step"`
	StartedAt  time.Time `json:"started_at"`
	EndedAt    time.Time `json:"ended_at"`
	Error      string    `json:"error,omitempty"`
}

type Customer struct {
	ID            int       `json:"id"`
	Name          string    `json:"name"`
	WalletAddress string    `json:"wallet_address"`
	Contract      string    `json:"contract"`
	Hashrate      string    `json:"hashrate"`
	StartDate     time.Time `json:"start_date"`
	EndDate       time.Time `json:"end_date"`
	Email         string    `json:"email"`
}
