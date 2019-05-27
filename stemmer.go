package main

import "strings"

var stoppedWords = map[string]bool{
	"a":   true,
	"e":   true,
	"o":   true,
	"da":  true,
	"de":  true,
	"do":  true,
	"na":  true,
	"no":  true,
	"em":  true,
	"com": true,
}

var ignoreSuffixes = [...]string{
	"onassemos",
	"onariamos",
	"onavamos",
	"onasseis",
	"onarieis",
	"onaremos",
	"onaramos",
	"onaveis",
	"onastes",
	"onasses",
	"onassem",
	"onarmos",
	"onarias",
	"onariam",
	"onareis",
	"onardes",
	"issemos",
	"iriamos",
	"essemos",
	"eriamos",
	"assemos",
	"ariamos",
	"onemos",
	"onavas",
	"onavam",
	"onaste",
	"onasse",
	"onaria",
	"onares",
	"onarem",
	"onarei",
	"onaras",
	"onarao",
	"onaram",
	"onamos",
	"isseis",
	"irieis",
	"iremos",
	"iramos",
	"esseis",
	"erieis",
	"eremos",
	"eramos",
	"avamos",
	"asseis",
	"arieis",
	"aremos",
	"aramos",
	"íreis",
	"oneis",
	"onava",
	"onara",
	"onais",
	"istes",
	"isses",
	"issem",
	"irmos",
	"irias",
	"iriam",
	"ireis",
	"irdes",
	"iamos",
	"estes",
	"esses",
	"essem",
	"ermos",
	"erias",
	"eriam",
	"ereis",
	"erdes",
	"aveis",
	"astes",
	"asses",
	"assem",
	"armos",
	"arias",
	"ariam",
	"areis",
	"ardes",
	"onou",
	"ones",
	"onem",
	"onei",
	"onas",
	"onar",
	"onam",
	"iste",
	"isse",
	"iria",
	"ires",
	"irem",
	"irei",
	"iras",
	"irao",
	"iram",
	"imos",
	"imas",
	"imam",
	"ieis",
	"ides",
	"iais",
	"este",
	"esse",
	"eria",
	"eres",
	"erem",
	"erei",
	"eras",
	"erao",
	"eram",
	"emos",
	"avas",
	"avam",
	"aste",
	"asse",
	"aria",
	"ares",
	"arem",
	"arei",
	"aras",
	"arao",
	"aram",
	"amos",
	"ono",
	"one",
	"ona",
	"mos",
	"iri",
	"ira",
	"imo",
	"ima",
	"iem",
	"ide",
	"ias",
	"iam",
	"era",
	"eis",
	"ava",
	"ara",
	"ais",
	"ou",
	"iu",
	"is",
	"ir",
	"im",
	"ia",
	"eu",
	"es",
	"er",
	"em",
	"ei",
	"as",
	"ar",
	"ao",
	"am",
	"ai",
	"o",
	"i",
	"e",
	"a",
}

var stemExceptions = map[string]string{
	"tres":      "tres",
	"frances":   "frances",
	"portugues": "portugues",
	"lapis":     "lapis",
	"onibus":    "onibus",
	"exij":      "exig",
	"fiz":       "faz",
	"fac":       "faz",
	"faç":       "faz",
	"far":       "faz",
	"fizer":     "faz",
	"vir":       "ver",
	"via":       "ver",
	"viam":      "ver",
	"vír":       "ver",
	"víss":      "ver",
	"viss":      "ver",
	"vird":      "ver",
	"vej":       "ver",
	"ved":       "ver",
	"tem":       "ter",
	"com":       "cmr",
}

func pluralToSingular(s string) string {
	z := len(s)
	if s[z-1:] != "s" {
		return s
	} else if v, exists := stemExceptions[s]; exists {
		return v
	} else if z > 2 && s[z-2:z-1] == "n" {
		// albuns, batons, marrons
		return s[:z-2] + "m"
	} else if z > 4 && s[z-2:] == "is" && isOneOf(s[z-3], "aeiou") && !isOneOf(s[z-4], "dmps") {
		// aneis, anzois, jornais
		return s[:z-2] + "l"
	} else if z > 3 && s[z-3:] == "aes" {
		// caes, paes
		return s[:z-2] + "o"
	} else if z > 3 && s[z-3:] == "oes" {
		// leoes
		return s[:z-3] + "ao"
	}
	return s[:z-1]
}

func stemWord(s string) string {
	s = pluralToSingular(s)
	size := len(s)
	for _, suffix := range ignoreSuffixes {
		if size-len(suffix) >= 2 && s[size-len(suffix):] == suffix {
			r := s[:size-len(suffix)]
			if v, exists := stemExceptions[r]; exists {
				return v
			}
			return r
		}
	}
	return s
}

func stemText(s string) ([]string, [][]int) {
	s = removeAccents(s)
	s = strings.ToLower(s)
	words, index := splitWords(s)
	for i, word := range words {
		word = stemWord(word)
		if stoppedWords[word] {
			word = ""
		}
		words[i] = word
	}
	return words, index
}
