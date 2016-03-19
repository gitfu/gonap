package goprofitbricks

import "encoding/json"

type RestRequest_Properties struct {
	Method  string            `json:"method"`
	Headers map[string]string `json:"headers"`
	Body    string            `json:"body"`
	Url     string            `json:"url"`
}

func toRestRequest(resp Resp) Instance {
	var rr Instance
	json.Unmarshal(resp.Body, &rr)
	rr.Resp = resp
	return rr
}

func toRestRequests(resp Resp) Collection {
	var rrs Collection
	json.Unmarshal(resp.Body, &rrs)
	rrs.Resp = resp
	return rrs
}

func ListRequests() Collection {
	path := request_col_path()
	return isList(path)
}

func GetRequest(requestid string) Instance {
	path := request_path(requestid)
	return isGet(path)
}

func StatusRequest(requestid string) Instance {
	path := request_status_path(requestid)
	return isGet(path)
}
