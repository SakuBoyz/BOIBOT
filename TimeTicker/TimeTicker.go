package main

import (
	"fmt"
	"time"
	"github.com/Ohmsmmm/boibot"
)
func statusUpdate() string { return "Running" }

func main() {
	c := time.Tick(5 * time.Second)
	for next := range c {
		fmt.Printf("%v %s\n", next, UpdateTotalThailandCovid)
	}
}