package gonap

import "encoding/json"

type NicProperties struct {
	Name           string   `json:"name,omitempty"`
	Ips            []string `json:"ips,omitempty"`
	Dhcp           bool     `json:"dhcp,omitempty"`
	Lan            int      `json:"lan"`
	FirewallActive bool     `json:"firewallActive,omitempty"`
	Firewallrules  []FwRule `json:"firewallrules,omitempty"`
}

type Nic struct {
	Id         string        `json:"id,omitempty"`
	Type       string        `json:"type,omitempty"`
	Href       string        `json:"href,omitempty"`
	MetaData   MetaData      `json:"metadata,omitempty"`
	Properties NicProperties `json:"properties,omitempty"`
}

type Nics struct {
	Id    string `json:"id,omitempty"`
	Type  string `json:"type,omitempty"`
	Href  string `json:"href,omitempty"`
	Items []Nic  `json:"items,omitempty"`
}

func AsNics(body []byte) Nics {
	var Nics Nics
	json.Unmarshal(body, &Nics)
	return Nics
}

func AsNic(body []byte) Nic {
	var Nic Nic
	json.Unmarshal(body, &Nic)
	return Nic
}

/**
Nic
Nics collection
	List nics
	Create a nic
Nic
	Retrieve a nic
	Replace properties of nic
	Partially update a nic
	Remove nic
**/

// ListNics returns a PBResp with nic collection data in PBResp.Body
func ListNics(dcid, srvid string) PBResp {
	path := nic_col_path(dcid, srvid)
	return is_get(path)

}

/**
	CreateNic takes []byte jason to create a nic
	Returns a PBResp struct with Nic data in PBResp.Body .
**/
func CreateNic(dcid string, srvid string, jason []byte) PBResp {
	path := nic_col_path(dcid, srvid)
	return is_post(path, jason)
}

// GetNic returns a PBResp with nic data in PBResp.Body
func GetNic(dcid, srvid, nicid string) PBResp {
	path := nic_path(dcid, srvid, nicid)
	return is_get(path)

}

func UpdateNic(dcid string, srvid string, nicid string, jason []byte) PBResp {
	path := nic_path(dcid, srvid, nicid)
	return is_put(path, jason)

}

func PatchNic(dcid string, srvid string, nicid string, jason []byte) PBResp {
	path := nic_path(dcid, srvid, nicid)
	return is_patch(path, jason)

}

func DeleteNic(dcid, srvid, nicid string) PBResp {
	path := nic_path(dcid, srvid, nicid)
	return is_delete(path)
}
