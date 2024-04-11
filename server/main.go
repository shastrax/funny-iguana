package main

import (
	"context"
	"fmt"
	"net"

	v1 "github.com/shastrax/funny-iguana/gen/iguana/v1"
	//	"github.com/shastrax/funny-iguana/gen/iguana/v1/iguanav1connect"
	"google.golang.org/grpc"

	"github.com/bufbuild/connect-go"
)

// buf curl --schema proto --data '{"source":"bogus"}' http://localhost:8080/iguana.v1.IguanaService/Ping

const address = "localhost:8080"

func main() {
	/*
		path, handler := iguanav1connect.NewIguanaServiceHandler(&iguanaServiceServer{})

		mux := http.NewServeMux()
		mux.Handle(path, handler)

		fmt.Println("... Listening on", address)
		http.ListenAndServe(
			address,
			// Use h2c so we can serve HTTP/2 without TLS.
			h2c.NewHandler(mux, &http2.Server{}),
		)
	*/

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", 50053))
	if err != nil {
		fmt.Println(err)
		//at.SugarLog.Fatalf("failed to listen: %v", err)
	}

	//st := ServerType{SugarLog: at.SugarLog}
	//st := iguanaServiceServer{}

	grpcServer := grpc.NewServer()
	//x v1.RegisterIguanaServiceServer(grpcServer, &st)
	//	pb.RegisterMixerServer(grpcServer, &st)
	//	at.SugarLog.Infof("server listening at %v", listener.Addr())

	if err := grpcServer.Serve(listener); err != nil {
		fmt.Println(err)
		//			at.SugarLog.Fatalf("failed to serve: %v", err)
	}
}

type iguanaServiceServer struct {
	//x iguanav1connect.UnimplementedIguanaServiceHandler
}

func (ss *iguanaServiceServer) Ping(ctx context.Context, req *connect.Request[v1.PingRequest]) (*connect.Response[v1.PingResponse], error) {
	fmt.Println("ping ping ping")

	//fmt.Println(req.Msg)

	source := req.Msg.GetSource()
	fmt.Println("ping request source:", source)

	return connect.NewResponse(&v1.PingResponse{Status: "happy"}), nil
}
