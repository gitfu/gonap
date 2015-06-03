package gonap

import "encoding/json"

type LbalProps struct {
	Name string `json:"name,omitempty"`
	Ip   string `json:"ip,omitempty"`
	Dhcp bool   `json:"dhcp,omitempty"`
}

type LbalEnts struct {
	BalancedNics Nics `json:"balancednics,omitempty"`
}

type Loadbalancer struct {
	Id         string    `json:"id,omitempty"`
	Type       string    `json:"type,omitempty"`
	Href       string    `json:"href,omitempty"`
	MetaData   MetaData  `json:"metadata,omitempty"`
	Properties LbalProps `json:"properties,omitempty"`
	Entities   LbalEnts  `json:"entities,omitempty"`
}

type Loadbalancers struct {
	Id    string         `json:"id,omitempty"`
	Type  string         `json:"type,omitempty"`
	Href  string         `json:"href,omitempty"`
	Items []Loadbalancer `json:"items,omitempty"`
}

func AsLoadbalancers(body []byte) Loadbalancers {
	var Loadbalancers Loadbalancers
	json.Unmarshal(body, &Loadbalancers)
	return Loadbalancers
}

func AsLoadbalancer(body []byte) Loadbalancer {
	var Loadbalancer Loadbalancer
	json.Unmarshal(body, &Loadbalancer)
	return Loadbalancer
}

func ListLoadbalancers(dcid string) PBResp {
	path := lbal_col_path(dcid)
	return is_get(path)
}

func CreateLoadbalancer(dcid string, jason []byte) PBResp {
	path := lbal_col_path(dcid)
	return is_post(path, jason)
}

func GetLoadbalancer(dcid, lbalid string) PBResp {
	path := lbal_path(dcid, lbalid)
	return is_get(path)
}

func UpdateLoadbalancer(dcid string, lbalid string, jason []byte) PBResp {
	path := lbal_path(dcid, lbalid)
	return is_put(path, jason)
}

func PatchLoadBalancer(dcid string, lbalid string, jason []byte) PBResp {
	path := lbal_path(dcid, lbalid)
	return is_patch(path, jason)
}

func DeleteLoadbalancer(dcid, lbalid string) PBResp {
	path := lbal_path(dcid, lbalid)
	return is_delete(path)
}

func ListBalancedNics(dcid, lbalid string) PBResp {
	path := balnic_col_path(dcid, lbalid)
	return is_get(path)
}

func AssociateNics(dcid string, lbalid string, jason []byte) PBResp {
	path := balnic_col_path(dcid, lbalid)
	return is_post(path, jason)
}

func GetBalancedNic(dcid, lbalid, balnicid string) PBResp {
	path := balnic_path(dcid, lbalid, balnicid)
	return is_get(path)
}

func DeleteBalancedNic(dcid, lbalid, balnicid string) PBResp {
	path := balnic_path(dcid, lbalid, balnicid)
	return is_delete(path)
}
