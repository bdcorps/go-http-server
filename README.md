# Assignment

## Installation

```
go get github.com/stretchr/testify/assert
go get github.com/julienschmidt/httprouter

go run .*
```

### Available Endpoints

- [localhost:8090/students](localhost:8090/students)
- [localhost:8090/students/Juanita65](localhost:8090/students/:id)
- [localhost:8090/exams](localhost:8090/exams)
- [localhost:8090/exams/45](localhost:8090/exams/:id)

## Run Unit Tests

```
go test
```

## Notes

- The data server needs to be woken up
- Better to send JSON data than mixed content
- `exam` should be string and not int

# go-http-server

## Assumptions

- No duplicate entries
