package spi

import (
	"encoding/xml"
	"fmt"
	"testing"
)

func TestTwoLink(t *testing.T) {

	var topo Experiment

	var a Computer
	a.Name = "a"
	a.Interfaces = append(a.Interfaces, Interface{
		Name:      "ifx0",
		Substrate: "ab-link",
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

	var ab Substrate
	ab.Name = "ab-link"
	ab.Capacity = Capacity{1000, Kind{"max"}}
	ab.Latency = Latency{10, Kind{"max"}}

	topo.Elements.Elements = append(topo.Elements.Elements, a)
	topo.Elements.Elements = append(topo.Elements.Elements, b)
	topo.Substrates = append(topo.Substrates, ab)

	topdl, err := xml.MarshalIndent(topo, "  ", "  ")
	if err != nil {
		fmt.Println(err)
		t.Error("failed to serialize topology to topdl xml")
	}
	fmt.Println(string(topdl))

	//create a session with the DeterLab SPI
	err = Login("deterboss", "muffins")
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
