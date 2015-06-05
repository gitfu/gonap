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
}

func (lan *Lan) Listlanmembers() Nics {
	pbresp := is_get(lan.Entities.Nics.Href)
	return toNics(pbresp.Body)
}

type Lans struct {
	Id    string `json:"id,omitempty"`
	Type  string `json:"type,omitempty"`
	Href  string `json:"href,omitempty"`
	Items []Lan  `json:"items,omitempty"`
}

func ToLans(body []byte) Lans {
	var Lans Lans
	json.Unmarshal(body, &Lans)
	return Lans
}

func ToLan(body []byte) Lan {
	var Lan Lan
	json.Unmarshal(body, &Lan)
	return Lan
}

func UpdateLan(dcid string, lanid string, jason []byte) PBResp {
	path := lan_path(dcid, lanid)
	return is_put(path, jason)
}

func PatchLan(dcid string, lanid string, jason []byte) PBResp {
	path := lan_path(dcid, lanid)
	return is_patch(path, jason)
}

func DeleteLan(dcid, lanid string) PBResp {
	path := lan_path(dcid, lanid)
	return is_delete(path)
}

func ListLanMembers(dcid, lanid string) PBResp {
	path := lan_nic_col(dcid, lanid)
	return is_get(path)
}
