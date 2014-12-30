package questionAnalyzer

import (
	"testing"
)

func TestIsYesNoQuestion(t *testing.T) {

	questions := [...]string{"Is there any way out?",
		"Who knows?",
		"Does anyone know?",
		"Does anyone know...",
		"What if God was one of us?",
		"Do you have an ulcer? If so, contact us",
	}

	expected := [len(questions)]bool{true, false, true, false, false, false}
	
	for i,question := range questions {
		if expected[i] != IsYesNoQuestion(question) {
			t.Errorf("For question '%s', I expected %t, but got %t.\n", question, expected[i], !expected[i])
		} else {
			t.Logf("For question '%s', I passed!.\n", question)
		}
	}
}

