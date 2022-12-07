package main

import "sort"

func (person *Person) potentialHighestValue() int {
	total := 0
	for _, relation := range person.Relations {
		total += relation.Value
	}
	return total
}

func AssignSeatingNaive(bar *Bar) {
	sort.Slice(bar.People, func(i, j int) bool {
		return bar.People[i].potentialHighestValue() < bar.People[j].potentialHighestValue()
	})

	sort.Slice(bar.Seats, func(i, j int) bool {
		return len(bar.Seats[i].AdjacentSeats) < len(bar.Seats[j].AdjacentSeats)
	})

	for i, person := range bar.People {
		person.Assign(bar.Seats[i])
	}
}

func AssignSeatingDFS(bar *Bar) {
	currentSeat := bar.Seats[0]
	currentPerson := bar.People[0]

	for !bar.EveryoneSeated() {
		currentPerson.Assign(currentSeat)
		currentPerson = currentPerson.Relations[0].Target

		if currentPerson.AssignedSeat != nil {
			for _, person := range bar.People {
				if person.AssignedSeat == nil {
					currentPerson = person
				}
			}
		}

		currentSeat = currentSeat.FirstEmptySeat()
		if currentSeat == nil {
			for _, seat := range bar.Seats {
				if seat.Person == nil {
					currentSeat = seat
				}
			}
		}
	}
}
