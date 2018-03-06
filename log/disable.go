package main

import (
	"log"
	"os"
	"io/ioutil"
)

func main() {
	logger := log.New(os.Stdout, "", 0)
	quiet := log.New(ioutil.Discard, "", 0)
	// use ioutil.Discard for a quick and dirty. For: perf. concerns, see:
	// https://stackoverflow.com/questions/10571182/

	logger.Println("Hello world")
	quiet.Println("Can you hear me?")
}
