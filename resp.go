package gonap

import "net/http"
import "fmt"

//import "encoding/json"

// PBResp is the struct returned by all Rest request functions
type PBResp struct {
	req_url    string
	StatusCode int
	Headers    http.Header
	Body       []byte
}

// PrintHeaders prints the http headers as k,v pairs
func (r *PBResp) PrintHeaders() {
	for key, value := range r.Headers {
		fmt.Println(key, " : ", value[0])
	}

}
