package main

import (
	"fmt"
	"github.com/beevik/ntp"
	"io"
	"os"
	"time"
)

func main() {
	ntpTime, err := ntp.Time("pool.ntp.org")
	if err != nil {
		fmt.Println(err)
		_, err = io.WriteString(os.Stderr, err.Error())
		if err != nil {
			fmt.Println(err)
		}
		os.Exit(0)
	}

	ntpTimeFormatted := ntpTime.Format(time.UnixDate)

	fmt.Printf("Network time: %v\n", ntpTime)
	fmt.Printf("Unix Date Network time: %v\n", ntpTimeFormatted)
	fmt.Println("+++++++++++++++++++++++++++++++")
	timeFormatted := time.Now().Local().Format(time.UnixDate)
	fmt.Printf("System time: %v\n", time.Now())
	fmt.Printf("Unix Date System time: %v\n", timeFormatted)

}
