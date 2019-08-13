package gist

import (
	"encoding/json"
	)

func NewGistManager()*GistManager{
	return &GistManager{}
}

func (manager *GistManager) CreateGist(gist *Gist) (bool, error) {
	headers := map[string]string{
		"Authorization": "Token "+ GIST_OAUTH_TOKEN,
	}
	resp := NewRequest("POST", GIST_API_URL).Headers(headers).Body(gist).DoRequest()
	if resp.err != nil {
		return false, resp.err
	}
	return true, nil
}

func (manager *GistManager)UpdateGist(gist *Gist) (bool, error) {
	headers := map[string]string{
		"Authorization": "Token "+ GIST_OAUTH_TOKEN,
	}
	url := GIST_API_URL + "/" + gist.Id
	resp := NewRequest("POST", url).Headers(headers).Body(gist).DoRequest()
	if resp.err != nil {
		return false, resp.err
	}
	return true, nil
}

func (manager *GistManager)DeleteGist(id string) (bool, error) {
	headers := map[string]string{
		"Authorization": "Token "+ GIST_OAUTH_TOKEN,
	}
	url := GIST_API_URL  + "/" + id
	resp := NewRequest("DELETE", url).Headers(headers).DoRequest()
	if resp.err != nil {
		return false, resp.err
	}
	return true, nil
}

func (manager *GistManager)ShowGist(id string) (*Gist, error) {
	headers := map[string]string{
		"Authorization": "Token "+ GIST_OAUTH_TOKEN,
	}
	url := GIST_API_URL + "/" + id
	response := NewRequest("GET", url).Headers(headers).DoRequest()
	if response.err != nil {
		return nil, response.err
	}
	defer response.resp.Body.Close()
	gist := &Gist{}
	json.NewDecoder(response.resp.Body).Decode(gist)
	return gist, nil
}

func (manager *GistManager) ListGists() ([]*Gist, error) {
	headers := map[string]string{
		"Authorization": "Token "+ GIST_OAUTH_TOKEN,
	}
	response := NewRequest("GET", GIST_API_URL).Headers(headers).DoRequest()
	if response.err != nil {
		return nil, response.err
	}
	defer response.resp.Body.Close()
	var gists []*Gist
	json.NewDecoder(response.resp.Body).Decode(&gists)
	return gists, nil
}