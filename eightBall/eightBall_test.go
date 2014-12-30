package eightBall

import (
	"testing"
)


func TestEightBallAnswer(t *testing.T) {

	var answer string = EightBallAnswer()
	t.Logf("First answer: %s", answer)
	
	if answer == "" {
		t.Error("Answer is empty!")
	}
}
