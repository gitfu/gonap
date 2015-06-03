package gonap

import "encoding/json"

type SnapProps struct {
	Name                string `json:"name,omitempty"`
	Description         string `json:"description,omitempty"`
	Location            string `json:"location,omitempty"`
	Size                int    `json:"size,omitempty"`
	CpuHotPlug          bool   `json:"cpuHotPlug,omitempty"`
	CpuHotUnplug        bool   `json:"cpuHotUnplug,omitempty"`
	RamHotPlug          bool   `json:"ramHotPlug,omitempty"`
	RamHotUnplug        bool   `json:"ramHotUnplug,omitempty"`
	NicHotPlug          bool   `json:"nicHotPlug,omitempty"`
	NicHotUnplug        bool   `json:"nicHotUnplug,omitempty"`
	DiscVirtioHotPlug   bool   `json:"discVirtioHotPlug,omitempty"`
	DiscVirtioHotUnplug bool   `json:"discVirtioHotUnplug,omitempty"`
	DiscScsiHotPlug     bool   `json:"discScsiHotPlug,omitempty"`
	DiscScsiHotUnplug   bool   `json:"discScsiHotUnplug,omitempty"`
	LicenceType         string `json:"licenceType,omitempty"`
}

type Snapshot struct {
	Id         string    `json:"id,omitempty"`
	Type       string    `json:"type,omitempty"`
	Href       string    `json:"href,omitempty"`
	MetaData   MetaData  `json:"metadata,omitempty"`
	Properties SnapProps `json:"properties,omitempty"`
}

func ToSnapshot(body []byte) Snapshot {
	var Snapshot Snapshot
	json.Unmarshal(body, &Snapshot)
	return Snapshot
}

type Snapshots struct {
	Id    string     `json:"id,omitempty"`
	Type  string     `json:"type,omitempty"`
	Href  string     `json:"href,omitempty"`
	Items []Snapshot `json:"items,omitempty"`
}

func ToSnapshots(body []byte) Snapshots {
	var Snapshots Snapshots
	json.Unmarshal(body, &Snapshots)
	return Snapshots
}

func ListSnapshots() PBResp {
	path := snapshot_col_path()
	return is_get(path)
}

func GetSnapshot(snapid string) PBResp {
	path := snapshot_path(snapid)
	return is_get(path)
}

func UpdateSnapshot(snapid string, jason []byte) PBResp {
	path := snapshot_path(snapid)
	return is_put(path, jason)
}

func PatchSnapshot(snapid string, jason []byte) PBResp {
	path := snapshot_path(snapid)
	return is_patch(path, jason)
}

func DeleteSnapshot(snapid string) PBResp {
	path := snapshot_path(snapid)
	return is_delete(path)
}
