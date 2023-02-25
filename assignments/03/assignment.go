package assignment_03

import (
	"strings"
)

type Student struct {
	id   int
	name string
}

type Course struct {
	id       int
	name     string
	students []Student
}

func (course *Course) Name() string {
	return course.name
}

func (course *Course) EnrollStudent(s Student) error {
	course.students = append(course.students, s)
	return nil
}

type DataSource interface{}

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
	return false
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

func EnrollStudentToCourse(dataSource DataSource, studentId int, courseId int) error {
	return nil
}
