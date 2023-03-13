package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {

	log := log.New(newFileLogger(), "[S3]", 1)

	res := Result{
		value:   "qq",
		isError: false,
	}

	res.setRandomValue()

	str := res.ToString()

	log.Println(str)

	// r := mux.NewRouter()
	// routes.RegisterBookStoreRoutes(r)
	// http.Handle("/", r)
	// log.Fatal(http.ListenAndServe("localhost:9091", r))
}

type FileLogger struct{}

func newFileLogger() io.Writer {
	return &FileLogger{}
}

func (fl *FileLogger) Write(p []byte) (n int, err error) {
	err = os.WriteFile("./test.log", p, 0644)
	n = len(p)
	return
}

type Result struct {
	value   string
	isError bool
}

func (r Result) ToString() string {
	return fmt.Sprintf("Value is %s", r.value)
}

func (r *Result) setRandomValue() {
	r.value = "cc"
}
