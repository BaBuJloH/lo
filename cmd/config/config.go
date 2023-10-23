package config

import (
	"os"
)

func ConfigSetup() {
	// Database settings
	os.Setenv("DB_USERNAME", "postgres")
	os.Setenv("DB_PASSWORD", "123")
	os.Setenv("DB_HOST", "localhost:5432")
	os.Setenv("DB_NAME", "postgres")

	os.Setenv("DB_POOL_MAXCONN", "5")
	os.Setenv("DB_POOL_MAXCONN_LIFETIME", "300")

	// NATS-Streaming settings
	os.Setenv("NATS_HOSTS", "nats://127.0.0.1:4222")
	os.Setenv("NATS_CLUSTER_ID", "test-cluster")
	os.Setenv("NATS_CLIENT_ID", "client-123")
	os.Setenv("NATS_SUBJECT", "go.nats")
	os.Setenv("NATS_DURABLE_NAME", "my-durable")
	os.Setenv("NATS_ACK_WAIT_SECONDS", "30")

	// Cache settings
	os.Setenv("CACHE_SIZE", "10")
	os.Setenv("APP_KEY", "WB-1")
}
