package spi

import (
	"encoding/xml"
)

type ExperimentAspect struct {
	Data string `xml:"data"`
	Type string `xml:"type"`
}

type CreateExperimentEnvelope struct {
	Envelope
	Body struct {
		Body
		CreateExperiment struct {
			XMLName xml.Name           `xml:"http://api.testbed.deterlab.net/xsd createExperiment"`
			EID     string             `xml:"eid"`
			Owner   string             `xml:"owner"`
			Aspects []ExperimentAspect `xml:"aspects"`
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
