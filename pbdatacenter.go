package gonap

import "encoding/json"

type DcProperties struct {
	Name        string `json:"name"`
	Location    string `json:"location"`
	Description string `json:"description,omitempty"`
}

type DcEntities struct {
	Servers       Servers       `json:"servers,omitempty"`
	Loadbalancers Loadbalancers `json:"loadbalancers,omitempty"`
	Lans          Lans          `json:"lans,omitempty"`
	Volumes       Volumes       `json:"volumes,omitempty"`
}

// Datacenter is  struct to hold data for a datacenter
type Datacenter struct {
	Id         string       `json:"id,omitempty"`
	Type       string       `json:"type,omitempty"`
	Href       string       `json:"href,omitempty"`
	MetaData   MetaData     `json:"metadata,omitempty"`
	Properties DcProperties `json:"properties,omitempty"`
	Entities   DcEntities   `json:"entities,omitempty"`
}

/**
func (dc *Datacenter) Asjson() []byte {

}
**/
func (dc *Datacenter) Listservers() Servers {
	pbresp := ListServers(dc.Id)
	return AsServers(pbresp.Body)
}
func (dc *Datacenter) Getserver(srvid string) Server {
	path := dc.Entities.Servers.Href + slash(srvid)
	pbresp := is_get(path)
	return AsServer(pbresp.Body)

}
func (dc *Datacenter) Listvolumes() Volumes {
	pbresp := ListVolumes(dc.Id)
	return AsVolumes(pbresp.Body)
}
func (dc *Datacenter) Listlans() Lans {
	pbresp := ListLans(dc.Id)
	return AsLans(pbresp.Body)
}

func (dc *Datacenter) Listloadbalancers() Loadbalancers {
	pbresp := ListLoadbalancers(dc.Id)
	return AsLoadbalancers(pbresp.Body)
}

// AsDatacenter unmarshalls a []byte into a Datacenter struct
func AsDatacenter(body []byte) Datacenter {
	var Datacenter Datacenter
	json.Unmarshal(body, &Datacenter)
	return Datacenter
}

// Datacenters is a struct for Datacenter collections
type Datacenters struct {
	Id    string       `json:"id,omitempty"`
	Type  string       `json:"type,omitempty"`
	Href  string       `json:"href,omitempty"`
	Items []Datacenter `json:"items,omitempty"`
}

// AsDatacenter unmarshalls a []byte into a Datacenters struct

func AsDatacenters(body []byte) Datacenters {
	var Datacenters Datacenters
	json.Unmarshal(body, &Datacenters)
	return Datacenters
}

func ListDatacenters() PBResp {
	path := dc_col_path()
	return is_get(path)

}
func CreateDatacenter(jason []byte) PBResp {
	path := dc_col_path()
	return is_post(path, jason)

}

func GetDatacenter(dcid string) PBResp {
	path := dc_path(dcid)
	return is_get(path)
}

func UpdateDatacenter(dcid string, jason []byte) PBResp {
	path := dc_path(dcid)
	return is_put(path, jason)
}

func PatchDatacenter(dcid string, jason []byte) PBResp {
	path := dc_path(dcid)
	return is_patch(path, jason)
}

func DeleteDatacenter(dcid string) PBResp {
	path := dc_path(dcid)
	return is_delete(path)
}
