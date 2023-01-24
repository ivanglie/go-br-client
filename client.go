package br

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gocolly/colly/v2"
	"github.com/gocolly/colly/v2/extensions"
)

const (
	// Example: https://www.banki.ru/products/currency/cash/usd/moskva/
	baseURL = "https://www.banki.ru/products/currency/cash/%s/%s/"

	// Currency
	USD Currency = "USD"
	EUR Currency = "EUR"
	GBP Currency = "GBP"
	CHF Currency = "CHF"
	JPY Currency = "JPY"
	CNY Currency = "CNY"
	CZK Currency = "CZK"
	PLN Currency = "PLN"

	// Region
	Barnaul         Region = "barnaul"
	Voronezh        Region = "voronezh"
	Volgograd       Region = "volgograd"
	Vladivostok     Region = "vladivostok"
	Ekaterinburg    Region = "ekaterinburg"
	Irkutsk         Region = "irkutsk"
	Izhevsk         Region = "izhevsk"
	Kazan           Region = "kazan~"
	Krasnodar       Region = "krasnodar"
	Krasnoyarsk     Region = "krasnoyarsk"
	Kaliningrad     Region = "kaliningrad"
	Kirov           Region = "kirov"
	Kemerovo        Region = "kemerovo"
	Moscow          Region = "moskva"
	Novosibirsk     Region = "novosibirsk"
	NizhnyNovgorod  Region = "nizhniy_novgorod"
	Omsk            Region = "omsk"
	Orenburg        Region = "orenburg"
	Perm            Region = "perm~"
	RostovOnDon     Region = "rostov-na-donu"
	SaintPetersburg Region = "sankt-peterburg"
	Samara          Region = "samara"
	Saratov         Region = "saratov"
	Sochi           Region = "krasnodarskiy_kray/sochi"
	Tyumen          Region = "tyumen~"
	Tolyatti        Region = "samarskaya_oblast~/tol~yatti"
	Tomsk           Region = "tomsk"
	Ufa             Region = "ufa"
	Khabarovsk      Region = "habarovsk"
	Chelyabinsk     Region = "chelyabinsk"
)

type Client struct {
	collector *colly.Collector
}

func NewClient() *Client {
	c := colly.NewCollector()

	t := &http.Transport{}
	t.RegisterProtocol("file", http.NewFileTransport(http.Dir("/")))

	c.WithTransport(t)
	c.AllowURLRevisit = true

	extensions.RandomUserAgent(c)

	return &Client{c}
}

var Debug bool

// Get cash currency exchange rates by currency (USD, if empty) and region (Moscow, if empty).
func (s *Client) GetRates(c Currency, r Region) (*Rates, error) {
	if len(c) == 0 {
		c = USD
	}

	if len(r) == 0 {
		r = Moscow
	}

	url := fmt.Sprintf(baseURL, strings.ToLower(string(c)), r)
	if Debug {
		log.Printf("Fetching the currency rate from %s", url)
	}

	rates := &Rates{Currency: c, Region: r}
	branches, err := s.parseBranches(url)
	if err != nil {
		rates = nil
	} else {
		rates.Branches = branches
	}

	return rates, err
}

// Parse banks and their branches info.
func (s *Client) parseBranches(url string) ([]Branch, error) {
	var branches []Branch
	var err error

	s.collector.OnHTML("div.table-flex.trades-table.table-product", func(e *colly.HTMLElement) {
		e.ForEach("div.table-flex__row.item.calculator-hover-icon__container", func(i int, row *colly.HTMLElement) {
			bank := row.ChildText("a.font-bold")

			a := strings.Split(row.ChildAttr("a.font-bold", "data-currency-rates-tab-item"), "_")
			address := a[len(a)-1]

			subway := row.ChildText("div.font-size-small")
			currency := row.ChildAttr("div.table-flex__rate.font-size-large", "data-currencies-code")

			var buy float64
			buy, err = strconv.ParseFloat(row.ChildAttr("div.table-flex__rate.font-size-large", "data-currencies-rate-buy"), 64)
			if err != nil {
				log.Println(err)
				return
			}

			var sell float64
			sell, err = strconv.ParseFloat(row.ChildAttr("div.table-flex__rate.font-size-large.text-nowrap", "data-currencies-rate-sell"), 64)
			if err != nil {
				log.Println(err)
				return
			}

			var location *time.Location
			location, err = time.LoadLocation("Europe/Moscow")
			if err != nil {
				log.Println(err)
				return
			}

			var updated time.Time
			updated, err = time.ParseInLocation("02.01.2006 15:04", row.ChildText("span.text-nowrap"), location)
			if err != nil {
				log.Println(err)
				return
			}

			raw := newBranch(bank, address, subway, currency, buy, sell, updated)
			if raw != (Branch{}) && time.Now().In(location).Sub(updated) < 24*time.Hour && (buy != 0 || sell != 0) {
				branches = append(branches, raw)
			}
		})
	})

	s.collector.OnRequest(func(r *colly.Request) {
		log.Printf("UserAgent: %s", r.Headers.Get("User-Agent"))
	})

	s.collector.OnError(func(r *colly.Response, e error) {
		err = e
		log.Println(err)
	})

	err = s.collector.Visit(url)

	return branches, err
}
