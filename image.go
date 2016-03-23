package goprofitbricks

// ListImages returns an Collection struct
func ListImages() Collection {
	path := imageColPath()
	return isList(path)
}

// GetImage returns an Instance struct where id ==imageid
func GetImage(imageid string) Instance {
	path := imagePath(imageid)
	return isGet(path)
}

// UpdateImage updates all image properties from values in jason
//returns an Instance struct where id ==imageid
func UpdateImage(imageid string, jason []byte) Instance {
	path := imagePath(imageid)
	return isPut(path, jason)
}

// PatchImage replaces any image properties from values in jason
//returns an Instance struct where id ==imageid
func PatchImage(imageid string, jason []byte) Instance {
	path := imagePath(imageid)
	return isPatch(path, jason)
}

// Deletes an image where id==imageid
func DeleteImage(imageid string) Resp {
	path := imagePath(imageid)
	return isDelete(path)
}
