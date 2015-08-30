package spi

import (
	"encoding/base64"
	"fmt"
	"log"
)

//Common Variables -------------------------------------------------------------

const XPS_HTTPS = API_HTTPS + "/Experiments"
const REX_HTTPS = API_HTTPS + "/Realizations"

//API calls --------------------------------------------------------------------

// Create Experiment -----------------------------------------------------------

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

// Realize Experiment ----------------------------------------------------------

func RealizeExperiment(expId, circle, owner string) (
	*RealizeExperimentResponse, error) {

	e := RealizeExperimentEnvelope{}
	e.Body.RealizeExperiment.EID = expId
	e.Body.RealizeExperiment.UID = owner
	e.Body.RealizeExperiment.CID = circle

	var responseEnvelope RealizeExperimentResponseEnvelope

	rsp, _, err := spiCall(XPS_HTTPS+"/realizeExperiment", e, &responseEnvelope)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	if rsp.StatusCode != 200 {
		return nil, fmt.Errorf("Server did not accept the realizeExperiment call - %d",
			rsp.StatusCode)
	}

	response := &responseEnvelope.Body.RealizeExperimentResponse

	return response, nil

}

// Remove Realization ----------------------------------------------------------

func RemoveRealization(expId string) (*RemoveRealizationResponse, error) {

	e := RemoveRealizationEnvelope{}
	e.Body.RemoveRealization.Name = expId

	var responseEnvelope RemoveRealizationResponseEnvelope

	rsp, _, err := spiCall(REX_HTTPS+"/removeRealization", e, &responseEnvelope)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	if rsp.StatusCode != 200 {
		return nil, fmt.Errorf("Server did not accept the removeRealization call - %d",
			rsp.StatusCode)
	}

	response := &responseEnvelope.Body.RemoveRealizationResponse

	return response, nil
}

// Realize Experiment ----------------------------------------------------------

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
