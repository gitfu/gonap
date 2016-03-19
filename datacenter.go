package goprofitbricks

// ListDatacenters returns a Collection struct
func ListDatacenters() Collection {
	path := dc_col_path()
	return isList(path)
}

// CreateDatacenter creates a datacenter and returns a Instance struct
func CreateDatacenter(jason []byte) Instance {
	path := dc_col_path()
	return is_post(path, jason)
}

// GetDatacenter returns a Instance struct where id == dcid
func GetDatacenter(dcid string) Instance {
	path := dc_path(dcid)
	return isGet(path)
}

// UpdateDatacenter updates all Datacenter properties from values in jason
//returns an Instance struct where id ==dcid
func UpdateDatacenter(dcid string, jason []byte) Instance {
	path := dc_path(dcid)
	return isPut(path, jason)
}

// PatchDatacenter replaces any Datacenter properties with the values in jason
//returns an Instance struct where id ==dcid
func PatchDatacenter(dcid string, jason []byte) Instance {
	path := dc_path(dcid)
	return isPatch(path, jason)
}

// DeletesDatacenter  delete where id==dcid
func DeleteDatacenter(dcid string) Resp {
	path := dc_path(dcid)
	return isDelete(path)
}
