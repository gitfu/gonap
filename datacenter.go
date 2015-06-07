package gonap

import "encoding/json"

// Datacenter is  struct to hold data for a datacenter
type Datacenter Instance

func toDatacenter(pbresp PBResp) Datacenter {
	var DC Datacenter
	json.Unmarshal(pbresp.Body, &DC)
	DC.Resp = pbresp
	return DC
}

// Datacenters is a struct for Datacenter collections
type Datacenters Collection 

func toDatacenters(pbresp PBResp) Datacenters {
	var DCS Datacenters
	json.Unmarshal(pbresp.Body, &DCS)
	DCS.Resp = pbresp
	return DCS
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
