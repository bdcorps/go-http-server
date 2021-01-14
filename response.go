package main

// StudentsResponse Common response for the app endpoints
type StudentsResponse struct {
	Scores  []ServerEvent
	Average float64
}

// ExamsResponse Common response for the app endpoints
type ExamsResponse struct {
	Scores  []ServerEvent
	Average float64
}
