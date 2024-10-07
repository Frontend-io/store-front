package objects

type User struct {
	Name string
	Age  int
	Sex  string
}

func MapPerson() map[string]interface{} {
	userMap := make(map[string]interface{})
	person := User{
		Name: "Jeffers",
		Age:  23,
		Sex:  "Male",
	}

	userMap["person1"] = person
	return userMap
}
