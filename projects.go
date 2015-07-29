package spi

import (
	"errors"
	"log"
	"net/http"
)

//Common Variables -------------------------------------------------------------

const PROJECTS_HTTPS = API_HTTPS + "/Projects"

//API calls --------------------------------------------------------------------

func ViewProjects(uid string, regex string) (*ViewProjectsResponse, error) {

	log.Printf("cert count: %d",
		len(client.Transport.(*http.Transport).TLSClientConfig.Certificates))

	//log.Printf("cert count: %d",
	//		client.Transport.(*http.Transport).TLSClientConfig.NameToCertificate)

	//create the envelope
	e := ViewProjectsEnvelope{}
	e.Body.ViewProjects.UID = uid
	e.Body.ViewProjects.Regex = regex

	//allocate struct for the result
	var vpre ViewProjectsResponseEnvelope

	//make the spi call
	rsp, body, err := spiCall(PROJECTS_HTTPS+"/viewProjects", e, &vpre)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	//check the result
	if rsp.StatusCode != 200 {
		log.Printf("[viewProjects] status code %d\n", rsp.StatusCode)
		log.Printf(body)
		return nil, errors.New("server rejected view projects request")
	}

	result := &vpre.Body.ViewProjectsResponse

	return result, nil
}
