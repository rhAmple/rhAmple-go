package main

import (
	"fmt"
	"log"

	"github.com/rhAmple/rhAmple-go"
)

func main() {
	rhAmple := rhAmple.GetRhAmple()

	signal, err := rhAmple.StrategySignal()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Signal: %s\n", signal.String())
}
