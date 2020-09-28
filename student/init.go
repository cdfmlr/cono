package main

import log "github.com/sirupsen/logrus"

func init() {
	// For production environment:
	log.SetFormatter(&log.JSONFormatter{}) // log in JSON
	log.SetReportCaller(true)              // include the calling method as a field.
}
