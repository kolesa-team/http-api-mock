package notify

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/kolesa-team/http-api-mock/definition"
	"github.com/kolesa-team/http-api-mock/logging"
	"github.com/kolesa-team/http-api-mock/utils"
)

//RequestCaller makes remote http requests
type RequestCaller struct {
}

//Call makes a remote http request
func (caller RequestCaller) Call(request definition.Request) bool {

	requestURL, err := url.Parse(request.Path)
	if err != nil {
		logging.Printf("Invalid url(%s) passed: %s", request.Path, err)
		return false
	}

	if !requestURL.IsAbs() {
		request.Path = utils.GetServerAddress() + "/" + strings.TrimPrefix(request.Path, "/")
	}

	req, err := http.NewRequest(request.Method, request.Path, bytes.NewBufferString(request.Body))
	if err != nil {
		logging.Printf("Error creating http request: %s", err)
		return false
	}

	for header, values := range request.Headers {
		for _, value := range values {
			req.Header.Add(header, value)
		}
	}

	cookies := []string{}
	for cookie, value := range request.Cookies {
		cookies = append(cookies, fmt.Sprintf("%s=%s", cookie, value))
		req.Header.Add("Set-Cookie", strings.Join(cookies, ";"))
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		logging.Printf("Error executing request to %s. Error: %s", request.Path, err)
		return false
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	logging.Printf("Request to %s returned status code %d and body: %s", request.Path, resp.StatusCode, body)
	return true
}
