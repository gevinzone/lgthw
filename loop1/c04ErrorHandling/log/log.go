package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
)

func DefaultLog() {
	log.Println("DefaultLog")
}

func StdLog() {
	logger := log.Logger{}
	logger.SetOutput(os.Stdout)
	logger.SetFlags(log.LstdFlags)
	logger.Println("StdLog")
}

func Logger() {
	buf := bytes.Buffer{}
	logger := log.New(&buf, "logger-", log.Lshortfile|log.Ldate)

	logger.Println("Logger")
	logger.SetPrefix("new logger-")
	logger.Printf("you can also add args(%v) and use Fatalln to log log and crash", true)

	fmt.Println(buf.String())
}
