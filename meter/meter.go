package meter

import (
	"encoding/csv"
	"log"
	"os"
	"sort"
	"strconv"
)

//Type that contains all data of a row from cv
type DataRow struct {
	Id         int64
	Tp         bool
	Reading    float64
	Created_at int64
}

//csv to read from
var FILEPATH string = "../data.csv"

func GetData() []DataRow {
	rtn := readCsvFile()
	return parse(rtn)
}

//extracts data from csv and puts them in type
func parse(data [][]string) []DataRow {
	var dataRows []DataRow
	for _, row := range data {
		id, _ := strconv.ParseInt(row[0], 10, 32)
		reading, _ := strconv.ParseFloat(row[2], 64)
		isGass, _ := strconv.ParseInt(row[1], 10, 32)
		timeSt, _ := strconv.ParseInt(row[3], 10, 64)
		//appending seems slow, can probaly be done beter
		dataRows = append(dataRows, DataRow{id, isGass == 2, reading, timeSt})
	}

	return sortRows(dataRows)

}

//sort rows by id then by timestapm
//
func sortRows(dataRows []DataRow) []DataRow {
	sort.SliceStable(dataRows, func(i, j int) bool {
		if dataRows[i].Id < dataRows[j].Id {
			return true
		}
		if dataRows[i].Id > dataRows[j].Id {
			return false
		}
		return dataRows[i].Created_at < dataRows[j].Created_at
	})
	return dataRows
}

//read the csv file
func readCsvFile() [][]string {
	f, err := os.Open(FILEPATH)
	if err != nil {
		log.Fatal("Unable to read input file "+FILEPATH, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+FILEPATH, err)
	}
	return records
}
