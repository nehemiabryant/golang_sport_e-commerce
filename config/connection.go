package config

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

var DB *pgxpool.Pool

func ConnectDB() error {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Warning: .env tidak ditemukan, lanjut...")
	}

	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		return fmt.Errorf("DATABASE_URL tidak ditemukan di .env")
	}

	// Parse config pool
	cfg, err := pgxpool.ParseConfig(databaseURL)
	if err != nil {
		return fmt.Errorf("gagal parsing config: %v", err)
	}

	// Matikan prepared-statement cache dalam supabase
	cfg.ConnConfig.DefaultQueryExecMode = pgx.QueryExecModeSimpleProtocol

	// Optional: Pool tuning
	cfg.MaxConns = 10
	cfg.MinConns = 2
	cfg.MaxConnIdleTime = 3 * time.Minute

	// Buat DB pool
	ctx := context.Background()
	DB, err = pgxpool.NewWithConfig(ctx, cfg)
	if err != nil {
		return fmt.Errorf("gagal membuat koneksi ke DB: %v", err)
	}

	// Test ping
	err = DB.Ping(ctx)
	if err != nil {
		return fmt.Errorf("gagal ping DB: %v", err)
	}

	fmt.Println("Berhasil terhubung ke Supabase PostgreSQL!")
	return nil
}
