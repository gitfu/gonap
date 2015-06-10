package gonap

import "encoding/json"

func toVolume(pbresp PBResp) Instance {
	var volume Instance
	json.Unmarshal(pbresp.Body, &volume)
	volume.Resp = pbresp
	return volume
}

func toVolumes(pbresp PBResp) Collection {
	var volumes Collection
	json.Unmarshal(pbresp.Body, &volumes)
	return volumes
}

// ListVolumes returns a Collection struct for volumes in the Datacenter
func ListVolumes(dcid string) Collection {
	path := volume_col_path(dcid)
	return toVolumes(is_get(path))
}

func UpdateVolume(dcid string, volid string, jason []byte) Instance {
	path := volume_path(dcid, volid)
	return toVolume(is_put(path, jason))
}

func PatchVolume(dcid string, volid string, jason []byte) Instance {
	path := volume_path(dcid, volid)
	return toVolume(is_patch(path, jason))
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
