package main

import (
	"flag"
	"fmt"
	"github.com/deter-project/go-spi/cli"
	"github.com/deter-project/go-spi/spi"
)

//flags
var opF = flag.String("op", "", "[create | remove | view]")
var rzF = flag.String("rz", "", "SPI Realization Name")
var circF = flag.String("circle", "", "SPI Circle")

func main() {
	flag.Parse()

	switch *opF {
	case "create":
		create()
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

func create() {

	if *rzF == "" {
		cli.Fatal(
			"you must specify a realization name with the -rz flag, " +
				"use the -help flag for details")
	}

	if *circF == "" {
		cli.Fatal(
			"you must specify a circle name with the -circle flag, " +
				"use the -help flag for details")
	}

	cli.PreRun()

	rsp, err := spi.RealizeExperiment(*rzF, *circF, cli.User)
	if err != nil {
		cli.Fatal("create realization failed")
	}

	fmt.Println(rsp)

}

func remove() {

	if *rzF == "" {
		cli.Fatal(
			"you must specify a realization name with the -rz flag, " +
				"use the -help flag for details")
	}

	cli.PreRun()

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
		fmt.Printf("%s : %s\n", x.Name, x.Status)
	}

}
