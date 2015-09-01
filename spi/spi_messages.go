package spi

import (
	"fmt"
)

type FaultEnvelope struct {
	Envelope
	Body struct {
		Body
		Fault struct {
			Code struct {
				Value string
			}
			Reason struct {
				Text string
			}
			Detail struct {
				ExperimentsDeterFault struct {
					//XMLName    xml.Name `xml:"http://api.testbed.deterlab.net/xsd ExperimentsDeterFault"`
					DeterFault struct {
						//XMLName       xml.Name `xml:"http://api.testbed.deterlab.net/xsd DeterFault"`
						DetailMessage string `xml:"detailMessage"`
						ErrorCode     int    `xml:"errorCode"`
						ErrorMessage  string `xml:"errorMessage"`
					}
				}
			}
		}
	}
}

func (f FaultEnvelope) String() string {
	df := &f.Body.Fault.Detail.ExperimentsDeterFault.DeterFault
	return fmt.Sprintf("%s (%d): %s", df.ErrorMessage, df.ErrorCode, df.DetailMessage)
}
