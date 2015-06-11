package gonap

import "encoding/json"

func toSnapshot(resp Resp) Instance {
	var snap Instance
	json.Unmarshal(resp.Body, &snap)
	snap.Resp = resp
	return snap
}

func toSnapshots(resp Resp) Collection {
	var snaps Collection
	json.Unmarshal(resp.Body, &snaps)
	snaps.Resp = resp
	return snaps
}

// ListSnapshots retrieves a collection of snapshot data
// returns a Collection struct
func ListSnapshots() Collection {
	path := snapshot_col_path()
	return toSnapshots(is_get(path))
}

// GetSnapshot retrieves Instance data where id==snapid
// returns a` snapshot struct
func GetSnapshot(snapid string) Instance {
	path := snapshot_path(snapid)
	return toSnapshot(is_get(path))
}

// UpdateSnapshot replaces all snapshot properties from values in jason
//returns an Instance struct where id ==snapid
func UpdateSnapshot(snapid string, jason []byte) Instance {
	path := snapshot_path(snapid)
	return toSnapshot(is_put(path, jason))
}

// PatchSnapshot replaces any snapshot properties from values in jason
//returns an Instance struct where id ==snapid
func PatchSnapshot(snapid string, jason []byte) Instance {
	path := snapshot_path(snapid)
	return toSnapshot(is_patch(path, jason))
}

// Deletes a Snapshot with id == snapid
func DeleteSnapshot(snapid string) Resp {
	path := snapshot_path(snapid)
	return is_delete(path)
}
