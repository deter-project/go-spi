package spi

import (
	"bytes"
	"crypto/tls"
	"encoding/xml"
	"errors"
	"log"
	"net/http"
)

//Constants
const USERS_HTTPS = API_HTTPS + "/Users"

//Common variables
var tr = &http.Transport{
	TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
}
var client = &http.Client{Transport: tr}

//API calls -------------------------------------------------------------------

func requestChallenge(uid string) (*RequestChallengeResponse, error) {

	addr := USERS_HTTPS + "/requestChallenge"

	//create the envelope
	e := RequestChallengeEnvelope{}
	e.Body.RequestChallenge.UID = uid
	msg, err := xml.Marshal(&e)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	//create the request
	req, err := http.NewRequest("POST", addr, bytes.NewBuffer(msg))
	if err != nil {
		log.Println(err)
		return nil, err
	}
	req.Header.Add("Content-Type", "application/soap+xml")

	//send the request
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	//read the result
	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	bs := buf.String()

	//unmarshal the result
	var rcre RequestChallengeResponseEnvelope
	xml.Unmarshal([]byte(bs), &rcre)
	result := &rcre.Body.RequestChallengeResponse.Return.RequestChallengeResponse

	//sanity check on result
	if result.ChallengeID == 0 {
		err := errors.New("Failed to get challengeID from DeterSPI")
		log.Println(err)
		return nil, err
	}

	return result, nil

}
