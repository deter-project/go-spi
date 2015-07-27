package spi

import (
	"fmt"
	"testing"
)

func TestRequestChallenge(t *testing.T) {

	response, err := requestChallenge("deterboss")

	if err != nil {
		t.Error(err)
		return
	}

	if response.ChallengeID == 0 {
		t.Error("Uninitialized challengeID")
	}

	fmt.Printf("challengeID: %d\n", response.ChallengeID)

}
