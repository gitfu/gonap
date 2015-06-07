package gonap

import "encoding/json"
import "fmt"

type Volume Instance

func (vol *Volume) Save() {
	path := vol.Href
	jason, err := json.MarshalIndent(&vol.Properties, "", "    ")
	if err != nil {
		panic(err)
	}
	pbreq := is_patch(path, jason)
	fmt.Println("save status code is ", pbreq.StatusCode)
}

func toVolume(pbresp PBResp) Volume {
	var volume Volume
	json.Unmarshal(pbresp.Body, &volume)
	volume.Resp = pbresp
	return volume
}

type Volumes Collection
func toVolumes(pbresp PBResp) Volumes {
	var volumes Volumes
	json.Unmarshal(pbresp.Body, &volumes)
	return volumes
}

// ListVolumes returns a Volumes struct for volumes in the Datacenter
func ListVolumes(dcid string) Volumes {
	path := volume_col_path(dcid)
	return toVolumes(is_get(path))
}

func UpdateVolume(dcid string, volid string, jason []byte) Volume {
	path := volume_path(dcid, volid)
	return toVolume(is_put(path, jason))
}

func PatchVolume(dcid string, volid string, jason []byte) Volume {
	path := volume_path(dcid, volid)
	return toVolume(is_patch(path, jason))
}

func DeleteVolume(dcid, volid string) Volume {
	path := volume_path(dcid, volid)
	return toVolume(is_delete(path))
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
