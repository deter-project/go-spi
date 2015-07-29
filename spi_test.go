package spi

import (
	"log"
	"testing"
)

func TestRequestChallenge(t *testing.T) {

	//send challenge and check result
	response, err := RequestChallenge("deterboss")
	if err != nil {
		t.Error(err)
		return
	}
	log.Printf("challengeID: %d\n", response.ChallengeID)

	//respond to challenge
	cresponse, err := ChallengeResponse(response.ChallengeID, "muffins")
	if err != nil {
		t.Error(err)
		return
	}
	log.Printf("challengeResponse accepted\n")
	log.Printf("\n%s\n", cresponse.Return)

	//use the certificate we got back from ChallengeResponse for future comms
	err = setCertificate([]byte(cresponse.Return))
	if err != nil {
		t.Error(err)
		return
	}

	//view deterboss projects
	vpr, err := ViewProjects("deterboss", ".*")
	if err != nil {
		t.Error(err)
		return
	}
	log.Println("deterboss projects:")
	log.Printf("\n%v\n", vpr.Return)

}
