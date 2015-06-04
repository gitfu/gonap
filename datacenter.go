package gonap

import "encoding/json"
import "fmt"

const datacenters_path = "/datacenters"

type DcProperties struct {
	Name        string `json:"name"`
	Location    string `json:"location"`
	Description string `json:"description,omitempty"`
}
// DCEntities is a struct inside a Datacenter struct to hold collections of other structs
type DcEntities struct {
	Servers       Servers       `json:"servers,omitempty"`
	Loadbalancers Loadbalancers `json:"loadbalancers,omitempty"`
	Lans          Lans          `json:"lans,omitempty"`
	Volumes       Volumes       `json:"volumes,omitempty"`
}

// Datacenter is  struct to hold data for a datacenter
type Datacenter struct {
	Id         string       `json:"id"`
	Type       string       `json:"type"`
	Href       string       `json:"href"`
	MetaData   MetaData     `json:"metadata,omitempty"`
	Properties DcProperties `json:"properties"`
	Entities   DcEntities   `json:"entities,omitempty"`
}

func (dc *Datacenter)Save() {
	path := dc.Href
	jason,err := json.MarshalIndent(&dc.Properties,"","    ")
	if err !=nil {panic(err)}
	pbreq := is_patch(path, jason)
	fmt.Println("save status code is ",pbreq.StatusCode)
}

// Datacenter.Tojson marshals the Datacenter struct into json
func (dc *Datacenter) Tojson() []byte {
	jason, err := json.MarshalIndent(&dc,"","    ")
	if err !=nil {
		fmt.Println(err)
	}
	//fmt.Println(string(jason))
	return jason
}

// Listservers lists the servers in the datacenter
func (dc *Datacenter) Listservers() Servers {
	pbresp:=is_get(dc.Entities.Servers.Href)
	return ToServers(pbresp.Body)
}
// Createserver creates a server from a jason []byte and returns a Server struct
func  (dc *Datacenter) Createserver(jason []byte) Server {
	path := dc.Entities.Servers.Href
	pbresp := is_post(path, jason)
	return ToServer(pbresp.Body)
}

func (dc * Datacenter) Getserver(srvid string) Server {
	path := dc.Entities.Servers.Href+ slash(srvid)
	pbresp:=is_get(path)
	return ToServer(pbresp.Body)
}


func (dc * Datacenter) Deleteserver(srvid string) PBResp {
	path := dc.Entities.Servers.Href+ slash(srvid)
	pbresp:=is_delete(path)
	return pbresp
}

func (dc *Datacenter) Listlans() Lans {
	pbresp:=is_get(dc.Entities.Lans.Href)
	return ToLans(pbresp.Body)
}

func  (dc *Datacenter) Createlan(jason []byte) Lan {
	path := dc.Entities.Lans.Href
	pbresp := is_post(path, jason)
	return ToLan(pbresp.Body)
}

func (dc * Datacenter) Getlan(lanid string) Lan {
	path := dc.Entities.Lans.Href+ slash(lanid)
	pbresp:=is_get(path)
	return ToLan(pbresp.Body)
}

func (dc * Datacenter) Deletelan(lanid string) PBResp {
	path := dc.Entities.Lans.Href+ slash(lanid)
	pbresp:=is_delete(path)
	return pbresp
}

func (dc *Datacenter) Listloadbalancers() Loadbalancers {
	pbresp := ListLoadbalancers(dc.Id)
	return ToLoadbalancers(pbresp.Body)
}


func  (dc *Datacenter) Createloadbalancer(jason []byte) Loadbalancer {
	path := dc.Entities.Loadbalancers.Href
	pbresp := is_post(path, jason)
	return ToLoadbalancer(pbresp.Body)
}

func (dc * Datacenter) Getloadbalancer(lbalid string) Loadbalancer {
	path := dc.Entities.Loadbalancers.Href+ slash(lbalid)
	pbresp:=is_get(path)
	return ToLoadbalancer(pbresp.Body)
}

func (dc *Datacenter) Listvolumes() Volumes {
	pbresp:=is_get(dc.Entities.Volumes.Href)
	return ToVolumes(pbresp.Body)
}


// ToDatacenter unmarshalls a []byte into a Datacenter struct
func ToDatacenter(body []byte) Datacenter {
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

func ToDatacenters(body []byte) Datacenters {
	var Datacenters Datacenters
	json.Unmarshal(body, &Datacenters)
	return Datacenters
}



func ListDatacenters() PBResp {
	return is_get(datacenters_path)
}
func CreateDatacenter(jason []byte) PBResp {
	return is_post(datacenters_path, jason)
}

func GetDatacenter(dcid string) PBResp {
	path := datacenters_path+slash(dcid)
	return is_get(path)
}

func UpdateDatacenter(dcid string, jason []byte) PBResp {
	path := datacenters_path+slash(dcid)
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
