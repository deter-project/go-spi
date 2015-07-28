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

}
