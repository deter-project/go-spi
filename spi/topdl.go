package spi

import (
	"encoding/xml"
)

type Experiment struct {
	XMLName    xml.Name    `xml:"experiment"`
	Substrates []Substrate `xml:"substrates"`
	Elements   Elements    `xml:"elements"`
}

type Substrate struct {
	Name     string   `xml:"name"`
	Capacity Capacity `xml:"capacity"`
	Latency  Latency  `xml:"latency"`
}

type Element interface{}

type Elements struct {
	XMLName  xml.Name  `xml:"elements"`
	Elements []Element `xml:"element"`
}

type Computer struct {
	XMLName    xml.Name    `xml:"computer"`
	Name       string      `xml:"name"`
	Interfaces []Interface //`xml:"interfaces"`
	OSs        []OS        //`xml:"os"`
	Attributes []TopDLAttribute
}

type Interface struct {
	XMLName   xml.Name `xml:"interface"`
	Name      string   `xml:"name"`
	Substrate string   `xml:"substrate"`
	Capacity  Capacity `xml:"capacity"`
	Latency   Latency  `xml:"latency"`
}

type OS struct {
	XMLName xml.Name `xml:"os"`
	Name    string   `xml:"name"`
	Version string   `xml:"version"`
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

type TopDLAttribute struct {
	XMLName   xml.Name `xml:"attribute"`
	Attribute string   `xml:"attribute"`
	Value     string   `xml:"value"`
}
