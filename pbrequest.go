package gonap

func ListRequests() PBResp {
	path := request_col_path()
	return is_get(path)
}

func GetRequest(requestid string) PBResp {
	path := request_path(requestid)
	return is_get(path)
}

func StatusRequest(requestid string) PBResp {
	path := request_status_path(requestid)
	return is_get(path)
}
