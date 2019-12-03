package handler

import (
	"context"

	"github.com/micro/go-log"

	kubernetes "micros/kubernetes/proto/kubernetes"
)

type Kubernetes struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Kubernetes) Call(ctx context.Context, req *kubernetes.Request, rsp *kubernetes.Response) error {
	log.Log("Received Kubernetes.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *Kubernetes) Stream(ctx context.Context, req *kubernetes.StreamingRequest, stream kubernetes.Kubernetes_StreamStream) error {
	log.Logf("Received Kubernetes.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Logf("Responding: %d", i)
		if err := stream.Send(&kubernetes.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *Kubernetes) PingPong(ctx context.Context, stream kubernetes.Kubernetes_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Logf("Got ping %v", req.Stroke)
		if err := stream.Send(&kubernetes.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
