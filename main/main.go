package main

import (
	"encoding/csv"
	"fmt"
	"joeri/weave/energy"
	"joeri/weave/meter"
	"os"
)

var FILEPATH string = "../result.csv"

func main() {
	//get Meter data
	val := meter.GetData()

	//declare variabels
	output := make(map[int64]float64)
	var cost float64 //cost of the energy
	var oldCost float64
	var inc float64 //increment

	for k, row := range val {

		if k != 0 { // check if k is no first element, to preven null pointer
			if val[k-1].Id != row.Id { //new ID so there is no incremetn or oldCost
				inc = 0
				oldCost = 0
			} else {
				//diff between this and last reading is bigger then 100 or below 0
				if (row.Reading-val[k-1].Reading) > 100 || row.Reading-val[k-1].Reading < 0 {
					output[row.Id] += inc //We have a wrong value so we assume they cosumed same as last increment
					continue
				}
			}
		}

		cost = energy.EnergyToCost(row.Reading, row.Tp, row.Created_at) // what did it cost
		inc = cost - oldCost
		output[row.Id] += inc // add increment to total
		oldCost = cost
	}
	toCsv(output)
}

//save to csv
func toCsv(data map[int64]float64) {
	file, _ := os.Create(FILEPATH)
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for k, value := range data {
		writer.Write([]string{fmt.Sprint(k), fmt.Sprintf("%f", value)})
	}
}
