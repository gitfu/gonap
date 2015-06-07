package gonap

import "encoding/json"

// FwRule is struct for Firewall rule instance data
type FwRule Instance

func toFwRule(pbresp PBResp) FwRule {
	var fwr FwRule
	json.Unmarshal(pbresp.Body, &fwr)
	fwr.Resp = pbresp
	return fwr
}

// FwRules is the struct for firewall rule collections
type FwRules Collection

func toFwRules(pbresp PBResp) FwRules {
	var fwrs FwRules
	json.Unmarshal(pbresp.Body, &fwrs)
	fwrs.Resp = pbresp
	return fwrs
}

// ListFwRules returns a collection of firewall rules
func ListFwRules(dcid, srvid, nicid string) FwRules {
	path := fwrule_col_path(dcid, srvid, nicid)
	return toFwRules(is_get(path))
}

// CreateFwRule uses the json values in jason to create a new firewall rule
// and returns a FwRule struct
func CreateFwRule(dcid string, srvid string, nicid string, jason []byte) FwRule {
	path := fwrule_col_path(dcid, srvid, nicid)
	return toFwRule(is_post(path, jason))
}

// GetFwRule Retrieve a firewall rule and returns fwRule struct
func GetFwRule(dcid, srvid, nicid, fwruleid string) FwRule {
	path := fwrule_path(dcid, srvid, nicid, fwruleid)
	return toFwRule(is_get(path))
}

// UpdateFwRule Replaces all the properties of firewall rule,
// returns a FwRule struct
func UpdateFwRule(dcid string, srvid string, nicid string, fwruleid string, jason []byte) FwRule {
	path := fwrule_path(dcid, srvid, nicid, fwruleid)
	return toFwRule(is_put(path, jason))
}

// PatchFwRule 	Partially updates a firewall rule with data from []byte jason
// returns FwRule struct
func PatchFWRule(dcid string, srvid string, nicid string, fwruleid string, jason []byte) FwRule {
	path := fwrule_path(dcid, srvid, nicid, fwruleid)
	return toFwRule(is_patch(path, jason))
}

// DeleteFwRule removes firewall rule
func DeleteFWRule(dcid, srvid, nicid, fwruleid string) FwRule {
	path := fwrule_path(dcid, srvid, nicid, fwruleid)
	return toFwRule(is_delete(path))
}
