package main

import (
	"fmt"
	"github.com/beevik/ntp"
	"log"
)

func main() {
	t, err := ntp.Time("time.google.com")

	if err != nil {
		log.Println("Error gettin time...")
		return
	}

	fmt.Println("Time: ", t)
}
