package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"

	"golang.org/x/exp/slices"
)

func findSeatByID(seats []*Seat, id string) *Seat {
	for _, seat := range seats {
		if seat.ID == id {
			return seat
		}
	}
	return nil
}

func ReadSeatsFromFile(path string) ([]*Seat, error) {
	path = "assets/seats/" + path
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	seats := []*Seat{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		ids := strings.Split(scanner.Text(), " ")

		if len(ids) < 2 {
			continue
		}

		currentSeat := findSeatByID(seats, ids[0])
		if currentSeat == nil {
			currentSeat = &Seat{ID: ids[0]}
			seats = append(seats, currentSeat)
		}

		for _, id := range ids[1:] {
			seat := findSeatByID(seats, id)
			if seat == nil {
				seat = &Seat{ID: id}
				seats = append(seats, seat)
			}

			if !slices.Contains(currentSeat.AdjacentSeats, seat) {
				currentSeat.AdjacentSeats = append(currentSeat.AdjacentSeats, seat)
			}
			if !slices.Contains(seat.AdjacentSeats, currentSeat) {
				seat.AdjacentSeats = append(seat.AdjacentSeats, currentSeat)
			}
		}
	}

	return seats, nil
}

func findPersonByName(people []*Person, name string) *Person {
	for _, person := range people {
		if person.Name == name {
			return person
		}
	}
	return nil
}

func ReadPeopleFromFile(path string) ([]*Person, []*Relation, error) {
	path = "assets/people/" + path
	file, err := os.Open(path)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	people := []*Person{}
	relations := []*Relation{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		peopleStrings := strings.Split(scanner.Text(), " ")

		currentPerson := findPersonByName(people, peopleStrings[0])
		if currentPerson == nil {
			currentPerson = &Person{Name: peopleStrings[0]}
			people = append(people, currentPerson)
		}

		for _, relationString := range peopleStrings[1:] {
			split := strings.Split(relationString, "-")
			personName := split[0]
			personRelation := split[1]

			person := findPersonByName(people, personName)
			if person == nil {
				person = &Person{Name: personName}
				people = append(people, person)
			}

			var relation *Relation
			if personRelation == "p" {
				relation = currentPerson.AddRelation(person, 0, true)
			} else {
				relationValue, err := strconv.Atoi(personRelation)
				if err != nil {
					return nil, nil, err
				}

				relation = currentPerson.AddRelation(person, relationValue, false)
			}

			if relation != nil {
				relations = append(relations, relation)
			}
		}
	}

	return people, relations, nil
}
