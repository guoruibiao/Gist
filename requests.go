package gist

import (
		"net/http"
	"bytes"
	"encoding/json"
	"io/ioutil"
)


func NewRequest(method, url string) *Request {
	return &Request{
		method:method,
		url:url,
		headers: map[string]string{},
	}
}

func (request *Request) Token(token string) *Request {
	if token != "" {
		request.token = token
	}
	return request
}

func (request *Request) Headers(headers map[string]string) *Request {
	if headers != nil {
		for key, value := range headers {
			request.headers[key] = value
		}
	}
	return request
}

func (request *Request) Body(payload interface{}) *Request {
	request.body = payload
	return request
}

func (request *Request) DoRequest() *Response {
	body := bytes.NewBuffer(nil)
	if request.body != nil {
		err := json.NewEncoder(body).Encode(request.body)

		if err != nil {
			return &Response{resp: nil, err: err}
		}
	}
	req, err := http.NewRequest(request.method, request.url, body)
	if err != nil {
		return &Response{resp:nil, err:err}
	}
	/*
	if request.token != "" {
		req.Header.Add("Authorization", "Token "+request.token)
	}
	*/
	if request.headers != nil {
		for key, value := range request.headers {
			req.Header.Add(key, value)
		}
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return &Response{resp:nil, err:err}
	}
	return &Response{resp:resp, err:nil}
}


func (response *Response)GetResponse() (string, error){
	bytes, err := ioutil.ReadAll(response.resp.Body)
	if err != nil {
		return "", err
	}
	defer response.resp.Body.Close()
	return string(bytes), nil
}
