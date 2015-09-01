package main

import (
	"flag"
	"fmt"
	"github.com/deter-project/go-spi/cli"
	"github.com/deter-project/go-spi/spi"
)

//flags
var opF = flag.String("op", "", "[remove | view]")
var rzF = flag.String("rz", "", "SPI Realization Name")

func main() {
	flag.Parse()

	switch *opF {
	case "remove":
		remove()
	case "view":
		view()
	case "":
		cli.Fatal(
			"you must specify an operation with the -op flag, " +
				"use the -help flag for details")
	default:
		cli.Fatal(
			"unknown operation " + *opF + ", " +
				"use the -help flag for details")

	}

}

func remove() {

	cli.PreRun()
	if *rzF == "" {
		cli.Fatal(
			"you must specify a realization name with the -rz flag, " +
				"use the -help flag for details")
	}

	rsp, err := spi.RemoveRealization(*rzF)
	if err != nil {
		cli.Fatal(fmt.Sprintf("remove realization failed: %v", err))
	}

	if rsp.Return != true {
		cli.Fatal("Removal of the realization failed")
	}

}

func view() {

	cli.PreRun()

	rsp, err := spi.ViewRealizations(cli.User, ".*")
	if err != nil {
		cli.Fatal(fmt.Sprintf("view realizations failed: %v", err))
	}
	for _, x := range rsp.Return {
		fmt.Println(x.Name)
	}

}
