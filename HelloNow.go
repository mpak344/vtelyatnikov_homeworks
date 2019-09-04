package main

import (
	"fmt"
	"time"
	"os"
	"github.com/beevik/ntp"
)
var host = "0.ru.pool.ntp.org"

func main() {

	currentTime := time.Now()
	fmt.Printf("Current time, %s\n", currentTime.Local())

	if time, err := ntp.Time(host); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	} else {
		fmt.Printf("Ntp time: %s", time)
	}
}
