package main

import (
	"errors"
	"fmt"
	"math/rand"
	"reflect"
	"runtime"
)

type Bar struct {
	Name      string
	Seats     []*Seat
	People    []*Person
	Relations []*Relation
}

func (bar *Bar) Shuffle() {
	rand.Shuffle(len(bar.People), func(i, j int) {
		bar.People[i], bar.People[j] = bar.People[j], bar.People[i]
	})

	rand.Shuffle(len(bar.Seats), func(i, j int) {
		bar.Seats[i], bar.Seats[j] = bar.Seats[j], bar.Seats[i]
	})

	rand.Shuffle(len(bar.Relations), func(i, j int) {
		bar.Relations[i], bar.Relations[j] = bar.Relations[j], bar.Relations[i]
	})
}

func (bar *Bar) TestAlgorithms(algorithms []func(bar *Bar)) {
	for _, algorithm := range algorithms {
		total := 0
		runs := 10000
		for i := 0; i < runs; i += 1 {
			bar.Shuffle()
			algorithm(bar)
			total += bar.TotalValue()
			bar.Reset()
		}

		name := runtime.FuncForPC(reflect.ValueOf(algorithm).Pointer()).Name()
		fmt.Printf("%30s on %30s : %10f\n", name, bar.Name, float32(total)/float32(runs*2))
	}
}

func NewBar(seatsPath string, peoplePath string) (*Bar, error) {
	seats, err := ReadSeatsFromFile(seatsPath)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	people, relations, err := ReadPeopleFromFile(peoplePath)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	if len(seats) != len(people) {
		return nil, errors.New("number of people not equal to number of seats")
	}

	return &Bar{Name: seatsPath + " " + peoplePath, Seats: seats, People: people, Relations: relations}, nil
}

func (bar *Bar) Reset() {
	for _, person := range bar.People {
		person.Assign(nil)
	}
}

func (bar *Bar) EveryoneSeated() bool {
	for _, seat := range bar.Seats {
		if seat.Person == nil {
			return false
		}
	}
	return true
}

func (bar *Bar) TotalValue() int {
	total := 0
	for _, person := range bar.People {
		total += person.TotalValue()
	}
	return total
}
