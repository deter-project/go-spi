package spi

import (
	"encoding/xml"
)

type ViewProjectsEnvelope struct {
	Envelope
	Body struct {
		Body
		ViewProjects struct {
			XMLName xml.Name `xml:"http://api.testbed.deterlab.net/xsd viewProjects"`
			UID     string   `xml:"uid"`
			Regex   string   `xml:"regex"`
		}
	}
}

type ProjectDescription struct {
	Approved  bool   `xml:"http://api.testbed.deterlab.net/xsd approved"`
	Members   Member `xml:"http://api.testbed.deterlab.net/xsd members"`
	Owner     string `xml:"http://api.testbed.deterlab.net/xsd owner"`
	ProjectId string `xml:"http://api.testbed.deterlab.net/xsd projectId"`
}

type Member struct {
	Permissions string `xml:"http://api.testbed.deterlab.net/xsd permissions"`
	uid         string `xml:"http://api.testbed.deterlab.net/xsd uid"`
}

type ViewProjectsResponse struct {
	Return []ProjectDescription `xml:"http://api.testbed.deterlab.net/xsd return"`
}

type ViewProjectsResponseEnvelope struct {
	Envelope
	Body struct {
		Body
		ViewProjectsResponse ViewProjectsResponse `xml:"http://api.testbed.deterlab.net/xsd viewProjectsResponse"`
	}
}
