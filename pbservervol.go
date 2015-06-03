package gonap

/**
Server volumes collection
	List volumes attached
	Attach volume to server
**/

func ListAttachedVolumes(dcid, srvid string) PBResp {
	path := server_volume_col_path(dcid, srvid)
	return is_get(path)
}

func AttachVolume(dcid string, srvid string, volid string) PBResp {
	jason := []byte(`{"id":"` + volid + `"}`)
	path := server_volume_path(dcid, srvid, volid)
	return is_post(path, jason)
}

/**
Server volume
	Retrieve attached volume
	Detach volume from server

**/

func GetAttachedVolume(dcid, srvid, volid string) PBResp {
	path := server_volume_path(dcid, srvid, volid)
	return is_get(path)
}

func DetachVolume(dcid, srvid, volid string) PBResp {
	path := server_volume_path(dcid, srvid, volid)
	return is_delete(path)
}
