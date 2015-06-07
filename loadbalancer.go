package gonap

import "encoding/json"

type Loadbalancer Instance

func toLoadbalancer(pbresp PBResp) Loadbalancer {
	var lb Loadbalancer
	json.Unmarshal(pbresp.Body, &lb)
	lb.Resp = pbresp
	return lb
}

type Loadbalancers Collection

func toLoadbalancers(pbresp PBResp) Loadbalancers {
	var lbs Loadbalancers
	json.Unmarshal(pbresp.Body, &lbs)
	lbs.Resp = pbresp
	return lbs
}

// Listloadbalancers returns a Loadbalancers struct
// for loadbalancers in the Datacenter
func ListLoadbalancers(dcid string) Loadbalancers {
	path := lbal_col_path(dcid)
	return toLoadbalancers(is_get(path))
}

// Createloadbalancer creates a loadbalancer in the datacenter
//from a jason []byte and returns a Loadbalancer struct
func CreateLoadbalancer(dcid string, jason []byte) Loadbalancer {
	path := lbal_col_path(dcid)
	return toLoadbalancer(is_post(path, jason))
}

// GetLoadbalancer pulls data for the Loadbalancer
// where id = lbalid returns a Loadbalancer struct
func GetLoadbalancer(dcid, lbalid string) Loadbalancer {
	path := lbal_path(dcid, lbalid)
	return toLoadbalancer(is_get(path))
}

func UpdateLoadbalancer(dcid string, lbalid string, jason []byte) Loadbalancer {
	path := lbal_path(dcid, lbalid)
	return toLoadbalancer(is_put(path, jason))
}

func PatchLoadbalancer(dcid string, lbalid string, jason []byte) Loadbalancer {
	path := lbal_path(dcid, lbalid)
	return toLoadbalancer(is_patch(path, jason))
}

func DeleteLoadbalancer(dcid, lbalid string) Loadbalancer {
	path := lbal_path(dcid, lbalid)
	return toLoadbalancer(is_delete(path))
}

func ListBalancedNics(dcid, lbalid string) Nics {
	path := balnic_col_path(dcid, lbalid)
	return toNics(is_get(path))
}

func AssociateNics(dcid string, lbalid string, jason []byte) Nics {
	path := balnic_col_path(dcid, lbalid)
	return toNics(is_post(path, jason))
}

func GetBalancedNic(dcid, lbalid, balnicid string) Nic {
	path := balnic_path(dcid, lbalid, balnicid)
	return toNic(is_get(path))
}

func DeleteBalancedNic(dcid, lbalid, balnicid string) Nic {
	path := balnic_path(dcid, lbalid, balnicid)
	return toNic(is_delete(path))
}
