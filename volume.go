package goprofitbricks

import "encoding/json"
import "fmt"

// ListVolumes returns a Collection struct for volumes in the Datacenter
func ListVolumes(dcid string) Collection {
	path := volume_col_path(dcid)
	return isList(path)
}

func UpdateVolume(dcid string, volid string, jason []byte) Instance {
	path := volumePath(dcid, volid)
	return isPut(path, jason)
}

func PatchVolume(dcid string, volid string, jason []byte) Instance {
	path := volumePath(dcid, volid)
	return isPatch(path, jason)
}

func DeleteVolume(dcid, volid string) Resp {
	path := volumePath(dcid, volid)
	return isDelete(path)
}

func CreateSnapshot(dcid string, volid string, jason []byte) Resp {

	empty := `
		{}
		`
	var path = volumePath(dcid, volid)
	var data StringMap
	json.Unmarshal(jason, &data)
	for key, value := range data {
		path += ("&" + key + "=" + value)
		fmt.Println(path)
	}
	return is_command(path, empty)
}

/**



	restoreSnapshot : function(dc_id,volume_id,jason,callback){
		pbreq.is_post([ "datacenters",dc_id,"volumes",volume_id,"restore-snapshot" ],str,callback)
	}
**/
