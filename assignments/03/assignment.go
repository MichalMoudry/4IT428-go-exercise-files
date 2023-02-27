package assignment_03

import (
	"errors"
	"strings"
)

// Method for reversing array of strings.
func Reverse(input []string) []string {
	for i, j := 0, len(input)-1; i < j; i, j = i+1, j-1 {
		input[i], input[j] = input[j], input[i]
	}
	return input
}

func Palindrome(input []string) bool {
	input_length := len(input)
	for index, value := range input {
		if input[(index-input_length+1)*-1] != value {
			return false
		}
	}
	return true
}

func Anagram(firstString string, secondString string) bool {
	if len(firstString) != len(secondString) {
		return false
	}

	firstStringFrequencies := make(map[rune]int)
	secondStringFrequencies := make(map[rune]int)

	for index, char := range firstString {
		firstStringFrequencies[char] += 1
		secondStringFrequencies[rune(secondString[index])] += 1
	}
	for index := range firstStringFrequencies {
		if firstStringFrequencies[index] != secondStringFrequencies[index] {
			return false
		}
	}
	return true
}

func RemoveDigits(input string) string {
	builder := strings.Builder{}
	for _, v := range input {
		if !(v >= 49 && v <= 57) {
			builder.WriteRune(v)
		}
	}
	return builder.String()
}

func ReplaceDigits(input string, replacement string) string {
	builder := strings.Builder{}
	for _, v := range input {
		if !(v >= 49 && v <= 57) {
			builder.WriteRune(v)
		} else {
			builder.WriteString(replacement)
		}
	}
	return builder.String()
}

type Student struct {
	id   int
	name string
}

func (student *Student) Name() string {
	return student.name
}

type Course struct {
	id   int
	name string
}

func (course *Course) Name() string {
	return course.name
}

type DataSource interface {
	ReadStudent(int) (Student, error)
	ReadCourse(int) (Course, error)
}

type Repository struct {
	students []Student
	courses  []Course
}

func (dataSource *Repository) ReadStudent(studentId int) (Student, error) {
	for _, student := range dataSource.students {
		if student.id == studentId {
			return student, nil
		}
	}
	return Student{}, errors.New("unable to find student by id")
}

func (dataSource *Repository) ReadCourse(courseId int) (Course, error) {
	for _, course := range dataSource.courses {
		if course.id == courseId {
			return course, nil
		}
	}
	return Course{}, errors.New("unable to find course by id")
}

func EnrollStudentToCourse(dataSource DataSource, studentId int, courseId int) error {
	dataSource.ReadStudent(studentId)
	dataSource.ReadCourse(courseId)
	//student, err := dataSource.ReadStudent()
	return nil
}
