package grpcagent

import (
	"github.com/golang/protobuf/proto"
	"google.golang.org/grpc/metadata"
)

type Agent struct {
	client  interface{}
	handler func(interface{}, string, interface{}, metadata.MD) (proto.Message, error)
}

func (c *Agent) Request(method string, data interface{}, header metadata.MD) (proto.Message, error) {
	return c.handler(c.client, method, data, header)
}

func (c *Agent) Client() interface{} {
	return c.client
}

func CreateAgent(client interface{}, handler func(interface{}, string, interface{}, metadata.MD) (proto.Message, error)) *Agent {
	return &Agent{
		client,
		handler,
	}
}
