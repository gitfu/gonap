package gonap

import "encoding/json"

type IpblockProperties struct {
	Location string `json:"location"`
	Size     int    `json:"size"`
}

type Ipblock struct {
	Id         string            `json:"id,omitempty"`
	Type       string            `json:"type,omitempty"`
	Href       string            `json:"href,omitempty"`
	MetaData   MetaData          `json:"metadata,omitempty"`
	Properties IpblockProperties `json:"properties"`
}

type Ipblocks struct {
	Id    string    `json:"id,omitempty"`
	Type  string    `json:"type,omitempty"`
	Href  string    `json:"href,omitempty"`
	Items []Ipblock `json:"items,omitempty"`
}

func AsIpblocks(body []byte) Ipblocks {
	var Ipblocks Ipblocks
	json.Unmarshal(body, &Ipblocks)
	return Ipblocks
}

func AsIpblock(body []byte) Ipblock {
	var Ipblock Ipblock
	json.Unmarshal(body, &Ipblock)
	return Ipblock
}

func ListIpBlocks() PBResp {
	path := ipblock_col_path()
	return is_get(path)
}
func ReserveIpBlock(jason []byte) PBResp {
	path := ipblock_col_path()
	return is_post(path, jason)

}
func GetIpBlock(ipblockid string) PBResp {
	path := ipblock_path(ipblockid)
	return is_get(path)
}

func ReleaseIpBlock(ipblockid string) PBResp {
	path := ipblock_path(ipblockid)
	return is_delete(path)
}
