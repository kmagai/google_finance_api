package main

import (
	"fmt"

	"github.com/kmagai/googleFinance"
)

func main() {
	api := googleFinance.API{}
	code := "NI225" // Nikkei 225
	stock, err := api.GetStock(code)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(stock)
	return
}
