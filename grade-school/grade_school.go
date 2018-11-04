package school

import (
	"sort"
)

type School struct {
	grades map[int][]string
}

type Grade struct {
	grade int
	students []string
}

func New() *School {
	return &School{grades: map[int][]string{}}
}

func (school *School) Enrollment() []Grade {
	return school.sortedListOfGrades()
}

func (school *School) sortedListOfGrades() []Grade {
	grades := school.listOfGrades()
	sort.Slice(grades, func(i, j int) bool {
		return grades[i].grade < grades[j].grade
	})
	return grades
}

func (school *School) listOfGrades() []Grade {
	school.sortStudents()
	grades := []Grade{}
	for g, students := range school.grades {
		grades = append(grades, Grade{grade: g, students: students})
	}
	return grades
}

func (school *School) sortStudents() {
	for _, students := range school.grades {
		sort.Slice(students, func(i, j int) bool {
			return students[i] < students[j]
		})
	}
}

func (school *School) Add(student string, g int) {
	if _, ok := school.grades[g]; !ok {
		school.grades[g] = []string{}
	}
	school.grades[g] = append(school.grades[g], student)
}

func (school *School) Grade(grade int) []string {
	return school.grades[grade]
}