package goprofitbricks

// Listloadbalancers returns a Collection struct
// for loadbalancers in the Datacenter
func ListLoadbalancers(dcid string) Collection {
	path := lbal_col_path(dcid)
	return is_list(path)
}

// Createloadbalancer creates a loadbalancer in the datacenter
//from a jason []byte and returns a Instance struct
func CreateLoadbalancer(dcid string, jason []byte) Instance {
	path := lbal_col_path(dcid)
	return is_post(path, jason)
}

// GetLoadbalancer pulls data for the Loadbalancer
// where id = lbalid returns a Instance struct
func GetLoadbalancer(dcid, lbalid string) Instance {
	path := lbal_path(dcid, lbalid)
	return is_get(path)
}

func UpdateLoadbalancer(dcid string, lbalid string, jason []byte) Instance {
	path := lbal_path(dcid, lbalid)
	return is_put(path, jason)
}

func PatchLoadbalancer(dcid string, lbalid string, jason []byte) Instance {
	path := lbal_path(dcid, lbalid)
	return is_patch(path, jason)
}

func DeleteLoadbalancer(dcid, lbalid string) Resp {
	path := lbal_path(dcid, lbalid)
	return is_delete(path)
}

func ListBalancedNics(dcid, lbalid string) Collection {
	path := balnic_col_path(dcid, lbalid)
	return is_list(path)
}

func AssociateNic(dcid string, lbalid string, nicid string) Instance {

	var sm StringMap
	sm["id"] = nicid
	jason := []byte(MkJson(sm))
	path := balnic_col_path(dcid, lbalid)
	return is_post(path, jason)
}

func GetBalancedNic(dcid, lbalid, balnicid string) Instance {
	path := balnic_path(dcid, lbalid, balnicid)
	return is_get(path)
}

func DeleteBalancedNic(dcid, lbalid, balnicid string) Resp {
	path := balnic_path(dcid, lbalid, balnicid)
	return is_delete(path)
}
