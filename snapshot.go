package gonap

import "encoding/json"


func toSnapshot(pbresp PBResp) Instance {
	var snap Instance
	json.Unmarshal(pbresp.Body, &snap)
	snap.Resp = pbresp
	return snap
}

func toSnapshots(pbresp PBResp) Collection {
	var snaps Collection
	json.Unmarshal(pbresp.Body, &snaps)
	snaps.Resp = pbresp
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
func DeleteSnapshot(snapid string) PBResp {
	path := snapshot_path(snapid)
	return is_delete(path)
}
