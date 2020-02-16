package main

import (
	"api/providers/locations_provider"
	"fmt"
)

func main() {
	country, err := locations_provider.GetCountry("AR")

	fmt.Println(err)
	fmt.Println(country)
}
