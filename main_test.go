package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddStudents(t *testing.T) {
	addStudents("Student 1", .10)
	addStudents("Student 1", .20)
	addStudents("Student 2", .40)
	expected := []string{"Student 1", "Student 2"}
	actual := getStudents()
	assert.Equal(t, actual, expected, "There should be two students saved")
}

func TestGetAllScoresByStudentID(t *testing.T) {
	expected := []float64{.10, .20}
	actual := getAllScoresByStudentID("Student 1")
	assert.Equal(t, actual, expected, "There should be scores for Student 1")
}

func TestStudentAverage(t *testing.T) {
	allScores := getAllScoresByStudentID("Student 1")
	expected := .15
	actual := round(average(allScores))
	assert.Equal(t, actual, expected, "There should be two scores for Student 1")
}

func TestAddExams(t *testing.T) {
	addExams("1", .10)
	addExams("1", .20)
	addExams("2", .40)
	expected := []string{"1", "2"}
	actual := getExams()
	assert.Equal(t, actual, expected, "There should be two students saved")
}

func TestGetAllScoresByExamID(t *testing.T) {
	expected := []float64{.10, .20}
	actual := getAllScoresByExamNumber("1")
	assert.Equal(t, actual, expected, "There should be scores for Exam 1")
}

func TestExamAverage(t *testing.T) {
	allScores := getAllScoresByExamNumber("1")
	expected := .15
	actual := round(average(allScores))
	assert.Equal(t, actual, expected, "There should be two  scores for Exam 1")
}
