package main

import (
	"fmt"
	s "strings"
	"unicode"
)

var f = fmt.Printf

func main() {
	upper := s.ToUpper("Hello there!")
	f("To Upper: %s\n", upper)
	f("To Lower: %s\n", s.ToLower(upper))
	f("%s\n", s.Title("this wiLL be a Title!"))
	f("EqualFold: %v\n", s.EqualFold("Oleg", "OLEG"))
	f("EqualFold: %v\n", s.EqualFold("Oleg", "OLE"))

	f("Prefix: %v\n", s.HasPrefix("Oleg", "Ol"))
	f("Prefix: %v\n", s.HasPrefix("Oleg", "ol"))
	f("Suffix: %v\n", s.HasSuffix("Oleg", "eg"))
	f("Suffix: %v\n", s.HasSuffix("Oleg", "EG"))
	f("Index: %v\n", s.Index("Oleg", "le"))
	f("Index: %v\n", s.Index("Oleg", "Le"))
	f("Count: %v\n", s.Count("Oleg", "l"))
	f("Count: %v\n", s.Count("Oleg", "L"))
	f("Repeat: %v\n", s.Repeat("Oleg", 5))

	f("TrimSpace: %s\n", s.TrimSpace(" \tThis is a line. \n"))
	f("TrimLeft: %s\n", s.TrimLeft(" \tThis is a line. \n", "\n\t "))
	f("TrimRight: %s\n", s.TrimRight(" \tThis is a line. \n", "\n\t "))

	f("Compare: %v\n", s.Compare("Oleg", "OLEG"))
	f("Compare: %v\n", s.Compare("Oleg", "Oleg"))
	f("Compare: %v\n", s.Compare("OLEG", "OLeg"))

	f("Fields: %v\n", s.Fields("This is a string!"))
	f("Fields: %v\n", s.Fields("This\nis\na\tstring!"))

	f("%s\n", s.Split("abcdefg", ""))

	f("%s\n", s.Replace("adcd efg", "", "_", -1))
	f("%s\n", s.Replace("adcd efg", "", "_", 4))
	f("%s\n", s.Replace("adcd efg", "", "_", 2))

	lines := []string{"Line 1", "Line 2", "Line 3"}
	f("Join: %s\n", s.Join(lines, "+++"))

	f("SplitAfter: %s\n", s.SplitAfter("123++432++", "++"))

	trimFuction := func(c rune) bool {
		return !unicode.IsLetter(c)
	}
	f("TrimFunc: %s\n", s.TrimFunc("123 abc ABC \t .", trimFuction))
}
