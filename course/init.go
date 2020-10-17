package main

import (
	"github.com/sirupsen/logrus"
	"log"
)

func init() {
	// go log
	log.SetPrefix("[conocourse] ")

	// logrus
	// DONE: uncomment the following lines for production environment:
	logrus.SetFormatter(&logrus.JSONFormatter{}) // log in JSON
	logrus.SetReportCaller(true)                 // include the calling method as a field.
}
