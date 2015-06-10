package gonap

import "encoding/json"

func toIpblock(pbresp PBResp) Instance {
	var ipb Instance
	json.Unmarshal(pbresp.Body, &ipb)
	ipb.Resp = pbresp
	return ipb
}

func toIpblocks(pbresp PBResp) Collection {
	var ipbs Collection
	json.Unmarshal(pbresp.Body, &ipbs)
	ipbs.Resp = pbresp
	return ipbs
}

// ListIpBlocks
func ListIpBlocks() Collection {
	path := ipblock_col_path()
	return toIpblocks(is_get(path))
}

func ReserveIpBlock(jason []byte) Instance {
	path := ipblock_col_path()
	return toIpblock(is_post(path, jason))

}
func GetIpBlock(ipblockid string) Instance {
	path := ipblock_path(ipblockid)
	return toIpblock(is_get(path))
}

func ReleaseIpBlock(ipblockid string) PBResp {
	path := ipblock_path(ipblockid)
	return is_delete(path)
}
