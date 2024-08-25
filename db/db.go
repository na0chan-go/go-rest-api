package db

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDB() (*gorm.DB, error) {
	if err := loadEnv(); err != nil {
		return nil, fmt.Errorf("環境変数の読み込みに失敗しました: %w", err)
	}

	url := buildConnectionURL()
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("データベース接続に失敗しました: %w", err)
	}

	log.Println("データベースに接続しました")
	return db, nil
}

func CloseDB(db *gorm.DB) error {
	sqlDB, err := db.DB()
	if err != nil {
		return fmt.Errorf("SQLデータベースの取得に失敗しました: %w", err)
	}
	if err := sqlDB.Close(); err != nil {
		return fmt.Errorf("データベース接続のクローズに失敗しました: %w", err)
	}
	return nil
}

func loadEnv() error {
	if os.Getenv("GO_ENV") == "dev" {
		if err := godotenv.Load(); err != nil {
			return fmt.Errorf(".envファイルの読み込みに失敗しました: %w", err)
		}
	}
	return nil
}

func buildConnectionURL() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_DB"),
	)
}
