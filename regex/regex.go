package regex

import "regexp"

func normal(input string) {
	r :=  regexp.MustCompile(`(?i)(https?|ftp)://(-\.)?([^\s/?\.#]+\.?)+(/[^\s]*)?`)
	r.MatchString(input)
}

var r =  regexp.MustCompile(`(?i)(https?|ftp)://(-\.)?([^\s/?\.#]+\.?)+(/[^\s]*)?`)
func cached(input string) {
	r.MatchString(input)
}