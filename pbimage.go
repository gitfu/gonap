package gonap

import "encoding/json"

type ImageProperties struct {
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

type Image struct {
	Id         string          `json:"id,omitempty"`
	Type       string          `json:"type,omitempty"`
	Href       string          `json:"href,omitempty"`
	MetaData   MetaData        `json:"metadata,omitempty"`
	Properties ImageProperties `json:"properties,omitempty"`
}

type Images struct {
	Id    string  `json:"id,omitempty"`
	Type  string  `json:"type,omitempty"`
	Href  string  `json:"href,omitempty"`
	Items []Image `json:"items,omitempty"`
}

func AsImage(body []byte) Image {
	var Image Image
	json.Unmarshal(body, &Image)
	return Image
}

func AsImages(body []byte) Images {
	var Images Images
	json.Unmarshal(body, &Images)
	return Images
}

func ListImages() PBResp {
	path := image_col_path()
	return is_get(path)
}

func GetImage(imageid string) PBResp {
	path := image_path(imageid)
	return is_get(path)
}

func UpdateImage(imageid string, jason []byte) PBResp {
	path := image_path(imageid)
	return is_put(path, jason)
}

func PatchImage(imageid string, jason []byte) PBResp {
	path := image_path(imageid)
	return is_patch(path, jason)
}

func DeleteImage(imageid string) PBResp {
	path := image_path(imageid)
	return is_delete(path)
}
