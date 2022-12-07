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

func main() {
	algorithms := []func(bar *Bar){AssignSeatingDFS, AssignSeatingNaive, AssignSeatingRandom}

	bar, err := NewBar("circleTables", "10people")
	if err != nil {
		fmt.Print(err)
		return
	}

	bar.TestAlgorithms(algorithms)
	fmt.Println()

	bar, err = NewBar("5table", "10people")
	if err != nil {
		fmt.Print(err)
		return
	}

	bar.TestAlgorithms(algorithms)
	fmt.Println()

	bar, err = NewBar("cinema", "10people")
	if err != nil {
		fmt.Print(err)
		return
	}

	bar.TestAlgorithms(algorithms)
	fmt.Println()
}
