package gonap

import "net/http"
import "fmt"

// PBResp is the struct returned by all Rest request functions
type PBResp struct {
	StatusCode int
	Headers    http.Header
	Body       []byte
}

// ToJson returns PBResp.Body as the raw json
func (r *PBResp) ToJson() {

	fmt.Printf("%s\n", string(r.Body))

}

// PrintHeaders prints the http headers as k,v pairs
func (r *PBResp) PrintHeaders() {
	for key, value := range r.Headers {
		fmt.Println(key, " : ", value[0])
	}

}
