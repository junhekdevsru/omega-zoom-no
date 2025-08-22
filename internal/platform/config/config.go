package config

import "os"

type Config struct {
	GRPCAddr   string
	HTTPAddr   string // gateway
	PGConn     string
	Migrations string
}

func MustLoad() Config {
	return Config{
		GRPCAddr:   getenv("GRPC_ADDR", ":8080"),
		HTTPAddr:   getenv("HTTP_ADDR", ":8082"),
		PGConn:     getenv("PG_CONN", "postgres://agent:agent@localhost:5432/agentdb?sslmode=disable"),
		Migrations: getenv("MIGRATIONS_DIR", "migrations"),
	}
}

func getenv(k, def string) string {
	if v := os.Getenv(k); v != "" {
		return v
	}
	return def
}
