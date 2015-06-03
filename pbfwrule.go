package gonap

import "encoding/json"

type FwRuleProperties struct {
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
	Id         string           `json:"id,omitempty"`
	Type       string           `json:"type,omitempty"`
	Href       string           `json:"href,omitempty"`
	MetaData   MetaData         `json:"metadata,omitempty"`
	Properties FwRuleProperties `json:"properties,omitempty"`
}

func AsFwRule(body []byte) FwRule {
	var FwRule FwRule
	json.Unmarshal(body, &FwRule)
	return FwRule
}

type FwRules struct {
	Id    string   `json:"id,omitempty"`
	Type  string   `json:"type,omitempty"`
	Href  string   `json:"href,omitempty"`
	Items []FwRule `json:"items,omitempty"`
}

func AsFwRules(body []byte) FwRules {
	var FwRules FwRules
	json.Unmarshal(body, &FwRules)
	return FwRules
}

func ListFwRules(dcid, srvid, nicid string) PBResp {
	path := fwrule_col_path(dcid, srvid, nicid)
	return is_get(path)
}

func CreateFwRule(dcid string, srvid string, nicid string, jason []byte) PBResp {
	path := fwrule_col_path(dcid, srvid, nicid)
	return is_post(path, jason)
}

/**
Firewall rule
	Retrieve a firewall rule
	Replace properties of firewall rule
	Partially update a firewall rule
	Remove firewall rule
**/
func GetFwRule(dcid, srvid, nicid, fwruleid string) PBResp {
	path := fwrule_path(dcid, srvid, nicid, fwruleid)
	return is_get(path)
}

func UpdateFwRule(dcid string, srvid string, nicid string, fwruleid string, jason []byte) PBResp {
	path := fwrule_path(dcid, srvid, nicid, fwruleid)
	return is_put(path, jason)
}

func PatchFWRule(dcid string, srvid string, nicid string, fwruleid string, jason []byte) PBResp {
	path := fwrule_path(dcid, srvid, nicid, fwruleid)
	return is_patch(path, jason)
}

func DeleteFWRule(dcid, srvid, nicid, fwruleid string) PBResp {
	path := fwrule_path(dcid, srvid, nicid, fwruleid)
	return is_delete(path)
}
