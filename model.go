package gist

import (
	"time"
	"net/http"
)

type (
	// gist CRUD
	Gist struct {
		Id string `json:"id, omitempty"`
		Description string `json:"description, omitempty"`
		Public bool `json:"public, omitempty"`
		Files map[string]GistFile `json:"files, omitempty"`
		HTMLURL string `json:"html_url,omitempty"`
		UpdateAt time.Time `json:"update_at,omitempty"`
	}
	GistFile struct {
		Content string `json:"content,omitempty"`
	}
	// request & response
	Request struct {
		method string
		url string
		token string
		headers map[string]string
		body interface{}
	}
	Response struct {
		resp *http.Response
		err error
	}
	// gist manager
	GistManager struct{
		TOKEN string
	}
)

const GIST_OAUTH_TOKEN = "your gist oauth token here."
const GIST_API_URL = "https://api.github.com/gists"