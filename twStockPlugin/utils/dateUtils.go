package utils

import (
	"strconv"
	"goTwStock/regex"
)

type DateUtils struct {
}

func (d *DateUtils) RepublicEraToAc(republicEraStr string) string {
	/*
		民國轉西元
	*/
	if !regex.RepublicEra.MatchString(republicEraStr) {
		panic("Republic Era format error")
	}
	dateMap := map[string]string{}
	keyMap := regex.RepublicEra.SubexpNames()[1:]
	valueMap := regex.RepublicEra.FindStringSubmatch(republicEraStr)[1:]

	for i, _ := range keyMap {
		dateMap[keyMap[i]] = valueMap[i]
	}

	if dateMap["symbol1"] != dateMap["symbol2"] {
		panic("Republic Era format error")
	}

	year_, _ := strconv.Atoi(dateMap["year"])
	year_ = year_ + 1911
	dateMap["year"] = strconv.Itoa(year_)

	acStr := ""
	for _, key := range []string{"year", "symbol1", "month", "symbol2", "day"} {
		acStr += dateMap[key]
	}
	return acStr
}

func (d *DateUtils) AcToRepublicEra(acStr string) string {
	/*
		西元轉民國
	*/
	if !regex.AC.MatchString(acStr) {
		panic("AC format error")
	}
	dateMap := map[string]string{}
	keyMap := regex.AC.SubexpNames()[1:]
	valueMap := regex.AC.FindStringSubmatch(acStr)[1:]

	for i, _ := range keyMap {
		dateMap[keyMap[i]] = valueMap[i]
	}

	if dateMap["symbol1"] != dateMap["symbol2"] {
		panic("AC format error")
	}

	year_, _ := strconv.Atoi(dateMap["year"])
	year_ = year_ - 1911
	dateMap["year"] = strconv.Itoa(year_)

	republicEraStr := ""
	for _, key := range []string{"year", "symbol1", "month", "symbol2", "day"} {
		republicEraStr += dateMap[key]
	}
	return republicEraStr
}
