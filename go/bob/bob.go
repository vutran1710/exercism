package bob

import (
	r "regexp"
	str "strings"
)

func Hey(remark string) string {
	const (
		yellQuestionReply = "Calm down, I know what I'm doing!"
		yellReply = "Whoa, chill out!"
		questionReply = "Sure."
		silenceReply = "Fine. Be that way!"
		other = "Whatever."
	)
	
	isShouting := func(a string) bool {
		upper := str.ToUpper(a)
		notOnlyNumber, _ := r.Compile("[a-zA-Z]")
		return notOnlyNumber.MatchString(a) && str.Compare(upper, a) == 0
	}

	isQuestion := func(a string) bool {
		trimmed := str.Trim(a, " ")
		questionRegex, _ := r.Compile(".+\\?$")
		return questionRegex.MatchString(trimmed)
	}

	isYellingQuestion := func(a string) bool {
		return isShouting(a) && isQuestion(a)
	}

	isSilence := func(a string) bool {
		nonWordStr, _ := r.Compile("\\W+")
		regexTrimmedStr := nonWordStr.ReplaceAllLiteralString(a, "")
		trimmed := str.Trim(regexTrimmedStr, " ")
		return trimmed == ""
	}

	if isYellingQuestion(remark) { return yellQuestionReply }
	if isShouting(remark) { return yellReply }
	if isQuestion(remark) { return questionReply }
	if isSilence(remark) { return silenceReply }
	return other
}
