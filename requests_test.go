package gist

import "testing"

func TestRequest_DoRequest(t *testing.T) {
	payload := &map[string]string{
		"hello": "world",
		"name": "tiger",
	}
	headers := map[string]string{
		"Authorization": "Token "+ GIST_OAUTH_TOKEN,
	}
	response := NewRequest("POST", "http://www.01happy.com/demo/accept.php").Headers(headers).Body(payload).DoRequest()
	t.Log(response.resp)
}
