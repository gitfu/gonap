package gonap

import "net/http"
import "fmt"
import "encoding/json"

func Mkjson(i interface{}) string {
	jason, err := json.MarshalIndent(&i, "", "    ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(jason))
	return string(jason)
}


// MetaData is a map for metadata returned in a PBResp.Body
type StringMap map[string]string

type StringIfaceMap map[string]interface{}

// PBResp is the struct returned by all Rest request functions
type PBResp struct {
	Req        *http.Request
	StatusCode int
	Headers    http.Header
	Body       []byte
}

type Id_Type_Href struct {
	Id   string `json:"id"`
	Type string `json:"type"`
	Href string `json:"href"`
}

type MetaData StringIfaceMap

type Obj struct {
Id_Type_Href
MetaData	StringMap 		`json:"metaData"`
Properties	StringIfaceMap 	`json:"properties"`
Entities	StringIfaceMap		`json:"entities"`
Resp       	PBResp              	`json:"-"`
}

type Collection struct {
	Id_Type_Href
	Items []Obj `json:"items,omitempty"`
	Resp  PBResp  `json:"-"`
}

// PrintHeaders prints the http headers as k,v pairs
func (r *PBResp) PrintHeaders() {
	for key, value := range r.Headers {
		fmt.Println(key, " : ", value[0])
	}

}
