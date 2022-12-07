package main

import (
	"fmt"
)

type Seat struct {
	ID            string
	AdjacentSeats []*Seat
	Person        *Person
}

func (seat *Seat) FirstEmptySeat() *Seat {
	for _, adjacentSeat := range seat.AdjacentSeats {
		if adjacentSeat.Person == nil {
			return adjacentSeat
		}
	}
	return nil
}

type Relation struct {
	Target   *Person
	Value    int
	Partners bool
}

type Bar struct {
	Seats     []*Seat
	People    []*Person
	Relations []*Relation
}

func (bar *Bar) EveryoneSeated() bool {
	for _, seat := range bar.Seats {
		if seat.Person == nil {
			return false
		}
	}
	for _, person := range bar.People {
		if person.AssignedSeat == nil {
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

func main() {
	seats, err := ReadSeatsFromFile("./assets/seats/circleTables")
	if err != nil {
		fmt.Println(err)
		return
	}

	people, relations, err := ReadPeopleFromFile("./assets/people/example")
	if err != nil {
		fmt.Println(err)
		return
	}

	bar := Bar{Seats: seats, People: people, Relations: relations}

	AssignSeatingDFS(&bar)

	fmt.Println(bar.TotalValue() / 2)
}
