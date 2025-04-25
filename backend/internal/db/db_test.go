package db

import (
	"testing"
	"time"
)

func TestInitAndCloseDB(t *testing.T) {
	err := InitDB("mongodb://localhost:27017")
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	client := GetMongoClient()
	if client == nil {
		t.Errorf("expected MongoClient to be non-nil, got nil")
	}

	CloseDB()
	time.Sleep(1 * time.Second)
}
