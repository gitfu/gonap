package gonap

import "encoding/json"

func ToCdRoms(body []byte) Images {
	var CdRoms Images
	json.Unmarshal(body, &CdRoms)
	return CdRoms
}

func ToCdRom(body []byte) Image {
	var CdRom Image
	json.Unmarshal(body, &CdRom)
	return CdRom
}

/**
Server CD Roms collection
	List CD Roms attached
	Attach CD Rom to server

**/

func ListAttachedCdroms(dcid, srvid string) PBResp {
	path := server_cdrom_col_path(dcid, srvid)
	return is_get(path)
}

func AttachCdrom(dcid string, srvid string, cdid string) PBResp {
	jason := []byte(`{"id":"` + cdid + `"}`)
	path := server_cdrom_col_path(dcid, srvid)
	return is_post(path, jason)
}

/**
Server CD Rom
	Retrieve attached CD Rom
	Detach CD Rom from server

**/

func GetAttachedCdrom(dcid, srvid, cdid string) PBResp {
	path := server_cdrom_path(dcid, srvid, cdid)
	return is_get(path)
}

func DetachCdrom(dcid, srvid, cdid string) PBResp {
	path := server_cdrom_path(dcid, srvid, cdid)
	return is_delete(path)
}
