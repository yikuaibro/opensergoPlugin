package stream_plugin

import (
	"context"
	"fmt"

	"github.com/hashicorp/go-plugin"
	pb "github.com/opensergo/opensergo-control-plane/pkg/plugin/proto/stream"
	"google.golang.org/grpc"
)

type StreamPluginServer struct {
	pb.UnimplementedStreamGreeterServer
}

var _ pb.StreamGreeterServer = (*StreamPluginServer)(nil)

func (s StreamPluginServer) Greet(ctx context.Context, req *pb.StreamReq) (*pb.StreamResp, error) {
	fmt.Println("GreetGreetGreet")
	return &pb.StreamResp{
		Greet: "testtest",
	}, nil
}

type StreamPlugin struct {
	plugin.Plugin

	impl pb.StreamGreeterServer
}

func NewStreamPluginServiceServer(impl pb.StreamGreeterServer) (*StreamPlugin, error) {
	if impl == nil {
		return nil, fmt.Errorf("empty underlying stream plugin passed in")
	}
	return &StreamPlugin{
		impl: impl,
	}, nil
}

func (h *StreamPlugin) GRPCServer(broker *plugin.GRPCBroker, s *grpc.Server) error {
	pb.RegisterStreamGreeterServer(s, h.impl)
	return nil
}

func (h *StreamPlugin) GRPCClient(ctx context.Context, broker *plugin.GRPCBroker, c *grpc.ClientConn) (any, error) {
	return pb.NewStreamGreeterClient(c), nil
}
