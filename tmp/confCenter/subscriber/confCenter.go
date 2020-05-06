package subscriber

import (
	"context"
	"github.com/micro/go-micro/util/log"

	confCenter "confCenter/proto/confCenter"
)

type ConfCenter struct{}

func (e *ConfCenter) Handle(ctx context.Context, msg *confCenter.Message) error {
	log.Log("Handler Received message: ", msg.Say)
	return nil
}

func Handler(ctx context.Context, msg *confCenter.Message) error {
	log.Log("Function Received message: ", msg.Say)
	return nil
}
