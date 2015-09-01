package main

import (
	"flag"
	"fmt"
	"github.com/deter-project/go-spi/cli"
	"github.com/deter-project/go-spi/spi"
)

//flags
var opF = flag.String("op", "", "[remove | view | change-profile | change-acl]")
var xpF = flag.String("xp", "", "SPI Experiment Name")
var nameF = flag.String("name", "", "Argument Name")
var valueF = flag.String("value", "", "Argument Value")
var deleteF = flag.Bool("delete", false, "Delete Flag")
var circF = flag.String("circle", "", "Circle Flag")
var permF = flag.String("perm", "", "Permission Flag")

func remove() {

	cli.PreRun()

	if *xpF == "" {
		cli.Fatal(
			"you must specify an experiment name with the -xp flag, " +
				"use the -help flag for details")
	}

	rsp, err := spi.RemoveExperiment(*xpF)
	if err != nil {
		cli.Fatal(fmt.Sprintf("remove experiment failed: %v", err))
	}

	if rsp.Return != true {
		cli.Fatal("Removal of the experiment failed")
	}

}

func view() {

	cli.PreRun()

	rsp, err := spi.ViewExperiments(cli.User, ".*")
	if err != nil {
		cli.Fatal(fmt.Sprintf("view experiments failed: %v", err))
	}
	for _, x := range rsp.Return {
		fmt.Println(x)
	}

}

func changeProfile() {

	if *xpF == "" {
		cli.Fatal(
			"you must specify an experiment name with the -xp flag, " +
				"use the -help flag for details")
	}
	if *nameF == "" {
		cli.Fatal(
			"you must specify an argument name with the -name flag, " +
				"use the -help flag for details")
	}
	if *valueF == "" {
		cli.Fatal(
			"you must specify an argument value with the -value flag, " +
				"use the -help flag for details")
	}

	cli.PreRun()

	attr := spi.ChangeAttribute{}
	attr.Name = *nameF
	attr.Value = *valueF
	attr.Delete = *deleteF

	attrs := []spi.ChangeAttribute{attr}

	rsp, err := spi.ChangeExperimentProfile(*xpF, attrs)

	if err != nil {
		cli.Fatal(fmt.Sprintf("change profile failed: %v", err))
	}
	for _, x := range rsp.Return {
		fmt.Println(x)
	}

}

func changeACL() {

	if *xpF == "" {
		cli.Fatal(
			"you must specify an experiment name with the -xp flag, " +
				"use the -help flag for details")
	}

	if *circF == "" {
		cli.Fatal(
			"you must specify a circle name with the -circle flag, " +
				"use the -help flag for details")
	}

	if *permF == "" {
		cli.Fatal(
			"you must specify a permission name with the -perm flag, " +
				"use the -help flag for details")
	}

	cli.PreRun()

	acm := spi.AccessMember{}
	acm.CircleId = *circF
	acm.Permissions = []string{*permF}

	rsp, err := spi.ChangeExperimentACL(*xpF, []spi.AccessMember{acm})

	if err != nil {
		cli.Fatal(fmt.Sprintf("change acl failed: %v", err))
	}
	for _, x := range rsp.Return {
		fmt.Println(x)
	}

}

func main() {
	flag.Parse()

	switch *opF {
	case "remove":
		remove()
	case "view":
		view()
	case "change-profile":
		changeProfile()
	case "change-acl":
		changeACL()
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
