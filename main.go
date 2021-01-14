package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {

	// Runs the live server
	go func() {
		router := httprouter.New()

		router.GET("/students", studentsController)
		router.GET("/students/:id", studentsIDController)
		router.GET("/exams", examsController)
		router.GET("/exams/:number", examsNumberController)

		log.Fatal(http.ListenAndServe(":8090", router))
	}()

	resp, err := http.Get("http://live-test-scores.herokuapp.com/scores")
	fmt.Println("Response: Read", resp, "bytes", err)

	if err != nil {
		log.Fatal(err)
	}

	reader := bufio.NewReader(resp.Body)

	for {
		reader.ReadBytes('\n')
		reader.ReadBytes('\n')
		line, _ := reader.ReadBytes('\n')

		data := string(line)
		data = data[6:len(data)]

		ServerEvent := ServerEvent{}
		err := ServerEvent.UnmarshalJSON([]byte(data))
		if err != nil {
			fmt.Printf("Error %+v", err)
		}

		addStudents(ServerEvent)
		addExams(ServerEvent)
	}
}
