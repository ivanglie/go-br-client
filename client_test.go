package br

import (
	"os"
	"path/filepath"
	"testing"
)

func Test_parseBranches(t *testing.T) {
	dir, _ := os.Getwd()
	absFilePath := filepath.Join(dir, "/test/bankiru")

	b, _ := NewClient().parseBranches("file:" + absFilePath)

	if len(b) == 0 {
		t.Errorf("currency.branches is empty")
	}

	branchesCount := len(b)

	if branchesCount != 5 {
		t.Errorf("branchesCount got = %v, want %v", branchesCount, 5)
	}
}
