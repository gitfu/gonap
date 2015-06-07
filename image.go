package gonap

import "encoding/json"

type Image struct {
Obj
}

func toImage(pbresp PBResp) Obj {
	var img Obj
	json.Unmarshal(pbresp.Body, &img)
	img.Resp = pbresp
	return img
}

// Images is a struct for Image and cdrom collections
type Images struct {
Collection
}
//type Fu []Image 

func toImages(pbresp PBResp) Images {
	var imgs Images
	json.Unmarshal(pbresp.Body, &imgs)
	imgs.Resp = pbresp
	return imgs
}

// ListImages returns an Images struct
func ListImages() Images {
	path := image_col_path()
	return toImages(is_get(path))
}


// GetImage returns an Image struct where id ==imageid
func GetImage(imageid string) Obj {
	path := image_path(imageid)
	return toImage(is_get(path))
}

// UpdateImage updates all image properties from values in jason
//returns an Image struct where id ==imageid
func UpdateImage(imageid string, jason []byte) Obj {
	path := image_path(imageid)
	return toImage(is_put(path, jason))
}

// PatchImage replaces any image properties from values in jason
//returns an Image struct where id ==imageid
func PatchImage(imageid string, jason []byte) Obj {
	path := image_path(imageid)
	return toImage(is_patch(path, jason))
}

// Deletes an image where id==imageid
func DeleteImage(imageid string) Obj {
	path := image_path(imageid)
	return toImage(is_delete(path))
}
