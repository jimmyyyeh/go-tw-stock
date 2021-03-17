package main

import (
	"fmt"
	"twStockPlugin/utils"
)

func main () {
	//testing date utils
	dateUtils := utils.DateUtils{}
	fmt.Println(dateUtils.AcToRepublicEra("20200318"))
	fmt.Println(dateUtils.AcToRepublicEra("2020/03/18"))
	fmt.Println(dateUtils.AcToRepublicEra("2020-03-18"))
	fmt.Println(dateUtils.RepublicEraToAc("1100318"))
	fmt.Println(dateUtils.RepublicEraToAc("110/03/18"))
	fmt.Println(dateUtils.RepublicEraToAc("110-03-18"))
}
