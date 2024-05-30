package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"

	"GoAcademy/handlers"
)

func main() {
	beverages := flag.String("b", "both", "Request list of wines")
	flag.Parse()

	httpClient := &http.Client{Timeout: 10 * time.Second}
	beverageService := handlers.NewBeverageService(httpClient)

	switch *beverages {
	case "wines":
		wines := beverageService.GetWines()
		fmt.Println(wines) // Do something with the wines
	case "coffee":
		coffees := beverageService.GetCoffees()
		fmt.Println(coffees) // Do something with the coffees
	case "both":
		beverageService.GetBoth()
	}
}
