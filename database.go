package main

var students = make(map[string][]float64)
var exams = make(map[string][]float64)

// Returns all the seen students
func getStudents() []string {
	return getKeys(students)
}

// Returns all the scores by a particular student
func getAllScoresByStudentID(id string) []float64 {
	return students[id]
}

// Add a new student's scores
func addStudents(studentId string, score float64) {
	students[studentId] = append(students[studentId], score)
}

// Returns all the seen exams
func getExams() []string {
	return getKeys(exams)
}

// Returns all scores for a given exam
func getAllScoresByExamNumber(number string) []float64 {
	return exams[number]
}

// Add a new exam's scores
func addExams(examNo string, score float64) {
	exams[examNo] = append(exams[examNo], score)
}
