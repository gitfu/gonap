package goprofitbricks

import "encoding/json"
import "fmt"

// ListVolumes returns a Collection struct for volumes in the Datacenter
func ListVolumes(dcid string) Collection {
	path := volume_col_path(dcid)
	return is_list(path)
}

func UpdateVolume(dcid string, volid string, jason []byte) Instance {
	path := volume_path(dcid, volid)
	return is_put(path, jason)
}

func PatchVolume(dcid string, volid string, jason []byte) Instance {
	path := volume_path(dcid, volid)
	return is_patch(path, jason)
}

func DeleteVolume(dcid, volid string) Resp {
	path := volume_path(dcid, volid)
	return is_delete(path)
}

func CreateSnapshot(dcid string, volid string, jason []byte) Resp {

	empty := `
		{}
		`
	var path = volume_path(dcid, volid)
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
