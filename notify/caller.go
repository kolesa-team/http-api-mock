package notify

import "github.com/kolesa-team/http-api-mock/definition"

//Caller makes remote http requests
type Caller interface {
	//Call makes a remote http request
	Call(m definition.Request) bool
}
