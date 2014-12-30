package questionAnalyzer

// Regex used to identify yes/no questions from
// SUMMARIZATION OF YES/NO QUESTIONS USING A FEATURE FUNCTION MODEL
// by Jing He and Decheng Dai, available here:
// http://jmlr.org/proceedings/papers/v20/he11/he11.pdf

import (
	"regexp"
	"strings"
)

func IsYesNoQuestion(s string) bool {
	beVerbs := []string{"am", "is", "are", "been", "being", "was", "were"}
	modalVerbs := []string{"can", "could", "shall", "should", "will", "would", "may", "might"}
	auxVerbs := []string{"do", "did", "does", "have", "had", "has"}

	pattern := "(?i)^(?:" + strings.Join(append(beVerbs, append(modalVerbs, auxVerbs...)...), "|") + ")[^\\?]*\\?$"
	result, _ := regexp.MatchString(pattern, s)
	return result
}
