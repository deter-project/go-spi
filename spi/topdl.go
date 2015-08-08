package spi

import (
	"encoding/xml"
)

type Experiment struct {
	XMLName    xml.Name    `xml:"experiment"`
	Substrates []Substrate `xml:"substrates"`
	Elements   []Element   `xml:"elements"`
}

type Substrate struct {
	Name     string   `xml:"name"`
	Capacity Capacity `xml:"capacity"`
	Latency  Latency  `xml:"latency"`
}

type Element interface{}

type Computer struct {
	Name       string      `xml:"name"`
	Interfaces []Interface `xml:"interfaces"`
	OSs        []OS        `xml:"oss"`
}

type Interface struct {
	Name      string   `xml:"name"`
	Substrate string   `xml:"substrate"`
	Capacity  Capacity `xml:"capacity"`
	Latency   Latency  `xml:"latency"`
}

type OS struct {
	Name    string `xml:"name"`
	Version string `xml:"version"`
}

type Capacity struct {
	Rate float64 `xml:"rate"`
	Kind Kind    `xml:"kind"`
}

type Latency struct {
	Time float64 `xml:"time"`
	Kind Kind    `xml:"kind"`
}

type Kind struct {
	//should be either "max" or "average"
	Value string `xml:"value"`
}
