package gonap

import "encoding/json"

func toLocation(resp Resp) Instance {
	var loc Instance
	json.Unmarshal(resp.Body, &loc)
	loc.Resp = resp
	return loc
}

func toLocations(resp Resp) Collection {
	var locs Collection
	json.Unmarshal(resp.Body, &locs)
	locs.Resp = resp
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
