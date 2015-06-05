package gonap

import "encoding/json"

type Server_Properties struct {
	Name             string `json:"name,omitempty"`
	Cores            int    `json:"cores,omitempty"`
	Ram              int    `json:"ram,omitempty"`
	Availabilityzone string `json:"availabilityzone,omitempty"`
	Licencetype      string `json:"licencetype,omitempty"`
	Bootvolume       string `json:"bootvolume,omitempty"`
	Bootcdrom        string `json:"bootcdrom,omitempty"`
}

type Server_Entities struct {
	Cdroms  Images  `json:"cdroms,omitempty"`
	Nics    Nics    `json:"nics,omitempty"`
	Volumes Volumes `json:"volumes,omitempty"`
}

// Server is a struct for Server data
type Server struct {
	Id         string            `json:"id,omitempty"`
	Type       string            `json:"type,omitempty"`
	Href       string            `json:"href,omitempty"`
	MetaData   MetaData          `json:"metadata,omitempty"`
	Properties Server_Properties `json:"properties,omitempty"`
	Entities   Server_Entities   `json:"entities,omitempty"`
	Resp       PBResp            `json:"-"`
}

// toServer converts a PBResp struct into a Server struct
func toServer(pbresp PBResp) Server {
	var Server Server
	json.Unmarshal(pbresp.Body, &Server)
	Server.Resp = pbresp
	return Server
}

// Servers is a struct for Server struct collections
type Servers struct {
	Id    string   `json:"id,omitempty"`
	Type  string   `json:"type,omitempty"`
	Href  string   `json:"href,omitempty"`
	Items []Server `json:"items,omitempty"`
	Resp  PBResp   `json:"-"`
}

// toServers converts a PBResp struct into a Servers struct
func toServers(pbresp PBResp) Servers {
	var Servers Servers
	json.Unmarshal(pbresp.Body, &Servers)
	Servers.Resp = pbresp
	return Servers
}

// ListServers returns a server struct collection
func ListServers(dcid string) Servers {
	path := server_col_path(dcid)
	return toServers(is_get(path))
}

// CreateServer creates a server from a jason []byte and returns a Server struct
func CreateServer(dcid string, jason []byte) Server {
	path := server_col_path(dcid)
	return toServer(is_post(path, jason))
}

// GetServer pulls data for the server where id = srvid returns a Server struct
func GetServer(dcid, srvid string) Server {
	path := server_path(dcid, srvid)
	return toServer(is_get(path))
}

// UpdateServer is a full update of server properties passed in as jason []byte
// Returns Server struct
func UpdateServer(dcid string, srvid string, jason []byte) Server {
	path := server_path(dcid, srvid)
	return toServer(is_put(path, jason))
}

// PatchServer partial update of server properties passed in as jason []byte
// Returns Server struct
func PatchServer(dcid string, srvid string, jason []byte) Server {
	path := server_path(dcid, srvid)
	return toServer(is_patch(path, jason))
}

// DeleteServer deletes the server where id=srvid and returns Server struct
func DeleteServer(dcid, srvid string) Server {
	path := server_path(dcid, srvid)
	return toServer(is_delete(path))
}
