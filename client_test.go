package br

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

type MockURL struct{}

func (m *MockURL) build() string {
	dir, _ := os.Getwd()
	return "file:" + filepath.Join(dir, "/test/bankiru")
}

type MockInvalidURL struct{}

func (m *MockInvalidURL) build() string {
	dir, _ := os.Getwd()
	return "file:" + filepath.Join(dir, "/test/invalid-bankiru")
}

func TestClient_Rates(t *testing.T) {
	c := NewClient()
	c.url = &MockURL{}

	r, err := c.Rates("", "")
	if err != nil {
		t.Error(err)
	}

	b := r.Branches
	if len(b) == 0 {
		t.Error("b is empty")
	}

	if bCount := len(b); bCount != 5 {
		t.Errorf("bCount got = %v, want %v", bCount, 5)
	}
}

func TestClient_RatesError(t *testing.T) {
	Debug = true

	c := NewClient()
	c.url = &MockInvalidURL{}
	if _, err := c.Rates(CNY, Sochi); err == nil {
		t.Error(err)
	}
}

func TestURL_build(t *testing.T) {
	want := fmt.Sprintf(baseURL, strings.ToLower(string(Crnc)), Ct)
	if got := (&URL{}).build(); got != want {
		t.Errorf("URL.build() = %v, want %v", got, want)
	}
}
