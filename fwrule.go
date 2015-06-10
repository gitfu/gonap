package gonap

import "encoding/json"

func toFwRule(pbresp PBResp) Instance {
	var fwr Instance
	json.Unmarshal(pbresp.Body, &fwr)
	fwr.Resp = pbresp
	return fwr
}

func toFwRules(pbresp PBResp) Collection {
	var fwrs Collection
	json.Unmarshal(pbresp.Body, &fwrs)
	fwrs.Resp = pbresp
	return fwrs
}

// ListFwRules returns a collection of firewall rules
func ListFwRules(dcid, srvid, nicid string) Collection {
	path := fwrule_col_path(dcid, srvid, nicid)
	return toFwRules(is_get(path))
}

// CreateFwRule uses the json values in jason to create a new firewall rule
// and returns an Instance struct
func CreateFwRule(dcid string, srvid string, nicid string, jason []byte) Instance {
	path := fwrule_col_path(dcid, srvid, nicid)
	return toFwRule(is_post(path, jason))
}

// GetFwRule Retrieve a firewall rule and returns Instance struct
func GetFwRule(dcid, srvid, nicid, fwruleid string) Instance {
	path := fwrule_path(dcid, srvid, nicid, fwruleid)
	return toFwRule(is_get(path))
}

// UpdateFwRule Replaces all the properties of firewall rule,
// returns a Instance struct
func UpdateFwRule(dcid string, srvid string, nicid string, fwruleid string, jason []byte) Instance {
	path := fwrule_path(dcid, srvid, nicid, fwruleid)
	return toFwRule(is_put(path, jason))
}

// PatchFwRule 	Partially updates a firewall rule with data from []byte jason
// returns Instance struct
func PatchFWRule(dcid string, srvid string, nicid string, fwruleid string, jason []byte) Instance {
	path := fwrule_path(dcid, srvid, nicid, fwruleid)
	return toFwRule(is_patch(path, jason))
}

// DeleteFwRule removes firewall rule
func DeleteFWRule(dcid, srvid, nicid, fwruleid string) PBResp {
	path := fwrule_path(dcid, srvid, nicid, fwruleid)
	return is_delete(path)
}
