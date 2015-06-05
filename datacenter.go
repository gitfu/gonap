package gonap

import "encoding/json"
import "fmt"

const datacenters_path = "/datacenters"

type Datacenter_Properties struct {
	Name        string `json:"name"`
	Location    string `json:"location"`
	Description string `json:"description,omitempty"`
}

// Datacenter_Entities is a struct inside a Datacenter struct to hold collections of other structs
type Datacenter_Entities struct {
	Servers       Servers       `json:"servers,omitempty"`
	Loadbalancers Loadbalancers `json:"loadbalancers,omitempty"`
	Lans          Lans          `json:"lans,omitempty"`
	Volumes       Volumes       `json:"volumes,omitempty"`
}

// Datacenter is  struct to hold data for a datacenter
type Datacenter struct {
	Id         string                `json:"id"`
	Type       string                `json:"type"`
	Href       string                `json:"href"`
	MetaData   MetaData              `json:"metadata,omitempty"`
	Properties Datacenter_Properties `json:"properties"`
	Entities   Datacenter_Entities   `json:"entities,omitempty"`
	Resp       PBResp                `json:"-"`
}

func toDatacenter(pbresp PBResp) Datacenter {
	var DC Datacenter
	json.Unmarshal(pbresp.Body, &DC)
	DC.Resp = pbresp
	return DC
}

// Save converts the datacenter struct's properties to json
// and "patch"es them to the rest server
func (dc *Datacenter) Save() {
	path := dc.Href
	jason, err := json.MarshalIndent(&dc.Properties, "", "    ")
	if err != nil {
		panic(err)
	}
	pbreq := is_patch(path, jason)
	fmt.Println("save status code is ", pbreq.StatusCode)
}

// Datacenter.ToJson marshals the Datacenter struct into json
func (dc *Datacenter) ToJson() string {
	jason, err := json.MarshalIndent(&dc, "", "    ")
	if err != nil {
		panic(err)
	}
	return string(jason)
}

// Datacenters is a struct for Datacenter collections
type Datacenters struct {
	Id    string       `json:"id,omitempty"`
	Type  string       `json:"type,omitempty"`
	Href  string       `json:"href,omitempty"`
	Items []Datacenter `json:"items,omitempty"`
	Resp  PBResp       `json:"-"`
}

func toDatacenters(pbresp PBResp) Datacenters {
	var DCS Datacenters
	json.Unmarshal(pbresp.Body, &DCS)
	DCS.Resp = pbresp
	return DCS
}

// Datacenter.ToJson marshals the Datacenter struct into json
func (dcs *Datacenters) ToJson() string {
	jason, err := json.MarshalIndent(&dcs, "", "    ")
	if err != nil {
		fmt.Println(err)
	}
	return string(jason)
}

// ListDatacenters returns a Datacenter collection struct
func ListDatacenters() Datacenters {
	path := dc_col_path()
	return toDatacenters(is_get(path))
}

// CreateDatacenter creates a datacenter and returns a Datacenter struct
func CreateDatacenter(jason []byte) Datacenter {
	path := dc_col_path()
	return toDatacenter(is_post(path, jason))
}

// GetDatacenter returns a Datacenter struct where id == dcid
func GetDatacenter(dcid string) Datacenter {
	path := dc_path(dcid)
	return toDatacenter(is_get(path))
}

// UpdateDatacenter updates all Datacenter properties from values in jason
//returns an Datacenter struct where id ==dcid
func UpdateDatacenter(dcid string, jason []byte) Datacenter {
	path := dc_path(dcid)
	return toDatacenter(is_put(path, jason))
}

// PatchDatacenter replaces any Datacenter properties with the values in jason
//returns an Datacenter struct where id ==dcid
func PatchDatacenter(dcid string, jason []byte) Datacenter {
	path := dc_path(dcid)
	return toDatacenter(is_patch(path, jason))
}

// Deletes a Datacenter where id==dcid
func DeleteDatacenter(dcid string) Datacenter {
	path := dc_path(dcid)
	return toDatacenter(is_delete(path))
}
