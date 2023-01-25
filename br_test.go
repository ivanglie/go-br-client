package br

import (
	"reflect"
	"sort"
	"testing"
	"time"
)

func Test_newBranch(t *testing.T) {
	got := newBranch("b", "a", "s", "c", 12.34, 56.78, func() time.Time { t, _ := time.Parse("02.01.2006 15:04", "01.02.2018 12:35"); return t }())
	want := newBranch("b", "a", "s", "c", 12.34, 56.78, func() time.Time { t, _ := time.Parse("02.01.2006 15:04", "01.02.2018 12:35"); return t }())
	if !reflect.DeepEqual(got, want) {
		t.Errorf("newBranch() = %v, want %v", got, want)
	}
}

func TestBranch_ByBuySorter(t *testing.T) {
	got := []Branch{
		newBranch("bank", "address", "subway", "currency", 101.34, 58.78, func() time.Time { t, _ := time.Parse("02.01.2006 15:04", "01.02.2018 12:35"); return t }()),
		newBranch("bank", "address", "subway", "currency", 100.23, 58.78, func() time.Time { t, _ := time.Parse("02.01.2006 15:04", "01.02.2018 12:35"); return t }()),
		newBranch("bank", "address", "subway", "currency", 56.78, 58.78, func() time.Time { t, _ := time.Parse("02.01.2006 15:04", "01.02.2018 12:35"); return t }()),
		newBranch("bank", "address", "subway", "currency", 90.12, 58.78, func() time.Time { t, _ := time.Parse("02.01.2006 15:04", "01.02.2018 12:35"); return t }()),
		newBranch("bank", "address", "subway", "currency", 12.34, 58.78, func() time.Time { t, _ := time.Parse("02.01.2006 15:04", "01.02.2018 12:35"); return t }()),
	}
	sort.Sort(ByBuySorter(got))

	want := []Branch{
		newBranch("bank", "address", "subway", "currency", 12.34, 58.78, func() time.Time { t, _ := time.Parse("02.01.2006 15:04", "01.02.2018 12:35"); return t }()),
		newBranch("bank", "address", "subway", "currency", 56.78, 58.78, func() time.Time { t, _ := time.Parse("02.01.2006 15:04", "01.02.2018 12:35"); return t }()),
		newBranch("bank", "address", "subway", "currency", 90.12, 58.78, func() time.Time { t, _ := time.Parse("02.01.2006 15:04", "01.02.2018 12:35"); return t }()),
		newBranch("bank", "address", "subway", "currency", 100.23, 58.78, func() time.Time { t, _ := time.Parse("02.01.2006 15:04", "01.02.2018 12:35"); return t }()),
		newBranch("bank", "address", "subway", "currency", 101.34, 58.78, func() time.Time { t, _ := time.Parse("02.01.2006 15:04", "01.02.2018 12:35"); return t }()),
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("branch = %v, want %v", got, want)
	}
}

func TestBranch_BySellSorter(t *testing.T) {
	got := []Branch{
		newBranch("bank", "address", "subway", "currency", 12.34, 58.78, func() time.Time { t, _ := time.Parse("02.01.2006 15:04", "01.02.2018 12:35"); return t }()),
		newBranch("bank", "address", "subway", "currency", 12.34, 56.75, func() time.Time { t, _ := time.Parse("02.01.2006 15:04", "01.02.2018 12:35"); return t }()),
		newBranch("bank", "address", "subway", "currency", 12.34, 78.56, func() time.Time { t, _ := time.Parse("02.01.2006 15:04", "01.02.2018 12:35"); return t }()),
		newBranch("bank", "address", "subway", "currency", 12.34, 56.78, func() time.Time { t, _ := time.Parse("02.01.2006 15:04", "01.02.2018 12:35"); return t }()),
		newBranch("bank", "address", "subway", "currency", 12.34, 52.64, func() time.Time { t, _ := time.Parse("02.01.2006 15:04", "01.02.2018 12:35"); return t }()),
	}
	sort.Sort(BySellSorter(got))

	want := []Branch{
		newBranch("bank", "address", "subway", "currency", 12.34, 52.64, func() time.Time { t, _ := time.Parse("02.01.2006 15:04", "01.02.2018 12:35"); return t }()),
		newBranch("bank", "address", "subway", "currency", 12.34, 56.75, func() time.Time { t, _ := time.Parse("02.01.2006 15:04", "01.02.2018 12:35"); return t }()),
		newBranch("bank", "address", "subway", "currency", 12.34, 56.78, func() time.Time { t, _ := time.Parse("02.01.2006 15:04", "01.02.2018 12:35"); return t }()),
		newBranch("bank", "address", "subway", "currency", 12.34, 58.78, func() time.Time { t, _ := time.Parse("02.01.2006 15:04", "01.02.2018 12:35"); return t }()),
		newBranch("bank", "address", "subway", "currency", 12.34, 78.56, func() time.Time { t, _ := time.Parse("02.01.2006 15:04", "01.02.2018 12:35"); return t }()),
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("branch = %v, want %v", got, want)
	}
}

func TestRates_String(t *testing.T) {
	loc, _ := time.LoadLocation("Europe/Moscow")
	b := newBranch(
		"Банк «Открытие»",
		"630102, г. Новосибирск, ул. Кирова, дом. 44",
		"м. Октябрьская", "CNY",
		9.61,
		11.64,
		time.Date(2023, time.January, 24, 16, 54, 0, 0, loc))

	r := &Rates{}
	r.Currency = CNY
	r.City = Novosibirsk
	r.Branches = []Branch{b}

	got := r.String()
	want := `{"currency":"CNY","city":"novosibirsk","branches":[{` +
		`"bank":"Банк «Открытие»","address":"630102, г. Новосибирск, ул. Кирова, дом. 44","subway":"м. Октябрьская",` +
		`"currency":"CNY","buy":9.61,"sell":11.64,"updated":"2023-01-24T16:54:00+03:00"}]}`

	if got != want {
		t.Errorf("got= %v, want= %v", got, want)
	}
}
