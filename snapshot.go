package gonap

import "encoding/json"

// Snapshot struct for Snapshot data
type Snapshot Instance

func toSnapshot(pbresp PBResp) Snapshot {
	var snap Snapshot
	json.Unmarshal(pbresp.Body, &snap)
	snap.Resp = pbresp
	return snap
}

// Snapshots struct for a Snapshot collection
type Snapshots Collection

func toSnapshots(pbresp PBResp) Snapshots {
	var snaps Snapshots
	json.Unmarshal(pbresp.Body, &snaps)
	snaps.Resp = pbresp
	return snaps
}

// ListSnapshots retrieves a collection of snapshot data
// returns a Snapshots struct
func ListSnapshots() Snapshots {
	path := snapshot_col_path()
	return toSnapshots(is_get(path))
}

// GetSnapshot retrieves Snapshot data where id==snapid
// returns a` snapshot struct
func GetSnapshot(snapid string) Snapshot {
	path := snapshot_path(snapid)
	return toSnapshot(is_get(path))
}

// UpdateSnapshot replaces all snapshot properties from values in jason
//returns an Snapshot struct where id ==snapid
func UpdateSnapshot(snapid string, jason []byte) Snapshot {
	path := snapshot_path(snapid)
	return toSnapshot(is_put(path, jason))
}

// PatchSnapshot replaces any snapshot properties from values in jason
//returns an Snapshot struct where id ==snapid
func PatchSnapshot(snapid string, jason []byte) Snapshot {
	path := snapshot_path(snapid)
	return toSnapshot(is_patch(path, jason))
}

// Deletes a Snapshot with id == snapid
func DeleteSnapshot(snapid string) Snapshot {
	path := snapshot_path(snapid)
	return toSnapshot(is_delete(path))
}
