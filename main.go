package main

import (
	"fmt"
)

type Subject struct {
	Name    string
	Credits int
}

type Student struct {
	Name   string
	Course int
	Grades map[Subject]int
}

func (s Student) AverageGrade() float64 {
	totalScore := 0
	totalCredits := 0

	for subject, grade := range s.Grades {
		totalScore += grade * subject.Credits
		totalCredits += subject.Credits
	}

	if totalCredits == 0 {
		return 0
	}

	return float64(totalScore) / float64(totalCredits)
}

func FindBestStudent(students []Student) Student {
	best := students[0]

	for _, s := range students {
		if s.AverageGrade() > best.AverageGrade() {
			best = s
		}
	}

	return best
}

func main() {

	mathan := Subject{"Математичний аналіз", 5}
	discrete := Subject{"Дискретна математика", 4}
	programming := Subject{"Програмування", 6}
	physics := Subject{"Фізика", 3}

	students := []Student{
		{
			Name:   "Єлизавета Алам",
			Course: 4,
			Grades: map[Subject]int{
				mathan:      95,
				discrete:    92,
				programming: 98,
				physics:     90,
			},
		},
		{
			Name:   "Березовська Анастасія",
			Course: 3,
			Grades: map[Subject]int{
				mathan:      88,
				discrete:    85,
				programming: 90,
				physics:     84,
			},
		},
		{
			Name:   "Благодарна Вікторія",
			Course: 4,
			Grades: map[Subject]int{
				mathan:      91,
				discrete:    89,
				programming: 93,
				physics:     87,
			},
		},
		{
			Name:   "Мороз Дмитро",
			Course: 4,
			Grades: map[Subject]int{
				mathan:      76,
				discrete:    80,
				programming: 78,
				physics:     75,
			},
		},
	}

	fmt.Println("Список студентів та їх середній бал:")
	fmt.Println("-----------------------------------")

	for _, s := range students {
		fmt.Printf("%s (курс %d): %.2f\n", s.Name, s.Course, s.AverageGrade())
	}

	best := FindBestStudent(students)

	fmt.Println("-----------------------------------")
	fmt.Println("Студент з найвищим середнім балом:")
	fmt.Printf("%s (курс %d) — %.2f\n", best.Name, best.Course, best.AverageGrade())
}
