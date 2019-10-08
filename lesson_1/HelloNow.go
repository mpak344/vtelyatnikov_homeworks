package main

import (
	"fmt"
	"github.com/beevik/ntp"
	"os"
	"time"
)

var host = "0.ru.pool.ntp.org"

func main() {

	currentTime := time.Now()
	fmt.Printf("Current time, %s\n", currentTime.Local())

	time, err := ntp.Time(host)

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Printf("Ntp time: %s", time)
}
