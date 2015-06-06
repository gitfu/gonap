package gonap

import "net/http"
import "fmt"

//import "encoding/json"

type Id_Type_Href struct {
	Id   string `json:"id"`
	Type string `json:"type"`
	Href string `json:"href"`
}

// MetaData is a map for metadata returned in a PBResp.Body
type MetaData map[string]string

type Props map[string]string

// PBResp is the struct returned by all Rest request functions
type PBResp struct {
	Req        *http.Request
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
