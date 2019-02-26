package grpcagent

import "github.com/golang/protobuf/proto"

type Agent struct {
	client  interface{}
	handler func(interface{}, string, interface{}) (proto.Message, error)
}

func (c *Agent) Request(method string, data interface{}) (proto.Message, error) {
	return c.handler(c.client, method, data)
}

func (c *Agent) Client() interface{} {
	return c.client
}

func CreateAgent(client interface{}, handler func(interface{}, string, interface{}) (proto.Message, error)) *Agent {
	return &Agent{
		client,
		handler,
	}
}
