package gonap

import "encoding/json"

type LanProperties struct {
	Name   string `json:"name,omitempty"`
	Public bool   `json:"public,omitempty"`
}

type LanEntities struct {
	Nics Nics `json:"nics,omitempty"`
}

type Lan struct {
	Id         string        `json:"id,omitempty"`
	Type       string        `json:"type,omitempty"`
	Href       string        `json:"href,omitempty"`
	MetaData   MetaData      `json:"metadata,omitempty"`
	Properties LanProperties `json:"properties"`
	Entities   LanEntities   `json:"entities,omitempty"`
	Resp       PBResp        `json:"-"`
}

type Lans struct {
	Id    string `json:"id,omitempty"`
	Type  string `json:"type,omitempty"`
	Href  string `json:"href,omitempty"`
	Items []Lan  `json:"items,omitempty"`
	Resp  PBResp `json:"-"`
}

func toLans(pbresp PBResp) Lans {
	var lans Lans
	json.Unmarshal(pbresp.Body, &lans)
	return lans
}

// toLan converts a PBResp struct into a Lan struct
func toLan(pbresp PBResp) Lan {
	var lan Lan
	json.Unmarshal(pbresp.Body, &lan)
	lan.Resp = pbresp
	return lan
}

// ListLan returns a Lans struct collection for lans in the Datacenter
func ListLans(dcid string) Lans {
	path := lan_col_path(dcid)
	return toLans(is_get(path))
}

// CreateLan creates a lan in the datacenter
// from a jason []byte and returns a Lan struct
func CreateLan(dcid string, jason []byte) Lan {
	path := lan_col_path(dcid)
	return toLan(is_post(path, jason))
}

// GetLan pulls data for the lan where id = lanid returns a Lan struct
func GetLan(dcid, lanid string) Lan {
	path := lan_path(dcid, lanid)
	return toLan(is_get(path))
}

func UpdateLan(dcid string, lanid string, jason []byte) Lan {
	path := lan_path(dcid, lanid)
	return toLan(is_put(path, jason))
}

func PatchLan(dcid string, lanid string, jason []byte) Lan {
	path := lan_path(dcid, lanid)
	return toLan(is_patch(path, jason))
}

// DeleteLan deletes a lan where id == lanid
func DeleteLan(dcid, lanid string) Lan {
	path := lan_path(dcid, lanid)
	return toLan(is_delete(path))

}

// ListLanMembers returns a Nic struct collection for the Lan
func ListLanMembers(dcid, lanid string) Nics {
	path := lan_nic_col(dcid, lanid)
	return toNics(is_get(path))
}
