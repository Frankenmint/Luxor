# KVM Manager

A modular Go-based orchestrator for provisioning KVM resources, managing cloud mining contracts, and integrating with cloud providers like AWS, GCP, Azure, and Terraform.

---

## 🚀 Getting Started

### Prerequisites
- Go 1.21+
- SQLite3
- Libvirt (for KVM)
- AWS/GCP/Azure CLI or Terraform installed (optional)

### Folder Structure
```
.
├── main.go                  # App entrypoint
├── api/
│   └── provision.go         # HTTP handlers
├── db/
│   └── db.go                # SQLite logic
├── internal/
│   ├── orchestrator.go      # Provisioning workflow
│   └── config.go            # Config loader
├── types/
│   └── types.go             # Shared types (JobStatus, Config, Customer)
├── config.json              # Your configuration
└── kvm_manager.db           # Auto-created SQLite DB
```

### Installation
```bash
git clone https://github.com/yourname/kvm-manager
cd kvm-manager
go mod tidy
go build -o kvm-manager
```

### Configuration
Create a `config.json` file in root:
```json
{
  "region": "us-east-1",
  "s3_bucket_name": "cloud-hashrate",
  "sales_team_net": "192.168.100.0/24",
  "email_domain": "hashcloud.io",
  "auth_token": "supersecrettoken"
}
```
Or set path:
```bash
export CONFIG_PATH=/path/to/config.json
```

---

## 🛠 Running
```bash
./kvm-manager
```
This launches an HTTP server on `:8080` with:

- `POST /provision` — Starts the provisioning workflow
- `GET /status` — Returns current job status
- `GET /history` — Returns recent provisioning logs

Use curl or Postman:
```bash
curl -X POST http://localhost:8080/provision \
  -H "Authorization: Bearer supersecrettoken"
```

---

## 🔌 Extending

### Add Terraform or Cloud SDKs
Replace stub logic in `internal/orchestrator.go`:
```go
func provisionCloudHashrate() error {
    // Example: Run Terraform or GCP SDK logic
    return nil
}
```

### Add More Providers
Split logic into interfaces like:
```go
type CloudProvider interface {
    Provision(customer Customer) error
}
```
Then register your providers and choose dynamically.

### Add Notifications
Extend `syncWithInfra()` to send emails, webhooks, or Slack updates.

---

## 🧪 Roadmap Ideas
- API auth per-user/token
- Web dashboard
- Kubernetes operator integration
- Advanced billing sync

---

## 📄 License
MIT or whatever you choose

---

Made with ❤️ for hybrid compute orchestration.

