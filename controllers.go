package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// Controller for the /students endpoint
func studentsController(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	students, err := json.Marshal(getStudents())
	if err != nil {
		log.Println(err)
	}
	w.Write(students)
}

// Controller for the /student/:id endpoint
func studentsIDController(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	allScores := getAllByStudentID(id)

	scores := StudentsResponse{
		Scores:  allScores,
		Average: average(getScores(allScores))}

	var jsonData []byte
	jsonData, err := json.Marshal(scores)
	if err != nil {
		log.Println(err)
	}

	w.Write(jsonData)
}

// Controller for the /exams endpoint
func examsController(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	exams, err := json.Marshal(getExams())
	if err != nil {
		log.Println(err)
	}
	w.Write(exams)
}

// Controller for the /exams/:number endpoint
func examsNumberController(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	number := string(ps.ByName("number"))
	allScores := getAllByExamNumber(number)

	scores := ExamsResponse{
		Scores:  allScores,
		Average: average(getScores(allScores))}

	var jsonData []byte
	jsonData, err := json.Marshal(scores)
	if err != nil {
		log.Println(err)
	}

	w.Write(jsonData)
}
