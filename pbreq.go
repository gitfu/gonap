package gonap

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

//fullheader is the standard header to include with all http requests except is_patch and is_command
const fullheader = "application/vnd.profitbricks.resource+json"

//patchheader is used with is_patch .
const patchheader = "application/vnd.profitbricks.partial-properties+json"

//commandheader is used with is_command
const commandheader = "application/x-www-form-urlencoded"

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
		return path
	}
	url := Endpoint + path
	return url
}

// is_delete performs an http.NewRequest DELETE  and returns a PBResp struct
func is_delete(path string) PBResp {
	url := mk_url(path)
	client := &http.Client{}
	req, _ := http.NewRequest("DELETE", url, nil)
	req.SetBasicAuth(Username, Passwd)
	req.Header.Add("Content-Type", fullheader)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)

	}
	defer resp.Body.Close()
	resp_body, _ := ioutil.ReadAll(resp.Body)
	var R PBResp
	R.Body = resp_body
	R.Headers = resp.Header
	R.StatusCode = resp.StatusCode
	return R

}

// is_get performs an http.NewRequest GET and returns a PBResp struct
func is_get(path string) PBResp {
	url := mk_url(path) + `?depth=` + Depth
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.SetBasicAuth(Username, Passwd)
	req.Header.Add("Content-Type", fullheader)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)

	}
	defer resp.Body.Close()
	resp_body, _ := ioutil.ReadAll(resp.Body)
	var R PBResp
	R.Body = resp_body
	R.Headers = resp.Header
	R.StatusCode = resp.StatusCode
	fmt.Printf("%s", R.Body)
	return R

}

// is_patch performs an http.NewRequest PATCH and returns a PBResp struct
func is_patch(path string, jason []byte) PBResp {
	url := mk_url(path)
	client := &http.Client{}
	req, _ := http.NewRequest("PATCH", url, bytes.NewBuffer(jason))
	req.SetBasicAuth(Username, Passwd)
	req.Header.Add("Content-Type", patchheader)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	resp_body, _ := ioutil.ReadAll(resp.Body)
	var R PBResp
	R.Body = resp_body
	R.Headers = resp.Header
	R.StatusCode = resp.StatusCode
	return R

}

// is_put performs an http.NewRequest PUT and returns a PBResp struct
func is_put(path string, jason []byte) PBResp {
	url := mk_url(path)
	client := &http.Client{}
	req, _ := http.NewRequest("PUT", url, bytes.NewBuffer(jason))
	req.SetBasicAuth(Username, Passwd)
	req.Header.Add("Content-Type", fullheader)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)

	}
	defer resp.Body.Close()
	resp_body, _ := ioutil.ReadAll(resp.Body)
	var R PBResp
	R.Body = resp_body
	R.Headers = resp.Header
	R.StatusCode = resp.StatusCode
	return R

}

// is_post performs an http.NewRequest POST and returns a PBResp struct
func is_post(path string, jason []byte) PBResp {
	url := mk_url(path)
	client := &http.Client{}
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jason))
	req.SetBasicAuth(Username, Passwd)
	req.Header.Add("Content-Type", fullheader)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)

	}
	defer resp.Body.Close()
	resp_body, _ := ioutil.ReadAll(resp.Body)
	var R PBResp
	R.Body = resp_body
	R.Headers = resp.Header
	R.StatusCode = resp.StatusCode
	return R

}

// is_command performs an http.NewRequest PUT and returns a PBResp struct
func is_command(path string, jason string) PBResp {
	url := mk_url(path)
	body := json.RawMessage(jason)
	client := &http.Client{}
	req, _ := http.NewRequest("PUT", url, bytes.NewBuffer(body))
	req.SetBasicAuth(Username, Passwd)
	req.Header.Add("Content-Type", commandheader)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)

	}
	defer resp.Body.Close()
	resp_body, _ := ioutil.ReadAll(resp.Body)
	var R PBResp
	R.Body = resp_body
	R.Headers = resp.Header
	R.StatusCode = resp.StatusCode
	return R

}
