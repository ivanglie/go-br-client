package br

import (
	"os"
	"path/filepath"
	"testing"
)

func TestClient_Rates(t *testing.T) {
	dir, _ := os.Getwd()
	absFilePath := filepath.Join(dir, "/test/bankiru")

	c := NewClient()
	c.url = "file:" + absFilePath
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
