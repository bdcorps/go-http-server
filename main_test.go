package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddStudents(t *testing.T) {
	event1 := ServerEvent{"Student 1", "1", .10}
	event2 := ServerEvent{"Student 1", "2", .20}
	event3 := ServerEvent{"Student 2", "2", .40}
	addStudents(event1)
	addStudents(event2)
	addStudents(event3)

	expected := []string{"Student 1", "Student 2"}
	actual := getStudents()
	assert.Equal(t, actual, expected, "There should be two students saved")
}

func TestGetAllByStudentID(t *testing.T) {
	expected := []float64{.10, .20}
	actual := getScores(getAllByStudentID("Student 1"))
	assert.Equal(t, actual, expected, "There should be scores for Student 1")
}

func TestStudentAverage(t *testing.T) {
	allScores := getScores(getAllByStudentID("Student 1"))
	expected := .15
	actual := round(average(allScores))
	assert.Equal(t, actual, expected, "There should be two scores for Student 1")
}

func TestAddExams(t *testing.T) {
	event1 := ServerEvent{"Student 3", "3", .10}
	event2 := ServerEvent{"Student 3", "4", .20}
	event3 := ServerEvent{"Student 4", "4", .40}
	addExams(event1)
	addExams(event2)
	addExams(event3)

	expected := []string{"3", "4"}
	actual := getExams()
	assert.Equal(t, actual, expected, "There should be two students saved")
}

func TestGetAllScoresByExamID(t *testing.T) {
	expected := []float64{.20, .40}
	actual := getScores(getAllByExamNumber("4"))
	assert.Equal(t, actual, expected, "There should be scores for Exam 1")
}

func TestExamAverage(t *testing.T) {
	allScores := getScores(getAllByExamNumber("4"))
	expected := .30
	actual := round(average(allScores))
	assert.Equal(t, actual, expected, "There should be two scores for Exam 1")
}
