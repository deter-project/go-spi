package spi

import (
	"encoding/base64"
	"errors"
	"fmt"
	"log"
)

//Common Variables -------------------------------------------------------------

const USERS_HTTPS = API_HTTPS + "/Users"

//API calls --------------------------------------------------------------------

/*
RequestChallenge returns the SPI's response to a challenge given a user-id.
*/
func RequestChallenge(uid string) (*RequestChallengeResponse, error) {

	//create the envelope
	e := RequestChallengeEnvelope{}
	e.Body.RequestChallenge.UID = uid

	//allocate a struct for the result
	var rcre RequestChallengeResponseEnvelope

	//make the spi call
	_, _, err := spiCall(USERS_HTTPS+"/requestChallenge", e, &rcre)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	//sanity check on result
	result := &rcre.Body.RequestChallengeResponse.Return.RequestChallengeResponse
	if result.ChallengeID == 0 {
		err := errors.New("Failed to get challengeID from DeterSPI")
		log.Println(err)
		return nil, err
	}

	return result, nil

}

/*
ChallengeResponse returns the SPI's response to a challenge with the certificate
decoded. The provided challengeID must be the result fo a RequestChallenge call
and the password is just a plain-text standard encoded string.
*/
func ChallengeResponse(challengeID int64, password string) (
	*ChallengeResponseResponse, error) {

	//create the envelope
	passB64 := base64.StdEncoding.EncodeToString([]byte(password))
	//log.Printf("encoded password: %s", passB64)
	e := ChallengeResponseEnvelope{}
	e.Body.ChallengeResponse.ResponseData = passB64
	e.Body.ChallengeResponse.ChallengeID = challengeID

	//allocate a struct for the result
	var crre ChallengeResponseResponseEnvelope

	//make the spi call
	rsp, _, err := spiCall(USERS_HTTPS+"/challengeResponse", e, &crre)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	//check the result
	if rsp.StatusCode != 200 {
		return nil, errors.New("server did not accept challenge response")
	}
	crr := &crre.Body.ChallengeResponseResponse
	if crr.Return == "" {
		log.Println("warning: empty certificate, already logged in?")
		return nil, nil
	}

	//decode the certificate
	cert, err := base64.StdEncoding.DecodeString(crr.Return)
	if err != nil {
		log.Println("invalid certificate (base64 decode)")
		return nil, err
	}
	crr.Return = string(cert)

	return crr, nil

}

func Login(user, password string) error {

	//send challenge and check result
	response, err := RequestChallenge(user)
	if err != nil {
		log.Println(err)
		return fmt.Errorf("[Login] Error sending request challenge")
	}
	log.Printf("[Login] challengeID: %d\n", response.ChallengeID)

	//respond to challenge
	cresponse, err := ChallengeResponse(response.ChallengeID, password)
	if err != nil {
		log.Println(err)
		return fmt.Errorf("[Login] Error sending challenge response")
	}
	log.Printf("[Login] challengeResponse accepted\n")
	//log.Printf("\n%s\n", cresponse.Return)

	//use the certificate we got back from ChallengeResponse for future comms
	err = setCertificate([]byte(cresponse.Return))
	if err != nil {
		log.Println(err)
		return fmt.Errorf("[Login] Error setting certificate")
	}

	return nil

}
