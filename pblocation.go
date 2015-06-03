package gonap

import "encoding/json"

// LocationProperties is aa struct for Locaation Properties
type LocationProperties struct {
	Name string `json:"name"`
}

// Location is the struct for a Location
type Location struct {
	Id         string             `json:"id,omitempty"`
	Type       string             `json:"type,omitempty"`
	Href       string             `json:"href,omitempty"`
	MetaData   MetaData           `json:"metadata,omitempty"`
	Properties LocationProperties `json:"properties"`
}

// AsLocation Unmarshals json into a Location struct
func AsLocation(body []byte) Location {
	var Location Location
	json.Unmarshal(body, &Location)
	return Location
}

// Locations is the struct for a Location Collection
type Locations struct {
	Id    string     `json:"id,omitempty"`
	Type  string     `json:"type,omitempty"`
	Href  string     `json:"href,omitempty"`
	Items []Location `json:"items,omitempty"`
}

// AsLocations Unmarshals json into a Locations struct
func AsLocations(body []byte) Locations {
	var Locations Locations
	json.Unmarshal(body, &Locations)
	return Locations
}

// ListLocations returns a PBResp with location collection data in the Body
func ListLocations() PBResp {
	path := location_col_path()
	return is_get(path)
}

// GetLocations returns a PBResp with data for a location in the Body
func GetLocation(locid string) PBResp {
	path := location_path(locid)
	return is_get(path)
}
