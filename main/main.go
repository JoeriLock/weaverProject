package main

import (
	"fmt"
	"joeri/weave/meter"
)

func main() {
	// test := energy.EnergyToCost(100, true, true, true)
	// fmt.Println(test)
	// val := meter.GetRow()
	// fmt.Println(val)
	// fmt.Println(meterdata.GetId())
	// a := []int{1, 2, 3}
	// fmt.Println(a[1:])
	val := meter.GetData()
	for _, row := range val {
		fmt.Println(row)
	}
}
