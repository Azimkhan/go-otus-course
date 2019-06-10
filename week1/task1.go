package main

import (
	"github.com/beevik/ntp"
	"log"
	"fmt"
)

func main() {
	t,err := ntp.Time("time.google.com")

	if err != nil {
		log.Println("Error gettin time...")
		return
	}

	fmt.Println("Time: ", t)
}


