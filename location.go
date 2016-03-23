package goprofitbricks

// ListLocations returns location collection data
func ListLocations() Collection {
	return isList(locationColPath())
}

// GetLocation returns location data
func GetLocation(locid string) Instance {
	return isGet(location_path(locid))
}
