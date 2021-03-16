package utils

import (
	"fmt"
	"strconv"
	"time"
	"twStockPlugin/regex"
)

type DateUtils struct {
}

func (d *DateUtils) RepublicEraToAc(republicEraStr string) string {
	//民國轉西元
	if !regex.RepublicEra.MatchString(republicEraStr) {
		panic("Republic Era format error")
	}
	y_, m_, d_ := "", "", ""
	if len(republicEraStr) == 7 {
		y_, m_, d_ = republicEraStr[:3], republicEraStr[3:5], republicEraStr[5:]
	} else {
		y_, m_, d_ = republicEraStr[:2], republicEraStr[2:4], republicEraStr[4:]
	}
	y__, _ := strconv.Atoi(y_)
	y__ += 1911
	acStr := fmt.Sprintf("%d%s%s", y__, m_, d_)
	return acStr
}

func (d *DateUtils) AcToRepublicEra(acStr string) string {
	//西元轉民國
	if !regex.AC.MatchString(acStr) {
		panic("AC format error")
	}
	republicEra, _ := time.Parse("20060102", acStr)
	y_ := strconv.Itoa(republicEra.Year() - 1911)
	m_ := fmt.Sprintf("%02d", int(republicEra.Month()))
	d_ := strconv.Itoa(republicEra.Day())
	republicEraStr := y_ + m_ + d_
	return republicEraStr
}
