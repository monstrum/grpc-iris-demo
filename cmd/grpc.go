package cmd

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/monstrum/grpc-iris-demo/internal"
	product "github.com/monstrum/grpc-iris-demo/pkg/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

func CreateGrpcServer() Command {
	return &grpcServer{}
}

type grpcServer struct {
}

func (s *grpcServer) Execute() error {
	handler := internal.CreateGrpcHandler()
	serv := grpc.NewServer()

	// I am end up serving grpc as usually
	// go s.serveGrpc(serv, handler)
	return s.serveIris(serv, handler)
}

func (s *grpcServer) serveGrpc(serv *grpc.Server, handler product.ApiServer) {
	lis, _ := net.Listen("tcp", "0.0.0.0:50051")
	product.RegisterApiServer(serv, handler)

	log.Printf("Now listening on: %s", "0.0.0.0:50051")
	err := serv.Serve(lis)
	if err != nil {
		log.Printf("grpc server didn't start. error: %s", err)
	}
}

func (s *grpcServer) serveIris(serv *grpc.Server, handler product.ApiServer) error {
	app := iris.New()
	app.Logger().SetLevel("debug")
	mvc.New(app).
		Handle(handler, mvc.GRPC{
			Server:      serv,          // Required.
			ServiceName: "product.Api", // Required.
			Strict:      false,
		})
	return app.Run(
		iris.TLS(":3443", "config/server.crt", "config/server.key"),
	)
}
