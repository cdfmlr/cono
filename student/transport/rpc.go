package transport

import (
	"conostudent/endpoint"
	"conostudent/service"
	"context"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
)

// 实例化 StudentRPC 服务
var studentRPCService endpoint.StudentRPCService

// initStudentRPCService  实例化 StudentRPC 服务
func initStudentRPCService() {
	log.Info("init StudentRPCService")

	studentRPCService = endpoint.StudentRPCService{
		GetStudentBySid: func(ctx context.Context, request *endpoint.GetStudentBySidRequest) (student *endpoint.Student, err error) {
			s, err := service.GetStudentBySid(request.Sid)
			return endpoint.StudentFromModel(s), err
		},
		GetStudentByWechatID: func(ctx context.Context, request *endpoint.GetStudentByWechatIDRequest) (student *endpoint.Student, err error) {
			s, err := service.GetStudentByWechatID(request.WechatId)
			return endpoint.StudentFromModel(s), err
		},
		GetAllStudents: func(ctx context.Context, empty *endpoint.Empty) (response *endpoint.GetAllStudentsResponse, err error) {
			students, err := service.GetAllStudents()
			if err != nil {
				return
			}
			response = new(endpoint.GetAllStudentsResponse)
			response.Students = []*endpoint.Student{}
			for _, s := range students {
				response.Students = append(response.Students, endpoint.StudentFromModel(&s))
			}
			return response, err
		},
		Save: func(ctx context.Context, student *endpoint.Student) (empty *endpoint.Empty, err error) {
			s := endpoint.StudentToModel(student)
			err = service.SaveStudent(s)
			return new(endpoint.Empty), err
		},
	}
}

// Serve 在指定地址 (e.g. "localhost:8080") 启动 StudentRPC 的 gRPC 服务
func Serve(address string) {
	// 实例化 gRPC，注册服务
	s := grpc.NewServer()
	endpoint.RegisterStudentRPCService(s, &studentRPCService)

	// 👇 For debug:
	reflection.Register(s)
	// 👆 Unnecessary for production environment

	logger := log.WithField("address", address)

	// 监听网络
	listener, err := net.Listen("tcp", address)
	if err != nil {
		logger.WithField("err", err).Fatal("listen error")
	}
	logger.Info("StudentRPC Listening tcp...")

	// 启动 gRPC 服务
	err = s.Serve(listener)
	if err != nil {
		logger.WithField("err", err).Fatal("gRPC Serve error")
	}
}
