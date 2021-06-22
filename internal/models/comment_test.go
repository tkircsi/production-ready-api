package models

import (
	"fmt"
	"testing"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestPostComment(t *testing.T) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", "localhost", "postgres", "example", "postgres", "5432")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Fatalf("Connecting error: %v", err)
	}

	sqlDb, err := db.DB()
	if err != nil {
		t.Fatalf("Connecting error: %v", err)
	}

	if err := sqlDb.Ping(); err != nil {
		t.Fatalf("Ping error: %v", err)
	}

	db.AutoMigrate(&Comment{})

	s := NewService(db)

	c := Comment{
		Slug:   "Slug",
		Body:   "This is a post",
		Author: "Author",
	}
	res, err := s.PostComment(c)
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	t.Logf("res: %v", res)
}
