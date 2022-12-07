package main

type Person struct {
	Name         string
	Relations    []*Relation
	AssignedSeat *Seat
}

func (person *Person) AddRelation(other *Person, value int, partners bool) *Relation {
	relation := &Relation{Target: other, Value: value, Partners: partners}
	if partners {
		relation.Value = 0
	}

	person.Relations = append(person.Relations, relation)

	return relation
}

func (person *Person) Assign(seat *Seat) {
	if seat.Person != nil {
		seat.Person.AssignedSeat = nil
	}

	seat.Person = person
	person.AssignedSeat = seat
}

func (person *Person) GetRelationValue(other *Person) int {
	for _, relation := range person.Relations {
		if relation.Target == other {
			return relation.Value
		}
	}
	return 0
}

func (person *Person) TotalValue() int {
	total := 0
	for _, adjacentSeat := range person.AssignedSeat.AdjacentSeats {
		if adjacentSeat.Person != nil {
			total += person.GetRelationValue(adjacentSeat.Person)
		}
	}
	return total
}
