package gonap

import "encoding/json"

type Loadbalancer Instance

func toLoadbalancer(resp Resp) Instance {
	var lb Instance
	json.Unmarshal(resp.Body, &lb)
	lb.Resp = resp
	return lb
}

type Loadbalancers Collection

func toLoadbalancers(resp Resp) Collection {
	var lbs Collection
	json.Unmarshal(resp.Body, &lbs)
	lbs.Resp = resp
	return lbs
}

// Listloadbalancers returns a Collection struct
// for loadbalancers in the Datacenter
func ListLoadbalancers(dcid string) Collection {
	path := lbal_col_path(dcid)
	return toLoadbalancers(is_get(path))
}

// Createloadbalancer creates a loadbalancer in the datacenter
//from a jason []byte and returns a Instance struct
func CreateLoadbalancer(dcid string, jason []byte) Instance {
	path := lbal_col_path(dcid)
	return toLoadbalancer(is_post(path, jason))
}

// GetLoadbalancer pulls data for the Loadbalancer
// where id = lbalid returns a Instance struct
func GetLoadbalancer(dcid, lbalid string) Instance {
	path := lbal_path(dcid, lbalid)
	return toLoadbalancer(is_get(path))
}

func UpdateLoadbalancer(dcid string, lbalid string, jason []byte) Instance {
	path := lbal_path(dcid, lbalid)
	return toLoadbalancer(is_put(path, jason))
}

func PatchLoadbalancer(dcid string, lbalid string, jason []byte) Instance {
	path := lbal_path(dcid, lbalid)
	return toLoadbalancer(is_patch(path, jason))
}

func DeleteLoadbalancer(dcid, lbalid string) Resp {
	path := lbal_path(dcid, lbalid)
	return is_delete(path)
}

func ListBalancedNics(dcid, lbalid string) Collection {
	path := balnic_col_path(dcid, lbalid)
	return toNics(is_get(path))
}

func AssociateNics(dcid string, lbalid string, jason []byte) Collection {
	path := balnic_col_path(dcid, lbalid)
	return toNics(is_post(path, jason))
}

func GetBalancedNic(dcid, lbalid, balnicid string) Instance {
	path := balnic_path(dcid, lbalid, balnicid)
	return toNic(is_get(path))
}

func DeleteBalancedNic(dcid, lbalid, balnicid string) Resp {
	path := balnic_path(dcid, lbalid, balnicid)
	return is_delete(path)
}
