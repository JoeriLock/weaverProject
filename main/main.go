package main

import (
	"fmt"
	"joeri/weave/energy"
	"joeri/weave/meter"
)

func main() {
	test := energy.EnergyToCost(100, true, true, true)
	fmt.Println(test)
	val := meter.GetRow()
	fmt.Println(val)

}
