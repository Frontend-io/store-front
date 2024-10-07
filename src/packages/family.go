package family

import (
	"encoding/json"
	"fmt"
	"log"
)

type Child struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type Person struct {
	Name     string   `json:"name"`
	Age      int      `json:"age"`
	Children []Person `json:"children"`
}

func FamilyPeople() (string, []Person) {
	peopleAndNames := `[
		{
			"name": "John",
			"age": 30,
			"children": [
				{
					"name": "Jill",
					"age": 5
				}
			]
		},
		{
			"name": "Jane",
			"age": 25,
			"children": [
				{
					"name": "Jack",
					"age": 3
				},
				{
					"name": "Jill",
					"age": 5
				}
			]
		},
		{
			"name": "Joe",
			"age": 20,
			"children": [
				{
					"name": "Jack",
					"age": 3
				},
				{
					"name": "Jill",
					"age": 5
				}
			]
		},
		{
			"name": "Jill",
			"age": 35,
			"children": [
				{
					"name": "Jack",
					"age": 3
				},
				{
					"name": "Jill",
					"age": 5
				}
			]
		}
	]`

	var people []Person

	err := json.Unmarshal([]byte(peopleAndNames), &people)
	if err != nil {
		log.Fatalf("Unable to marshal json data. Err:%v", err)
	}

	// Print names of each person and his children
	// there are x people comprising of - x. person is x years old and has x children who are, x and x years old

	return CallNames(people), people
}

func CallNames(people []Person) string {
	finalSentence := ""
	for i, person := range people {
		if i == 0 {
			finalSentence = fmt.Sprintf("There are %v people comprising of: \n", len(people))
		}

		childrenAndChildren := ""
		for j, child := range person.Children {
			if j == 0 {
				childrenAndChildren += fmt.Sprintf("%v %v", child.Name, child.Age)
				continue
			}
			childrenAndChildren += fmt.Sprintf(" and %v %v", child.Name, child.Age)
		}

		childCountSentence := ""
		if len(person.Children) > 1 {
			childCountSentence = "children who are"
		} else {
			childCountSentence = "a child who is"
		}

		if i == 0 {
			finalSentence += fmt.Sprintf("%v. %v is %v years old and has %v %v , %v", i+1, person.Name, person.Age, childCountSentence, len(person.Children), childrenAndChildren)
			continue
		}

		finalSentence += fmt.Sprintf("\n%v. %v is %v years old and has %v children who are, %v", i+1, person.Name, person.Age, len(person.Children), childrenAndChildren)
	}

	return finalSentence
}
