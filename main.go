package main

import (
	"flag"

	"GoAcademy/handlers"
)

func main() {

beverages := flag.String("b", "wines", "Request list of wines")

flag.Parse()
	if *beverages == "wines" {
		handlers.GetWines() 
		} else if *beverages == "coffee"{
		handlers.GetCoffees()
	}
}