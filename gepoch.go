// Package main reads the command line flags and arguments
package main

import (
	"flag"
	"fmt"
	"gepoch/epoch"
	"strconv"
)

func panicHandler() {
	if r := recover(); r != nil {
		fmt.Printf("\nFatal error: %s\n", r)
	}
}

func main() {

	defer panicHandler()

	var format string
	flag.StringVar(&format, "format", "plain", "output format [json/plain]")
	flag.StringVar(&format, "f", "plain", "output format [json/plain] (shorthand)")

	flag.Parse()
	args := flag.Args()

	// parse args to int64
	timestamps := []int64{}
	for _, ts := range args {
		i, err := strconv.ParseInt(ts, 10, 64)
		if err != nil {
			panic(fmt.Sprintf(`invalid "%s" timestamp`, ts))
		}
		timestamps = append(timestamps, i)
	}

	en := epoch.ConvertN(timestamps)
	en.Print(format)
}
