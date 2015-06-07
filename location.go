package gonap

import "encoding/json"

// Location is the struct for a Location
type Location Instance

func toLocation(pbresp PBResp) Location {
	var loc Location
	json.Unmarshal(pbresp.Body, &loc)
	loc.Resp = pbresp
	return loc
}

// Locations is the struct for a Location Collection
type Locations Collection

func toLocations(pbresp PBResp) Locations {
	var locs Locations
	json.Unmarshal(pbresp.Body, &locs)
	locs.Resp = pbresp
	return locs
}

// ListLocations returns a PBResp with location collection data in the Body
func ListLocations() Locations {
	return toLocations(is_get(location_col_path()))
}

// GetLocation returns a PBResp with data for a location in the Body
func GetLocation(locid string) Location {
	path := location_path(locid)
	return toLocation(is_get(path))
}
