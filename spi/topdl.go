package spi

type Topology struct {
	Substrates []Substrate
	Elements   []Element
}

type Substrate struct {
	Name     string
	Capacity Capacity
	Latency  Latency
}

type Element interface{}

type Computer struct {
	Name      string
	Interface []Interface
	OSs       []OS
}

type Interface struct {
	Name      string
	Substrate string
	Capacity  Capacity
	Latency   Latency
}

type OS struct {
	Name    string
	Version string
}

type Capacity struct {
	Rate float64
	Kind Kind
}

type Latency struct {
	Time float64
	Kind Kind
}

type Kind struct {
	Value String //should be either "max" or "average"
}
