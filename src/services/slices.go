package objects

import (
	"sort"
)

// Equivalent of an array
func CreateASlice() any {
	var slicees []*User

	newUser := User{
		Name: "Aron",
		Age:  12,
		Sex:  "Male",
	}
	newUser2 := User{
		Name: "Jeff",
		Age:  15,
		Sex:  "Male",
	}
	slicees = append(slicees, &newUser2)
	slicees = append(slicees, &newUser)

	sort.Slice(slicees, func(i int, j int) bool {
		return slicees[i].Name < slicees[j].Name
	})
	return slicees
}
