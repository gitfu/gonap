package gonap

import "encoding/json"

func toImage(resp Resp) Instance {
	var img Instance
	json.Unmarshal(resp.Body, &img)
	img.Resp = resp
	return img
}

func toImages(resp Resp) Collection {
	var imgs Collection
	json.Unmarshal(resp.Body, &imgs)
	imgs.Resp = resp
	return imgs
}

// ListImages returns an Collection struct
func ListImages() Collection {
	path := image_col_path()
	return toImages(is_get(path))
}

// GetImage returns an Instance struct where id ==imageid
func GetImage(imageid string) Instance {
	path := image_path(imageid)
	return toImage(is_get(path))
}

// UpdateImage updates all image properties from values in jason
//returns an Instance struct where id ==imageid
func UpdateImage(imageid string, jason []byte) Instance {
	path := image_path(imageid)
	return toImage(is_put(path, jason))
}

// PatchImage replaces any image properties from values in jason
//returns an Instance struct where id ==imageid
func PatchImage(imageid string, jason []byte) Instance {
	path := image_path(imageid)
	return toImage(is_patch(path, jason))
}

// Deletes an image where id==imageid
func DeleteImage(imageid string) Resp {
	path := image_path(imageid)
	return is_delete(path)
}
