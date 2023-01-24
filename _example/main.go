package main

import (
	"fmt"

	br "github.com/ivanglie/go-br-client"
)

func main() {
	client := br.NewClient()
	rates, err := client.Rates(br.CNY, br.Novosibirsk)
	if err != nil {
		panic(err)
	}
	fmt.Println(rates)
}
