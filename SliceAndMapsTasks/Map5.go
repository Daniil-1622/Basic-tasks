package main

import "fmt"

func main() {
	student := map[string]map[string]int{
		"Алиса": {
			"Математика": 5,
		},
	}

	addGrade(student, "Alice", "Матемитика", 8)
	fmt.Println("Student:", student)
}

func addGrade(grades map[string]map[string]int, student, subject string, grade int) {
	if grades[student] == nil {
		grades[student] = make(map[string]int)
	}
	grades[student][subject] = grade
}
