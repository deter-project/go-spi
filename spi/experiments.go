package spi

import (
	"encoding/base64"
	"fmt"
	"log"
)

//Common Variables -------------------------------------------------------------

const XPS_HTTPS = API_HTTPS + "/Experiments"

//API calls --------------------------------------------------------------------

func CreateExperiment(expId, owner, topdl string) (
	*CreateExperimentResponse, error) {

	e := CreateExperimentEnvelope{}
	e.Body.CreateExperiment.EID = expId
	e.Body.CreateExperiment.Owner = owner
	e.Body.CreateExperiment.Aspects = append(e.Body.CreateExperiment.Aspects, ExperimentAspect{
		Data: base64.StdEncoding.EncodeToString([]byte(topdl)),
		Type: "layout",
	})
	e.Body.CreateExperiment.Profile = append(e.Body.CreateExperiment.Profile, DescriptionAttr{
		"description", "This is not an experiment"})

	var responseEnvelope CreateExperimentResponseEnvelope

	rsp, _, err := spiCall(XPS_HTTPS+"/createExperiment", e, &responseEnvelope)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	if rsp.StatusCode != 200 {
		return nil, fmt.Errorf("Server did not accept the createExperiment call - %d",
			rsp.StatusCode)
	}

	response := &responseEnvelope.Body.CreateExperimentResponse

	return response, nil
}

func RemoveExperiment(expId string) (*RemoveExperimentResponse, error) {

	e := RemoveExperimentEnvelope{}
	e.Body.RemoveExperiment.EID = expId

	var responseEnvelope RemoveExperimentResponseEnvelope

	rsp, _, err := spiCall(XPS_HTTPS+"/removeExperiment", e, &responseEnvelope)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	if rsp.StatusCode != 200 {
		return nil, fmt.Errorf("Server did not accept the removeExperiment call - %d",
			rsp.StatusCode)
	}

	response := &responseEnvelope.Body.RemoveExperimentResponse

	return response, nil
}
