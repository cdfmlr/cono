package transport

import (
	"conocourse/config"
	"conocourse/endpoint"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

var StudentRPCClient endpoint.StudentRPCClient

func initStudentRPC() {
	log.Info("init StudentRPC")

	cliConn, err := grpc.Dial(*config.StudentRPCAddress, grpc.WithInsecure())
	if err != nil {
		log.Error("StudentRPC Dial error:", err)
	}
	//defer cliConn.Close()

	StudentRPCClient = endpoint.NewStudentRPCClient(cliConn)
}
