package goprofitbricks

// ListSnapshots retrieves a collection of snapshot data
// returns a Collection struct
func ListSnapshots() Collection {
	path := snapshot_col_path()
	return is_list(path)
}

// GetSnapshot retrieves Instance data where id==snapid
// returns a` snapshot struct
func GetSnapshot(snapid string) Instance {
	path := snapshotPath(snapid)
	return is_get(path)
}

// UpdateSnapshot replaces all snapshot properties from values in jason
//returns an Instance struct where id ==snapid
func UpdateSnapshot(snapid string, jason []byte) Instance {
	path := snapshotPath(snapid)
	return is_put(path, jason)
}

// PatchSnapshot replaces any snapshot properties from values in jason
//returns an Instance struct where id ==snapid
func PatchSnapshot(snapid string, jason []byte) Instance {
	path := snapshotPath(snapid)
	return is_patch(path, jason)
}

// Deletes a Snapshot with id == snapid
func DeleteSnapshot(snapid string) Resp {
	path := snapshotPath(snapid)
	return is_delete(path)
}
