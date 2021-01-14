package main

import (
	"encoding/json"
	"math"
	"strconv"
)

// Returns keys only for given map
func getKeys(dataMap map[string][]float64) []string {
	keys := make([]string, len(dataMap))

	i := 0
	for k := range dataMap {
		keys[i] = k
		i++
	}
	return keys
}

// Rounds to  decimal places
func round(val float64) float64 {
	return math.Round(val*100) / 100
}

//Averages a given array
func average(arr []float64) float64 {
	return sum(arr) / float64(len(arr))
}

// Sums the given array
func sum(arr []float64) float64 {
	if len(arr) == 0 {
		return 0.0
	}

	result := 0.0
	for _, v := range arr {
		result += v
	}
	return result
}

// Custom marshaller for events coming from server-side - making exam a string
func (s *ServerEvent) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}

	var StudentId string
	err = json.Unmarshal(*objMap["studentId"], &StudentId)
	if err != nil {
		return err
	}

	var Score float64
	err = json.Unmarshal(*objMap["score"], &Score)
	if err != nil {
		return err
	}

	var ExamString string
	err = json.Unmarshal(*objMap["exam"], &ExamString)
	if err != nil {
		var examInt int
		err = json.Unmarshal(*objMap["exam"], &examInt)
		if err != nil {
			return err
		}
		exam := strconv.Itoa(examInt)
		if err != nil {
			return err
		}
		s.Exam = exam

	} else {
		s.Exam = ExamString
	}

	s.StudentId = StudentId
	s.Score = Score

	return nil
}
