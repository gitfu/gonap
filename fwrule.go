package gonap

import "encoding/json"

type FwRule_Properties struct {
	Name           string `json:"name,omitempty"`
	Protocol       string `json:"protocol,omitempty"`
	SourceMac      string `json:"sourceMac,omitempty"`
	SourceIp       string `json:"sourceIp,omitempty"`
	TargetIp       string `json:"targetIp,omitempty"`
	PortRangeStart string `json:"portRangeStart,omitempty"`
	PortRangeEnd   string `json:"portRangeEnd,omitempty"`
	IcmpType       string `json:"icmpType,omitempty"`
	IcmpCode       string `json:"icmpCode,omitempty"`
}

type FwRule struct {
	Id_Type_Href
	MetaData   MetaData          `json:"metadata,omitempty"`
	Properties FwRule_Properties `json:"properties,omitempty"`
	Resp       PBResp            `json:"-"`
}

func toFwRule(pbresp PBResp) FwRule {
	var fwr FwRule
	json.Unmarshal(pbresp.Body, &fwr)
	fwr.Resp = pbresp
	return fwr
}

type FwRules struct {
	Id_Type_Href
	Items []FwRule `json:"items,omitempty"`
	Resp  PBResp   `json:"-"`
}

func toFwRules(pbresp PBResp) FwRules {
	var fwrs FwRules
	json.Unmarshal(pbresp.Body, &fwrs)
	fwrs.Resp = pbresp
	return fwrs
}

func ListFwRules(dcid, srvid, nicid string) FwRules {
	path := fwrule_col_path(dcid, srvid, nicid)
	return toFwRules(is_get(path))
}

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
