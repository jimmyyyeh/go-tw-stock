package utils

import (
	"encoding/csv"
	"net/http"
)

var utils = Utils{}

type CsvUtils struct {
}

func (c *CsvUtils) ReadCsv(url string, comma rune) interface{} {
	/*
		read url from csv, split by comma
	*/
	resp, err := http.Get(url)
	utils.ErrorHandler(err)
	defer resp.Body.Close()
	reader := csv.NewReader(resp.Body)
	reader.Comma = comma
	data, err := reader.ReadAll()
	utils.ErrorHandler(err)
	return data
}
