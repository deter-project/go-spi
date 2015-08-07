package spi

import (
	"encoding/xml"
)

type RequestChallengeEnvelope struct {
	Envelope
	Body struct {
		Body
		RequestChallenge struct {
			XMLName xml.Name `xml:"http://api.testbed.deterlab.net/xsd requestChallenge"`
			UID     string   `xml:"uid"`
		}
	}
}

type RequestChallengeResponse struct {
	ChallengeID int64  `xml:"http://api.testbed.deterlab.net/xsd challengeID"`
	Validity    string `xml:"http://api.testbed.deterlab.net/xsd validity"`
	Type        string `xml:"http://api.testbed.deterlab.net/xsd type"`
}

type RequestChallengeResponseEnvelope struct {
	Envelope
	Body struct {
		Body
		RequestChallengeResponse struct {
			XMLName xml.Name `xml:"http://api.testbed.deterlab.net/xsd requestChallengeResponse"`
			Return  struct {
				XMLName xml.Name `xml:"http://api.testbed.deterlab.net/xsd return"`
				RequestChallengeResponse
			}
		}
	}
}

type ChallengeResponseEnvelope struct {
	Envelope
	Body struct {
		Body
		ChallengeResponse struct {
			XMLName      xml.Name `xml:"http://api.testbed.deterlab.net/xsd challengeResponse"`
			ResponseData string   `xml:"http://api.testbed.deterlab.net/xsd responseData"`
			ChallengeID  int64    `xml:"http://api.testbed.deterlab.net/xsd challengeID"`
		}
	}
}

type ChallengeResponseResponse struct {
	Return string `xml:"http://api.testbed.deterlab.net/xsd return"`
}

type ChallengeResponseResponseEnvelope struct {
	Envelope
	Body struct {
		Body
		ChallengeResponseResponse ChallengeResponseResponse `xml:"http://api.testbed.deterlab.net/xsd challengeResponseResponse"`
	}
}
