package br

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestClient_Rates(t *testing.T) {
	c := NewClient()
	c.buildURL = func() string {
		dir, _ := os.Getwd()
		return "file:" + filepath.Join(dir, "/test/bankiru")
	}

	r, err := c.Rates("")
	if err != nil {
		t.Error(err)
	}

	b := r.Branches
	if len(b) == 0 {
		t.Error("b is empty")
	}

	if bCount := len(b); bCount != 4 {
		t.Errorf("bCount got = %v, want %v", bCount, 4)
	}
}

func TestClient_RatesError(t *testing.T) {
	Debug = true

	c := NewClient()
	c.buildURL = func() string {
		dir, _ := os.Getwd()
		return "file:" + filepath.Join(dir, "/test/invalid-bankiru")
	}

	if _, err := c.Rates(Sochi); err == nil {
		t.Error(err)
	}
}

func TestURL_build(t *testing.T) {
	buildURL := func() string {
		return fmt.Sprintf(baseURL, strings.ToLower(string(Moscow)))
	}

	want := (NewClient()).buildURL()

	if got := buildURL(); got != want {
		t.Errorf("URL.build() = %v, want %v", got, want)
	}
}
