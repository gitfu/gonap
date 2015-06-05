package gonap

import "encoding/json"
import "fmt"

type VolProps struct {
	Name                string `json:"name,omitempty"`
	Size                int    `json:"size"`
	Bus                 string `json:"bus,omitempty"`
	Image               string `json:"image,omitempty"`
	ImagePassword       string `json:"imagePassword,omitempty"`
	Type                string `json:"type,omitempty"`
	LicenceType         string `json:"licenceType,omitempty"`
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
}

type Volume struct {
	Id         string   `json:"id"`
	Type       string   `json:"type"`
	Href       string   `json:"href"`
	MetaData   MetaData `json:"metadata,omitempty"`
	Properties VolProps `json:"properties,omitempty"`
}

func (vol *Volume) Save() {
	path := vol.Href
	jason, err := json.MarshalIndent(&vol.Properties, "", "    ")
	if err != nil {
		panic(err)
	}
	pbreq := is_patch(path, jason)
	fmt.Println("save status code is ", pbreq.StatusCode)
}

func ToVolume(body []byte) Volume {
	var Volume Volume
	json.Unmarshal(body, &Volume)
	return Volume
}

type Volumes struct {
	Id    string   `json:"id"`
	Type  string   `json:"type"`
	Href  string   `json:"href"`
	Items []Volume `json:"items,omitempty"`
}

func ToVolumes(body []byte) Volumes {
	var Volumes Volumes
	json.Unmarshal(body, &Volumes)
	return Volumes
}

func UpdateVolume(dcid string, volid string, jason []byte) PBResp {
	path := volume_path(dcid, volid)
	return is_put(path, jason)
}

func PatchVolume(dcid string, volid string, jason []byte) PBResp {
	path := volume_path(dcid, volid)
	return is_patch(path, jason)
}

func DeleteVolume(dcid, volid string) PBResp {
	path := volume_path(dcid, volid)
	return is_delete(path)
}

/**

createSnapshot : function(dc_id,volume_id,jason,callback){
                var str=""
		if (jason){
                	if (jason['name']){
                        	str +=('&name='+jason['name'])
                	}
                	if (jason['description']){
                        	str +=('&description='+jason['description'])
                	}
		}
                pbreq.is_command([ "datacenters",dc_id,"volumes",volume_id,"create-snapshot" ],str,callback)
        },



	restoreSnapshot : function(dc_id,volume_id,jason,callback){
		pbreq.is_post([ "datacenters",dc_id,"volumes",volume_id,"restore-snapshot" ],str,callback)
	}
**/
