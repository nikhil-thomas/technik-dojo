package main

import (
	"fmt"
	s "strings"
	"unicode"
)

func main() {
	var f = fmt.Printf

	f("To upper: %s\n", s.ToUpper("hello There!"))
	f("To lower: %s\n", s.ToLower("hello There!"))
	f("To title: %s\n", s.ToTitle("thiS is a title!"))
	f("Equalfold: %v\n", s.EqualFold("Name", "NAME"))
	f("Equalfold: %v\n", s.EqualFold("Name", "NAE"))

	f("Prefix: %v\n", s.HasPrefix("Name", "Na"))
	f("Prefix: %v\n", s.HasPrefix("Name", "na"))

	f("Suffix: %v\n", s.HasSuffix("Name", "me"))
	f("Suffix: %v\n", s.HasSuffix("Name", "ME"))

	f("Index: %v\n", s.Index("Name", "am"))
	f("Index: %v\n", s.Index("Name", "Am"))
	f("Count: %v\n", s.Count("NameNameName", "e"))
	f("Count: %v\n", s.Count("NameNameName", "E"))

	f("Repeat]: %s\n", s.Repeat("xy", 5))

	f("Trimspace: %s\n", s.TrimSpace(" \tThis is a line. \n"))
	f("TrimLeft: %s\n", s.TrimLeft(" \tThis is a\t line. \n", "\n\t "))
	f("TrimRight: %s\n", s.TrimRight(" \tThis is a\t line. \n", "\n\t "))

	f("Compare: %v\n", s.Compare("Name", "NAME"))
	f("Compare: %v\n", s.Compare("Name", "name"))
	f("Compare: %v\n", s.Compare("Name", "NaMe"))

	f("Fields: %v\n", s.Fields("THis is a string"))
	f("Fields: %v\n", s.Fields("THisis\na\tstring"))

	f("%s\n ", s.Split("abcd efg", ""))

	f("%s\n", s.Replace("abcd efg", "", "_", -1))
	f("%s\n", s.Replace("abcd efg", "", "_", 4))
	f("%s\n", s.Replace("abcd efg", "", "_", 2))

	lines := []string{
		"Line 1",
		"Line 2",
		"Line 3",
	}

	f("Join: %s\n", s.Join(lines, "+++"))

	f("Split After %s\n", s.SplitAfter("abc***def***ghi", "***"))

	trimFunc := func(c rune) bool {
		return !unicode.IsLetter(c)
	}

	f("TrimFunc: %s\n", s.TrimFunc("123 abc ABC \t .", trimFunc))

}
