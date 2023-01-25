package br

import (
	"os"
	"path/filepath"
	"testing"
)

func TestClient_Rates(t *testing.T) {
	dir, _ := os.Getwd()

	c := NewClient()
	c.url = "file:" + filepath.Join(dir, "/test/bankiru")
	r, err := c.Rates("", "")
	if err != nil {
		t.Error(err)
	}

	b := r.Branches
	if len(b) == 0 {
		t.Error("b is empty")
	}

	bCount := len(b)
	if bCount != 5 {
		t.Errorf("bCount got = %v, want %v", bCount, 5)
	}
}

func TestClient_RatesError(t *testing.T) {
	dir, _ := os.Getwd()

	c := NewClient()
	c.url = "file:" + filepath.Join(dir, "/test/invalid-bankiru")
	_, err := c.Rates("", "")
	if err == nil {
		t.Error(err)
	}
}
