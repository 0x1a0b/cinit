package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	if len(os.Args) != 2 {
		usage()
	}
	sleepTime, err := strconv.Atoi(os.Args[1])
	if err != nil {
		usage()
	}

	i := 0
	fmt.Printf("sleeper %v - starting\n", sleepTime)
	for {
		t := time.NewTimer(time.Duration(sleepTime) * time.Second)
		<-t.C
		i++
		fmt.Printf("sleeper %v - %v\n", sleepTime, i)
	}

}

func usage() {
		fmt.Printf("usage: %v <sleepSeconds>\n", os.Args[0])
		os.Exit(1)
}
