package main

import "testing"

func TestAverageGrade(t *testing.T) {
	mathan := Subject{"Математичний аналіз", 5}
	discrete := Subject{"Дискретна математика", 4}
	programming := Subject{"Програмування", 6}
	physics := Subject{"Фізика", 3}

	s := Student{
		Name:   "Єлизавета Алам",
		Course: 4,
		Grades: map[Subject]int{
			mathan:      95,
			discrete:    92,
			programming: 98,
			physics:     90,
		},
	}

	result := s.AverageGrade()
	expected := 94.5

	if result != expected {
		t.Errorf("Очікували %.2f, отримали %.2f", expected, result)
	}
}

func TestAverageGrade_Empty(t *testing.T) {
	s := Student{
		Name:   "Без оцінок",
		Course: 1,
		Grades: map[Subject]int{},
	}

	result := s.AverageGrade()

	if result != 0 {
		t.Errorf("Очікували 0, отримали %.2f", result)
	}
}

func TestFindBestStudent(t *testing.T) {
	mathan := Subject{"Математичний аналіз", 5}

	s1 := Student{
		Name:   "Анастасія",
		Course: 2,
		Grades: map[Subject]int{
			mathan: 80,
		},
	}

	s2 := Student{
		Name:   "Вікторія",
		Course: 1,
		Grades: map[Subject]int{
			mathan: 100,
		},
	}

	students := []Student{s1, s2}
	best := FindBestStudent(students)

	if best.Name != "Вікторія" {
		t.Errorf("Очікували Вікторія, отримали %s", best.Name)
	}
}

func TestAverageGrade_TableDriven(t *testing.T) {
	mathan := Subject{"Математичний аналіз", 5}
	discrete := Subject{"Дискретна математика", 4}
	programming := Subject{"Програмування", 6}

	tests := []struct {
		name     string
		grades   map[Subject]int
		expected float64
	}{
		{
			name: "звичайний випадок",
			grades: map[Subject]int{
				mathan:      95,
				discrete:    92,
				programming: 98,
			},
			expected: 95.4,
		},
		{
			name:     "порожній map",
			grades:   map[Subject]int{},
			expected: 0,
		},
		{
			name: "один предмет",
			grades: map[Subject]int{
				mathan: 80,
			},
			expected: 80,
		},
		{
			name: "нульова оцінка",
			grades: map[Subject]int{
				mathan:   0,
				discrete: 100,
			},
			expected: 44.44444444444444,
		},
		{
			name: "максимум",
			grades: map[Subject]int{
				mathan:      100,
				discrete:    100,
				programming: 100,
			},
			expected: 100,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Student{
				Name:   "Тест",
				Course: 1,
				Grades: tt.grades,
			}

			result := s.AverageGrade()

			if result != tt.expected {
				t.Errorf("%s: очікували %f, отримали %f", tt.name, tt.expected, result)
			}
		})
	}
}

func TestFindBestStudent_Table(t *testing.T) {
	mathan := Subject{"Математичний аналіз", 5}

	tests := []struct {
		name     string
		students []Student
		expected string
	}{
		{
			name: "два студенти",
			students: []Student{
				{
					Name: "Перший",
					Grades: map[Subject]int{
						mathan: 80,
					},
				},
				{
					Name: "Другий",
					Grades: map[Subject]int{
						mathan: 95,
					},
				},
			},
			expected: "Другий",
		},
		{
			name: "один студент",
			students: []Student{
				{
					Name: "Єдиний",
					Grades: map[Subject]int{
						mathan: 100,
					},
				},
			},
			expected: "Єдиний",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			best := FindBestStudent(tt.students)

			if best.Name != tt.expected {
				t.Errorf("Очікували %s, отримали %s", tt.expected, best.Name)
			}
		})
	}
}
