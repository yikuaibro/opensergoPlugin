package main

import (
	"context"
	"log"

	"github.com/opensergo/opensergo-control-plane/pkg/plugin/pl"
	"github.com/opensergo/opensergo-control-plane/pkg/plugin/pl/builtin"
	pb "github.com/opensergo/opensergo-control-plane/pkg/plugin/proto/stream"
)

//nolint:gosimple
func main() {
	pluginServer := pl.NewPluginServer()
	//defer func() {
	//	if err := pluginServer.RunShutdownFuncs(); err != nil {
	//		log.Fatalln("Error:", err.Error())
	//	}
	//	log.Println("Server shutdown")
	//	return
	//}()
	err := pluginServer.InitPlugin()
	if err != nil {
		log.Fatalln("Error:", err.Error())
	}

	client, err := pluginServer.GetPluginClient(builtin.StreamServicePluginName)
	if err != nil {
		log.Fatalln("Error:", err.Error())
	}
	raw, ok := client.(pb.StreamGreeterClient)
	if !ok {
		log.Fatalln("Error: can't convert rpc plugin to normal wrapper")
	}

	ctx := context.Background()
	greet, err := raw.Greet(ctx, &pb.StreamReq{})
	if err != nil {
		log.Printf("Error: %s\n", err.Error())
	}
	log.Println("Greet:", greet)
	for {
		select {
		case <-pluginServer.ShutdownCh:
			pluginServer.ContextCancel()
			if err := pluginServer.RunShutdownFuncs(); err != nil {
				log.Fatalln("Error:", err.Error())
			}
			log.Println("Server shutdown")
			return
		}
	}

}
