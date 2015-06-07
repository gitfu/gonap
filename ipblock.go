package gonap

import "encoding/json"

// Ipblock is the struct for Ipblock data
type Ipblock Instance

func toIpblock(pbresp PBResp) Ipblock {
	var ipb Ipblock
	json.Unmarshal(pbresp.Body, &ipb)
	ipb.Resp = pbresp
	return ipb
}

// Ipblocks is the struct for an Ipblock collection
type Ipblocks Collection

func toIpblocks(pbresp PBResp) Ipblocks {
	var ipbs Ipblocks
	json.Unmarshal(pbresp.Body, &ipbs)
	ipbs.Resp = pbresp
	return ipbs
}

// ListIpBlocks
func ListIpBlocks() Ipblocks {
	path := ipblock_col_path()
	return toIpblocks(is_get(path))
}

func ReserveIpBlock(jason []byte) Ipblock {
	path := ipblock_col_path()
	return toIpblock(is_post(path, jason))

}
func GetIpBlock(ipblockid string) Ipblock {
	path := ipblock_path(ipblockid)
	return toIpblock(is_get(path))
}

func ReleaseIpBlock(ipblockid string) Ipblock {
	path := ipblock_path(ipblockid)
	return toIpblock(is_delete(path))
}
