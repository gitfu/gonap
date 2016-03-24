package goprofitbricks

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

//FullHeader is the standard header to include with all http requests except isPatch and is_command
const FullHeader = "application/vnd.profitbricks.resource+json"

//PatchHeader is used with isPatch .
const PatchHeader = "application/vnd.profitbricks.partial-properties+json"

//CommandHeader is used with is_command
const CommandHeader = "application/x-www-form-urlencoded"

var Depth = "5"

// SetDepth is used to set Depth
func SetDepth(newdepth string) string {
	Depth = newdepth
	return Depth
}

// mkURL  either:
// returns the path (if it`s a full url)
//			 or
//	returns	Endpoint+ path .
func mkURL(path string) string {
	if strings.HasPrefix(path, "http") {
		//REMOVE AFTER TESTING
		path := strings.Replace(path, "https://api.profitbricks.com/rest", Endpoint, 1)
		// END REMOVE
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

func do(req *http.Request) Resp {
	client := &http.Client{}
	req.SetBasicAuth(Username, Passwd)
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	resp_body, _ := ioutil.ReadAll(resp.Body)
	var R Resp
	R.Req = resp.Request
	R.Body = resp_body
	R.Headers = resp.Header
	R.StatusCode = resp.StatusCode
	return R
}

// isDelete performs an http.NewRequest DELETE  and returns a Resp struct
func isDelete(path string) Resp {
	url := mkURL(path)
	req, _ := http.NewRequest("DELETE", url, nil)
	req.Header.Add("Content-Type", FullHeader)
	return do(req)
}

// isList performs an http.NewRequest GET and returns a Collection struct
func isList(path string) Collection {
	url := mkURL(path) + `?depth=` + Depth
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Content-Type", FullHeader)
	return toCollection(do(req))
}

// isGet performs an http.NewRequest GET and returns an Instance struct
func isGet(path string) Instance {
	url := mkURL(path) + `?depth=` + Depth
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Content-Type", FullHeader)
	return toInstance(do(req))
}

// isPatch performs an http.NewRequest PATCH and returns an Instance struct
func isPatch(path string, jason []byte) Instance {
	url := mkURL(path)
	req, _ := http.NewRequest("PATCH", url, bytes.NewBuffer(jason))
	req.Header.Add("Content-Type", PatchHeader)
	return toInstance(do(req))
}

// isPut performs an http.NewRequest PUT and returns an Instance struct
func isPut(path string, jason []byte) Instance {
	url := mkURL(path)
	req, _ := http.NewRequest("PUT", url, bytes.NewBuffer(jason))
	req.Header.Add("Content-Type", FullHeader)
	return toInstance(do(req))
}

// is_post performs an http.NewRequest POST and returns an Instance struct
func is_post(path string, jason []byte) Instance {
	url := mkURL(path)
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jason))
	req.Header.Add("Content-Type", FullHeader)
	return toInstance(do(req))
}

// is_command performs an http.NewRequest PUT and returns a Resp struct
func is_command(path string, jason string) Resp {
	url := mkURL(path)
	body := json.RawMessage(jason)
	req, _ := http.NewRequest("PUT", url, bytes.NewBuffer(body))
	req.Header.Add("Content-Type", CommandHeader)
	return do(req)
}
