package main

import (
	"fmt"
	"github.com/beevik/ntp"
)

func main() {
	t,err := ntp.Time("time.google.com")

	if err != nil {
		fmt.Println("Error gettin time...")
		return
	}

	fmt.Println("Time: ", t)
}


