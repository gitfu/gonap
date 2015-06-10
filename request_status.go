package gonap

import "encoding/json"

type RestRequest_Properties struct {
	Method  string            `json:"method"`
	Headers map[string]string `json:"headers"`
	Body    string            `json:"body"`
	Url     string            `json:"url"`
}

func toRestRequest(pbresp PBResp) Instance {
	var rr Instance
	json.Unmarshal(pbresp.Body, &rr)
	rr.Resp = pbresp
	return rr
}

func toRestRequests(pbresp PBResp) Collection {
	var rrs Collection
	json.Unmarshal(pbresp.Body, &rrs)
	rrs.Resp = pbresp
	return rrs
}

func ListRequests() Collection {
	path := request_col_path()
	return toRestRequests(is_get(path))
}

func GetRequest(requestid string) Instance {
	path := request_path(requestid)
	return toRestRequest(is_get(path))
}

func StatusRequest(requestid string) Instance {
	path := request_status_path(requestid)
	return toRestRequest(is_get(path))
}
