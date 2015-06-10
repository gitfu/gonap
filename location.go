package gonap

import "encoding/json"

func toLocation(pbresp PBResp) Instance {
	var loc Instance
	json.Unmarshal(pbresp.Body, &loc)
	loc.Resp = pbresp
	return loc
}

func toLocations(pbresp PBResp) Collection {
	var locs Collection
	json.Unmarshal(pbresp.Body, &locs)
	locs.Resp = pbresp
	return locs
}

// ListLocations returns location collection data
func ListLocations() Collection {
	return toLocations(is_get(location_col_path()))
}

// GetLocation returns location data
func GetLocation(locid string) Instance {
	return toLocation(is_get(location_path(locid)))
}
