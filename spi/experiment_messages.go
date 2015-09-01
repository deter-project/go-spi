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

type RealizationContainment struct {
	XMLName xml.Name `xml:"http://api.testbed.deterlab.net/xsd containment"`
	Inner   string   `xml:"inner"`
	Outer   string   `xml:"outer"`
}

type RealizationMap struct {
	XMLName      xml.Name `xml:"http://api.testbed.deterlab.net/xsd mapping"`
	Resource     string   `xml:"resource"`
	TopologyName string   `xml:"topologyName"`
}

type RealizationDescription struct {
	XMLName     xml.Name                 `xml:"http://api.testbed.deterlab.net/xsd return"`
	Circle      string                   `xml:"circle"`
	Experiment  string                   `xml:"experiment"`
	Name        string                   `xml:"name"`
	Staus       string                   `xml:"status"`
	Containment []RealizationContainment `xml:"containment"`
	Mapping     []RealizationMap         `xml:"mapping"`
	Perms       []string                 `xml:"perms"`
}

type ExperimentDescription struct {
	XMLName xml.Name `xml:"http://api.testbed.deterlab.net/xsd return"`
	Name    string   `xml:"experimentId"`
	Owner   string   `xml:"owner"`
	Staus   string   `xml:"status"`
	Perms   []string `xml:"perms"`
}

// Remove Realization ---------------------------------------------------------

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

// Release Realization --------------------------------------------------------

type ReleaseRealizationEnvelope struct {
	Envelope
	Body struct {
		Body
		ReleaseRealization struct {
			XMLName xml.Name `xml:"http://api.testbed.deterlab.net/xsd releaseRealization"`
			Name    string   `xml:"name"`
		}
	}
}

type ReleaseRealizationResponse struct {
	Return bool `xml:"return"`
}

type ReleaseRealizationResponseEnvelope struct {
	Envelope
	Body struct {
		Body
		ReleaseRealizationResponse ReleaseRealizationResponse `xml:"releaseRealizationResponse"`
	}
}

// View Realizations -----------------------------------------------------------

type ViewRealizationsEnvelope struct {
	Envelope
	Body struct {
		Body
		ViewRealizations struct {
			XMLName xml.Name `xml:"http://api.testbed.deterlab.net/xsd viewRealizations"`
			UID     string   `xml:"uid"`
			Regex   string   `xml:"regex"`
		}
	}
}

type ViewRealizationsResponse struct {
	Return []RealizationDescription `xml:"return"`
}

type ViewRealizationsResponseEnvelope struct {
	Envelope
	Body struct {
		Body
		ViewRealizationsResponse ViewRealizationsResponse `xml:"viewRealizationsResponse"`
	}
}

// View Experiments -----------------------------------------------------------

type ViewExperimentsEnvelope struct {
	Envelope
	Body struct {
		Body
		ViewExperiments struct {
			XMLName xml.Name `xml:"http://api.testbed.deterlab.net/xsd viewExperiments"`
			UID     string   `xml:"uid"`
			Regex   string   `xml:"regex"`
		}
	}
}

type ViewExperimentsResponse struct {
	Return []ExperimentDescription `xml:"return"`
}

type ViewExperimentsResponseEnvelope struct {
	Envelope
	Body struct {
		Body
		ViewExperimentsResponse ViewExperimentsResponse `xml:"viewExperimentsResponse"`
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

// Change Experiment Profile ---------------------------------------------------

type ChangeExperimentProfileEnvelope struct {
	Envelope
	Body struct {
		Body
		ChangeExperimentProfile struct {
			XMLName xml.Name          `xml:"http://api.testbed.deterlab.net/xsd changeExperimentProfile"`
			EID     string            `xml:"eid"`
			Changes []ChangeAttribute `xml:"changes"`
		}
	}
}

type ChangeAttribute struct {
	XMLName xml.Name `xml:"changes"`
	Delete  bool     `xml:"delete"`
	Name    string   `xml:"name"`
	Value   string   `xml:"value"`
}

type ChangeExperimentProfileResponse struct {
	Return []ChangeResult `xml:"return"`
}

type ChangeExperimentProfileResponseEnvelope struct {
	Envelope
	Body struct {
		Body
		ChangeExperimentProfileResponse ChangeExperimentProfileResponse
	}
}

type ChangeResult struct {
	Name    string `xml:"name"`
	Reason  string `xml:"reason"`
	Success bool   `xml:"success"`
}
