package main

import (
	"regexp"
	"strings"
	"unicode"

	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

var reHTMLTags = regexp.MustCompile(`<\/?\w[^>]*>`)

var reHTMLComments = regexp.MustCompile(`<!--(.*?)-->`)

func sanitizeString(s string) (r string) {
	return removeHTML(s)
}

func removeHTML(s string) string {
	return removeHTMLTags(removeHTMLComments(s))
}

func removeHTMLComments(s string) string {
	return reHTMLComments.ReplaceAllStringFunc(s, func(match string) string { return strings.Repeat(" ", len(match)) })
}

func removeHTMLTags(s string) string {
	return reHTMLTags.ReplaceAllStringFunc(s, func(match string) string { return strings.Repeat(" ", len(match)) })
}

func isMn(r rune) bool {
	return unicode.Is(unicode.Mn, r)
}

func removeAccents(s string) string {
	t := transform.Chain(norm.NFD, transform.RemoveFunc(isMn), norm.NFC)
	result, _, _ := transform.String(t, s)
	return result
}
