package subscriber

import (
	"context"
	"github.com/micro/go-log"

	kubernetes "micros/kubernetes/proto/kubernetes"
)

type Kubernetes struct{}

func (e *Kubernetes) Handle(ctx context.Context, msg *kubernetes.Message) error {
	log.Log("Handler Received message: ", msg.Say)
	return nil
}

func Handler(ctx context.Context, msg *kubernetes.Message) error {
	log.Log("Function Received message: ", msg.Say)
	return nil
}
