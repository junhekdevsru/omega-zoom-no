package config

import "os"

type Config struct {
	GRPCAddr   string
	PGConn     string
	Migrations string
}

func MustLoad() Config {
	cfg := Config{
		GRPCAddr:   getEnv("GRPC_ADDR", ":8080"),
		PGConn:     getEnv("PG_CONN", "postgres://agent:agent@localhost:5432/agentdb?sslmode=disable"),
		Migrations: getEnv("MIGRATIONS_DIR", "migrations"),
	}
	return cfg
}

func getEnv(k, def string) string {
	if v := os.Getenv(k); v != "" {
		return v
	}
	return def
}
