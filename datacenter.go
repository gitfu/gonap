package gonap

// ListDatacenters returns a Collection struct
func ListDatacenters() Collection {
	path := dc_col_path()
	return toCollection(is_get(path))
}

// CreateDatacenter creates a datacenter and returns a Instance struct
func CreateDatacenter(jason []byte) Instance {
	path := dc_col_path()
	return toInstance(is_post(path, jason))
}

// GetDatacenter returns a Instance struct where id == dcid
func GetDatacenter(dcid string) Instance {
	path := dc_path(dcid)
	return toInstance(is_get(path))
}

// UpdateDatacenter updates all Datacenter properties from values in jason
//returns an Instance struct where id ==dcid
func UpdateDatacenter(dcid string, jason []byte) Instance {
	path := dc_path(dcid)
	return toInstance(is_put(path, jason))
}

// PatchDatacenter replaces any Datacenter properties with the values in jason
//returns an Instance struct where id ==dcid
func PatchDatacenter(dcid string, jason []byte) Instance {
	path := dc_path(dcid)
	return toInstance(is_patch(path, jason))
}

// Deletes a Datacenter where id==dcid
func DeleteDatacenter(dcid string) Resp {
	path := dc_path(dcid)
	return is_delete(path)
}
