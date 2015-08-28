package spi

import (
	"encoding/xml"
	"fmt"
	"testing"
)

func TestThreeLink(t *testing.T) {

	var topo Experiment

	var a Computer
	a.Name = "a"
	a.Interfaces = append(a.Interfaces, Interface{
		Name:      "ifx0",
		Substrate: "ab-link",
		Capacity:  Capacity{1000, Kind{"max"}},
		Latency:   Latency{10, Kind{"max"}},
	})
	a.Interfaces = append(a.Interfaces, Interface{
		Name:      "ifx1",
		Substrate: "ca-link",
		Capacity:  Capacity{1000, Kind{"max"}},
		Latency:   Latency{10, Kind{"max"}},
	})

	var b Computer
	b.Name = "b"
	b.Interfaces = append(b.Interfaces, Interface{
		Name:      "ifx0",
		Substrate: "ab-link",
		Capacity:  Capacity{1000, Kind{"max"}},
		Latency:   Latency{10, Kind{"max"}},
	})
	b.Interfaces = append(b.Interfaces, Interface{
		Name:      "ifx1",
		Substrate: "bc-link",
		Capacity:  Capacity{1000, Kind{"max"}},
		Latency:   Latency{10, Kind{"max"}},
	})

	var c Computer
	c.Name = "c"
	c.Interfaces = append(c.Interfaces, Interface{
		Name:      "ifx0",
		Substrate: "bc-link",
		Capacity:  Capacity{1000, Kind{"max"}},
		Latency:   Latency{10, Kind{"max"}},
	})
	c.Interfaces = append(c.Interfaces, Interface{
		Name:      "ifx0",
		Substrate: "ca-link",
		Capacity:  Capacity{1000, Kind{"max"}},
		Latency:   Latency{10, Kind{"max"}},
	})

	var ab Substrate
	ab.Name = "ab-link"
	ab.Capacity = Capacity{1000, Kind{"max"}}
	ab.Latency = Latency{10, Kind{"max"}}

	var bc Substrate
	bc.Name = "bc-link"
	bc.Capacity = Capacity{1000, Kind{"max"}}
	bc.Latency = Latency{10, Kind{"max"}}

	var ca Substrate
	ca.Name = "ca-link"
	ca.Capacity = Capacity{1000, Kind{"max"}}
	ca.Latency = Latency{10, Kind{"max"}}

	topo.Elements.Elements = append(topo.Elements.Elements, a)
	topo.Elements.Elements = append(topo.Elements.Elements, b)
	topo.Elements.Elements = append(topo.Elements.Elements, c)
	topo.Substrates = append(topo.Substrates, ab)
	topo.Substrates = append(topo.Substrates, bc)
	topo.Substrates = append(topo.Substrates, ca)

	topdl, err := xml.MarshalIndent(topo, "  ", "  ")
	if err != nil {
		fmt.Println(err)
		t.Error("failed to serialize topology to topdl xml")
	}
	fmt.Println(string(topdl))

	//remove this if you have an spi running and want to test that interaction
	return

	//create a session with the DeterLab SPI
	_, err = Login("deterboss", "muffins")
	if err != nil {
		t.Log(err)
		t.Error("unable to login as deterboss")
	}

	createResponse, err := CreateExperiment("deterboss:twolink", "deterboss", string(topdl))
	if err != nil {
		t.Log(err)
		t.Error("failed to create experiment")
	}
	t.Log(createResponse)

	removeResponse, err := RemoveExperiment("deterboss:twolink")
	if err != nil {
		t.Log(err)
		t.Error("failed to remove experiment")
	}
	t.Log(removeResponse)

}

func TestThreeLinkLan(t *testing.T) {

	var topo Experiment

	var a Computer
	a.Name = "a"
	a.Interfaces = append(a.Interfaces, Interface{
		Name:      "ifx0",
		Substrate: "lan",
		Capacity:  Capacity{1000, Kind{"max"}},
		Latency:   Latency{10, Kind{"max"}},
	})

	var b Computer
	b.Name = "b"
	b.Interfaces = append(b.Interfaces, Interface{
		Name:      "ifx0",
		Substrate: "lan",
		Capacity:  Capacity{1000, Kind{"max"}},
		Latency:   Latency{10, Kind{"max"}},
	})

	var c Computer
	c.Name = "c"
	c.Interfaces = append(c.Interfaces, Interface{
		Name:      "ifx0",
		Substrate: "lan",
		Capacity:  Capacity{1000, Kind{"max"}},
		Latency:   Latency{10, Kind{"max"}},
	})

	var lan Substrate
	lan.Name = "lan"
	lan.Capacity = Capacity{1000, Kind{"max"}}
	lan.Latency = Latency{10, Kind{"max"}}

	topo.Elements.Elements = append(topo.Elements.Elements, a)
	topo.Elements.Elements = append(topo.Elements.Elements, b)
	topo.Elements.Elements = append(topo.Elements.Elements, c)
	topo.Substrates = append(topo.Substrates, lan)

	topdl, err := xml.MarshalIndent(topo, "  ", "  ")
	if err != nil {
		fmt.Println(err)
		t.Error("failed to serialize topology to topdl xml")
	}
	fmt.Println(string(topdl))

}
