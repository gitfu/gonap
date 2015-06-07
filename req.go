package gonap

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

//FullHeader is the standard header to include with all http requests except is_patch and is_command
const FullHeader = "application/vnd.profitbricks.resource+json"

//PatchHeader is used with is_patch .
const PatchHeader = "application/vnd.profitbricks.partial-properties+json"

//CommandHeader is used with is_command
const CommandHeader = "application/x-www-form-urlencoded"

// Depth sets the level of detail returned from the REST server .
var Depth = "5"

// SetDepth is used to set Depth
func SetDepth(newdepth string) string {
	Depth = newdepth
	return Depth
}

// SetEnpoint is used to set the REST Endpoint. Endpoint is declared in config.go
func SetEndpoint(newendpoint string) string {
	Endpoint = newendpoint
	return Endpoint
}

// SetAuth is used to set Username and Passwd. Username and Passwd are declared in config.go

func SetAuth(u, p string) {
	Username = u
	Passwd = p
}

// mk_url  either:
// returns the path (if it`s a full url)
//			 or
//	returns	Endpoint+ path .
func mk_url(path string) string {
	if strings.HasPrefix(path, "http") {
		//REMOVE AFTER TESTING
		path := strings.Replace(path, "https://api.profitbricks.com/rest", Endpoint, 1)
		return path
	}
	if strings.HasPrefix(path, "<base>") {
		//REMOVE AFTER TESTING
		path := strings.Replace(path, "<base>", Endpoint, 1)
		return path
	}
	url := Endpoint + path
	return url
}

func do(req *http.Request) PBResp {
	client := &http.Client{}
	req.SetBasicAuth(Username, Passwd)
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	resp_body, _ := ioutil.ReadAll(resp.Body)
	var R PBResp
	R.Req = resp.Request
	R.Body = resp_body
	R.Headers = resp.Header
	R.StatusCode = resp.StatusCode
	return R
}

// is_delete performs an http.NewRequest DELETE  and returns a PBResp struct
func is_delete(path string) PBResp {
	url := mk_url(path)
	req, _ := http.NewRequest("DELETE", url, nil)
	req.Header.Add("Content-Type", FullHeader)
	return do(req)
}

// is_get performs an http.NewRequest GET and returns a PBResp struct
func is_get(path string) PBResp {
	url := mk_url(path) + `?depth=` + Depth
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Content-Type", FullHeader)
	return do(req)
}

// is_patch performs an http.NewRequest PATCH and returns a PBResp struct
func is_patch(path string, jason []byte) PBResp {
	url := mk_url(path)
	req, _ := http.NewRequest("PATCH", url, bytes.NewBuffer(jason))
	req.Header.Add("Content-Type", PatchHeader)
	return do(req)
}

// is_put performs an http.NewRequest PUT and returns a PBResp struct
func is_put(path string, jason []byte) PBResp {
	url := mk_url(path)
	req, _ := http.NewRequest("PUT", url, bytes.NewBuffer(jason))
	req.Header.Add("Content-Type", FullHeader)
	return do(req)
}

// is_post performs an http.NewRequest POST and returns a PBResp struct
func is_post(path string, jason []byte) PBResp {
	url := mk_url(path)
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jason))
	req.Header.Add("Content-Type", FullHeader)
	return do(req)
}

// is_command performs an http.NewRequest PUT and returns a PBResp struct
func is_command(path string, jason string) PBResp {
	url := mk_url(path)
	body := json.RawMessage(jason)
	req, _ := http.NewRequest("PUT", url, bytes.NewBuffer(body))
	req.Header.Add("Content-Type", CommandHeader)
	return do(req)
}
