package main

import (
	"fmt"
	"time"
)

// UTC: Universal Coordinated Time

var week time.Duration

func main() {

	timeFormat := "02 Jan 2006 15:04"
	t := time.Now().UTC()
	fmt.Println(t)
	fmt.Printf("%02d.%02d.%4d\n", t.Day(), t.Month(), t.Year())
	fmt.Println(t.Format(timeFormat))

	week = 60 * 60 * 24 * 7 * 1e9 // in nano seconds
	t = t.Add(week)
	fmt.Println(t.Format(timeFormat))
	fmt.Println(t.Format(time.RFC822))
	fmt.Println(t.Format(time.ANSIC))

	s := t.Format("2006 01 02")
	fmt.Println(t, "=>", s)

}
