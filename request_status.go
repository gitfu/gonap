package gonap

import "encoding/json"

type RestRequest_Properties struct {
	Method  string            `json:"method"`
	Headers map[string]string `json:"headers"`
	Body    string            `json:"body"`
	Url     string            `json:"url"`
}

type RestRequest struct {
	Id_Type_Href
	Metadata               MetaData `json:"metadata"`
	RestRequest_Properties `json:"properties"`
	Resp                   PBResp `json:"-"`
}

func toRestRequest(pbresp PBResp) RestRequest {
	var rr RestRequest
	json.Unmarshal(pbresp.Body, &rr)
	rr.Resp = pbresp
	return rr
}

func ListRequests() PBResp {
	path := request_col_path()
	return is_get(path)
}

func GetRequest(requestid string) RestRequest {
	path := request_path(requestid)
	return toRestRequest(is_get(path))
}

func StatusRequest(requestid string) PBResp {
	path := request_status_path(requestid)
	return is_get(path)
}
