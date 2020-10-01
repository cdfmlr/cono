package transport

import (
	"conostudent/config"
	"conostudent/model"
	"conostudent/service"
	"testing"
)

func TestServe(t *testing.T) {
	// Setup
	config.Init("/Users/c/Desktop/StudentConf.yml")
	model.Init()
	service.Init()
	Init()

	Serve(config.Serve.StudentRPCAddress)

	/*
		gRPC Test & Debug with grpcui, which requests:
		    reflection.Register(s)
		in func Serve.

		Install grpcui:
		    $ go get github.com/fullstorydev/grpcui
		    $ go install github.com/fullstorydev/grpcui/cmd/grpcui
		Start grpcui:
		    $ grpcui -plaintext localhost:8080
	*/
}
