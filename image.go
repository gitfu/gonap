package gonap

import "encoding/json"
// Image_Properties for image and cdrom data
type Image_Properties struct {
	Name                string `json:"name,omitempty"`
	Description         string `json:"description,omitempty"`
	Location            string `json:"location"`
	Size                int    `json:"size"`
	Public              bool   `json:"public,omitempty"`
	ImageType           string `json:"imageType,omitempty"`
	CpuHotPlug          bool   `json:"cpuHotPlug,omitempty"`
	CpuHotUnplug        bool   `json:"cpuHotUnplug,omitempty"`
	RamHotPlug          bool   `json:"ramHotPlug,omitempty"`
	RamHotUnplug        bool   `json:"ramHotUnplug,omitempty"`
	NicHotPlug          bool   `json:"nicHotPlug,omitempty"`
	NicHotUnplug        bool   `json:"nicHotUnplug,omitempty"`
	DiscVirtioHotPlug   bool   `json:"discVirtioHotPlug,omitempty"`
	DiscVirtioHotUnplug bool   `json:"discVirtioHotUnplug,omitempty"`
	DiscScsiHotPlug     bool   `json:"discScsiHotPlug,omitempty"`
	DiscScsiHotUnplug   bool   `json:"discScsiHotUnplug,omitempty"`
	LicenceType         string `json:"licenceType,omitempty"`
}

// Image is the struct for image and cdrom data
type Image struct {
	Id_Type_Href
	MetaData   MetaData         `json:"metadata,omitempty"`
	Properties Image_Properties `json:"properties,omitempty"`
	Resp       PBResp           `json:"-"`
}

func toImage(pbresp PBResp) Image {
	var img Image
	json.Unmarshal(pbresp.Body, &img)
	img.Resp = pbresp
	return img
}

// Images is a struct for Image and cdrom collections
type Images struct {
	Id_Type_Href
	Items []Image `json:"items,omitempty"`
	Resp  PBResp  `json:"-"`
}

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

// CreateImage creates an Image and returns an Image struct
func CreateImage(jason []byte) Image {
	path := image_col_path()
	return toImage(is_post(path, jason))
}

// GetImage returns an Image struct where id ==imageid
func GetImage(imageid string) Image {
	path := image_path(imageid)
	return toImage(is_get(path))
}

// UpdateImage updates all image properties from values in jason
//returns an Image struct where id ==imageid
func UpdateImage(imageid string, jason []byte) Image {
	path := image_path(imageid)
	return toImage(is_put(path, jason))
}

// PatchImage replaces any image properties from values in jason
//returns an Image struct where id ==imageid
func PatchImage(imageid string, jason []byte) Image {
	path := image_path(imageid)
	return toImage(is_patch(path, jason))
}

// Deletes an image where id==imageid
func DeleteImage(imageid string) Image {
	path := image_path(imageid)
	return toImage(is_delete(path))
}
