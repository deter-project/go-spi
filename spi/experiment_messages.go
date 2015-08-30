package spi

import (
	"encoding/xml"
)

// Attributes & Aspects --------------------------------------------------------

type ExperimentAspect struct {
	Data string `xml:"data"`
	Type string `xml:"type"`
}

type Attribute interface {
	GetName() string
}

type DescriptionAttr struct {
	Name  string `xml:"name"`
	Value string `xml:"value"`
}

func (d DescriptionAttr) GetName() string {
	return d.Name
}

// Create Experiment -----------------------------------------------------------

type CreateExperimentEnvelope struct {
	Envelope
	Body struct {
		Body
		CreateExperiment struct {
			XMLName xml.Name           `xml:"http://api.testbed.deterlab.net/xsd createExperiment"`
			EID     string             `xml:"eid"`
			Owner   string             `xml:"owner"`
			Aspects []ExperimentAspect `xml:"aspects"`
			Profile []Attribute        `xml:"profile"`
		}
	}
}

type CreateExperimentResponse struct {
	Return bool `xml:"return"`
}

type CreateExperimentResponseEnvelope struct {
	Envelope
	Body struct {
		Body
		CreateExperimentResponse CreateExperimentResponse `xml:"createExperimentResponse"`
	}
}

// Realize Experiment ----------------------------------------------------------

type RealizeExperimentEnvelope struct {
	Envelope
	Body struct {
		Body
		RealizeExperiment struct {
			XMLName xml.Name `xml:"http://api.testbed.deterlab.net/xsd realizeExperiment"`
			UID     string   `xml:"uid"`
			EID     string   `xml:"eid"`
			CID     string   `xml:"cid"`
		}
	}
}

type RealizeExperimentResponse struct {
	Return RealizationDescription `xml:"return"`
}

type RealizeExperimentResponseEnvelope struct {
	Envelope
	Body struct {
		Body
		RealizeExperimentResponse RealizeExperimentResponse `xml:"realizeExperimentResponse"`
	}
}

type RealizationDescription struct {
	XMLName    xml.Name `xml:"http://api.testbed.deterlab.net/xsd RealizationDescription"`
	Circle     string   `xml:"circle"`
	Experiment string   `xml:"experiment"`
	Name       string   `xml:"name"`
	Staus      string   `xml:"status"`
}

// Remove Realization

type RemoveRealizationEnvelope struct {
	Envelope
	Body struct {
		Body
		RemoveRealization struct {
			XMLName xml.Name `xml:"http://api.testbed.deterlab.net/xsd removeRealization"`
			Name    string   `xml:"name"`
		}
	}
}

type RemoveRealizationResponse struct {
	Return bool `xml:"return"`
}

type RemoveRealizationResponseEnvelope struct {
	Envelope
	Body struct {
		Body
		RemoveRealizationResponse RemoveRealizationResponse `xml:"removeRealizationResponse"`
	}
}

// Remove Experiment -----------------------------------------------------------

type RemoveExperimentEnvelope struct {
	Envelope
	Body struct {
		Body
		RemoveExperiment struct {
			XMLName xml.Name `xml:"http://api.testbed.deterlab.net/xsd removeExperiment"`
			EID     string   `xml:"eid"`
		}
	}
}

type RemoveExperimentResponse struct {
	XMLName xml.Name `xml:"http://api.testbed.deterlab.net/xsd removeExperimentResponse"`
	Return  bool     `xml:"return"`
}

type RemoveExperimentResponseEnvelope struct {
	Envelope
	Body struct {
		Body
		RemoveExperimentResponse RemoveExperimentResponse
	}
}
