package gonap

import "net/http"
import "fmt"

//import "encoding/json"

type Id_Type_Href struct {
	Id   string `json:"id"`
	Type string `json:"type"`
	Href string `json:"href"`
}

// MetaData is a struct for metadata returned in a PBResp.Body
type MetaData struct {
	LastModified   string `json:"lastModifiedDate,omitempty"`
	LastModifiedBy string `json:"lastModifiedBy,omitempty"`
	Created        string `json:"createdDate,omitempty"` //"2014-02-01T11:12:12Z",
	CreatedBy      string `json:"createdBy,omitempty"`
	State          string `json:"state,omitempty"`
	Etag           string `json:"etag,omitempty"`
}

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
