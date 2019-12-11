package amqp

import "github.com/kolesa-team/http-api-mock/definition"

//Sender sends messages to AMQP server
type Sender interface {
	//Send sends to amqp
	Send(m *definition.Mock) bool
}
