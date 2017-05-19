package main

import (
	"fmt"
	"regexp"
	"strings"
)

func playgroundMain() {
	var inps string

	inps = "\"@Donald Trump urged the Russians to hack his opponents...That is an established fact.\" (via @MSNBC)\""
	//inps = "RT @TheDemCoalition: Unhinged: @RealDonaldTrump just threatened Comey with possibility of secret \"tapes\" https://t.co/ahPSM1R9iJ via @eenay…"
	fmt.Println(string(inps))

	r, _ := regexp.Compile("https://([\\w\\W]+)")
	out := r.ReplaceAllString(inps, "")
	fmt.Println(out)

	r, _ = regexp.Compile("([\\W\\w\\s]+):")
	out = r.ReplaceAllString(out, "")
	fmt.Println(out)

	r, _ = regexp.Compile("([^\\w@]+)")
	out = r.ReplaceAllString(out, "")
	fmt.Println(out)

	out = strings.TrimSpace(out)
	fmt.Println(out)

	/*r, _ = regexp.Compile("@([\\w]+)")
	out = r.ReplaceAllString(out, "")
	fmt.Println(out)*/
}

//----------------------------------------------
