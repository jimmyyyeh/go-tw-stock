package stockInfo

import (
	"fmt"
	"strconv"
	"time"
	"twStockPlugin/constant"
	"twStockPlugin/regex"
	"twStockPlugin/schema"
	"twStockPlugin/utils"
)

var csvMap = map[string]string{
	"上市": "",
	"上櫃": "t187ap03_O.csv",
}

var categoryConstMap = map[string]map[string]int{
	"上市": constant.ListedCategory,
	"上櫃": constant.OTCCategory,
}

func getReverseCategoryMap(originMap map[string]int) map[int]string {
	/*
		reversed category map from key: value -> value: key
	*/
	var reversedMap = map[int]string{}
	for k, v := range originMap {
		reversedMap[v] = k
	}
	return reversedMap
}

func Fetch() map[string]schema.StockInfo {
	csvUtils := utils.CsvUtils{}
	results := map[string]schema.StockInfo{}

	for type_, csv := range csvMap {
		reversedCategoryMap := getReverseCategoryMap(categoryConstMap[type_])
		url := fmt.Sprintf("%s/opendata/%s", constant.DtsTwseDomain, csv)
		dataList := csvUtils.ReadCsv(url, ',').([][]string)
		for _, data := range dataList[1:] {
			incorporationDay, _ := time.Parse("20060102", data[14])
			listingDay, _ := time.Parse("20060102", data[15])
			value, _ := strconv.ParseFloat(regex.Value.FindString(data[16]), 64)
			capital, _ := strconv.ParseFloat(data[17], 64)
			privateStockShares, _ := strconv.ParseFloat(data[18], 64)
			preferredShares, _ := strconv.ParseFloat(data[19], 64)
			categoryCode, _ := strconv.Atoi(data[5])
			category := reversedCategoryMap[categoryCode]
			stockInfo := &schema.StockInfo{
				Code:               data[1],
				Name:               data[2],
				ShortName:          data[3],
				Category:           category,
				Type:               type_,
				Address:            data[6],
				UnifiedNo:          data[7],
				IncorporationDay:   incorporationDay,
				ListingDate:        listingDay,
				Value:              value,
				Capital:            capital,
				PrivateStockShares: privateStockShares,
				PreferredShares:    preferredShares,
			}
			results[stockInfo.Code] = *stockInfo
		}
	}
	return results
}
