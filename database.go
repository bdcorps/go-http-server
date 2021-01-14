package main

var students = make(map[string][]ServerEvent)
var exams = make(map[string][]ServerEvent)

// Returns all the seen students
func getStudents() []string {
	return getKeys(students)
}

// Returns all the scores by a particular student
func getAllByStudentID(id string) []ServerEvent {
	return students[id]
}

// Add a new student's scores
func addStudents(serverEvent ServerEvent) {
	students[serverEvent.StudentId] = append(students[serverEvent.StudentId], serverEvent)
}

// Returns all the seen exams
func getExams() []string {
	return getKeys(exams)
}

// Returns all scores for a given exam
func getAllByExamNumber(number string) []ServerEvent {
	return exams[number]
}

// Add a new exam's scores
func addExams(serverEvent ServerEvent) {
	exams[serverEvent.Exam] = append(exams[serverEvent.Exam], serverEvent)
}
