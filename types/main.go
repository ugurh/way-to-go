package main

import (
	"fmt"
	"time"
)

func main() {

	// array constructed with the basic type uint8
	var arr [3]uint8

	// pointer constructed with the basic type uint8
	var myPointer *uint8

	// function  constructed with the basic type string
	var displayName func(name, firstname string) string

	// slices constructed with the basic type uint8
	var roomNumbers []uint8

	// maps constructed with the basic types uint8 and string
	var score map[string]uint8

	// channel constructed with the basic type bool
	var received chan<- bool

	// struct, interface
	fmt.Println(arr, myPointer, displayName, roomNumbers, score, received)

	// new type "ExchangeRate"
	// underlying type is map[string]float64
	// map[string]float64 is a type litteral
	type ExchangeRate map[string]float64

	// new type "Birthdate"
	// underlying type : time.Time (type Time from the time package)
	type Birthdate time.Time

	// new type "Country"
	// underlying type : struct
	type Country struct {
		Name        string
		CapitalCity string
	}

	// new type "Hotel"
	// underlying type : struct
	type Hotel struct {
		Name     string
		Capacity uint8
		Rooms    uint8
		Smoking  bool
		Country
	}

	fmt.Println(Country{
		Name:        "Turkey",
		CapitalCity: "Ankara",
	})

	france := Country{
		Name:        "France",
		CapitalCity: "Paris",
	}
	fmt.Println(france)

	greece := Country{"Yunanistan", "Atina"}
	fmt.Println(greece)
	fmt.Println(greece.Name)

	usa := Country{
		Name: "United Sates of America",
	}
	usa.CapitalCity = "Washington DC"
	fmt.Println(usa)

	fmt.Println(
		Hotel{
			Name:     "The Asia Hotel",
			Capacity: 100,
			Rooms:    40,
			Smoking:  true,
			Country: Country{
				"Turkey",
				"Ankara"},
		},
	)

	type Episode struct {
		Name string
		*Country
	}

	episode := Episode{
		"Ezel",
		&Country{
			"Turkey",
			"Ankara",
		},
	}
	fmt.Println(*episode.Country)

}
