package gonap

import "encoding/json"

type SrvProps struct {
	Name             string `json:"name,omitempty"`
	Cores            int    `json:"cores,omitempty"`
	Ram              int    `json:"ram,omitempty"`
	Availabilityzone string `json:"availabilityzone,omitempty"`
	Licencetype      string `json:"licencetype,omitempty"`
	Bootvolume       string `json:"bootvolume,omitempty"`
	Bootcdrom        string `json:"bootcdrom,omitempty"`
}

type SrvEnts struct {
	Cdroms  Images  `json:"cdroms,omitempty"`
	Nics    Nics    `json:"nics,omitempty"`
	Volumes Volumes `json:"volumes,omitempty"`
}

type Server struct {
	Id         string   `json:"id,omitempty"`
	Type       string   `json:"type,omitempty"`
	Href       string   `json:"href,omitempty"`
	MetaData   MetaData `json:"metadata,omitempty"`
	Properties SrvProps `json:"properties,omitempty"`
	Entities   SrvEnts  `json:"entities,omitempty"`
}

func ToServer(body []byte) Server {
	var Server Server
	json.Unmarshal(body, &Server)
	return Server
}

type Servers struct {
	Id    string   `json:"id,omitempty"`
	Type  string   `json:"type,omitempty"`
	Href  string   `json:"href,omitempty"`
	Items []Server `json:"items,omitempty"`
}

func ToServers(body []byte) Servers {
	var Servers Servers
	json.Unmarshal(body, &Servers)
	return Servers
}

func ListServers(dcid string) PBResp {
	path := server_col_path(dcid)
	return is_get(path)
}

func CreateServer(dcid string, jason []byte) PBResp {
	path := server_col_path(dcid)
	return is_post(path, jason)
}

func GetServer(dcid, srvid string) PBResp {
	path := server_path(dcid, srvid)
	return is_get(path)
}

func UpdateServer(dcid string, srvid string, jason []byte) PBResp {
	path := server_path(dcid, srvid)
	return is_put(path, jason)
}

func PatchServer(dcid string, srvid string, jason []byte) PBResp {
	path := server_path(dcid, srvid)
	return is_patch(path, jason)
}

func DeleteServer(dcid, srvid string) PBResp {
	path := server_path(dcid, srvid)
	return is_delete(path)
}
