package gonap

import "encoding/json"

// Location_Properties is aa struct for Locaation Properties
type Location_Properties struct {
	Name string `json:"name"`
}

// Location is the struct for a Location
type Location struct {
	Id_Type_Href
	MetaData   MetaData            `json:"metadata,omitempty"`
	Properties Location_Properties `json:"properties"`
	Resp       PBResp              `json:"-"`
}

func toLocation(pbresp PBResp) Location {
	var loc Location
	json.Unmarshal(pbresp.Body, &loc)
	loc.Resp = pbresp
	return loc
}

// Locations is the struct for a Location Collection
type Locations struct {
	Id_Type_Href
	Items []Location `json:"items,omitempty"`
	Resp  PBResp     `json:"-"`
}

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
