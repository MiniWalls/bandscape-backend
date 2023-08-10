package mydb

import (
	"testing"
)

func TestDbConnection(t *testing.T) {
	DbConnection()
	result := GetDB()
	if result == nil {
		t.Errorf("Expected db object, got nil")
	} else {
		t.Logf("Expected db object, got %v", result)
	}
}

func TestGetPost(t *testing.T) {
	result, err := GetPost("1")
	if err != nil {
		t.Errorf("Expected nil error, got %v", err)
	} else {
		t.Logf("Expected first post, got %v", result)
	}
}

func TestGetAllPosts(t *testing.T) {
	result, err := GetPosts()
	if err != nil {
		t.Errorf("Expected nil error, got %v", err)
	} else {
		if len(result) == 0 {
			t.Errorf("Expected all posts, got empty array %v", result)
		} else {
			t.Logf("Expected all posts, got %v", result)
		}
	}
}
