package gonap

import "encoding/json"

func toIpblock(resp Resp) Instance {
	var ipb Instance
	json.Unmarshal(resp.Body, &ipb)
	ipb.Resp = resp
	return ipb
}

func toIpblocks(resp Resp) Collection {
	var ipbs Collection
	json.Unmarshal(resp.Body, &ipbs)
	ipbs.Resp = resp
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

func ReleaseIpBlock(ipblockid string) Resp {
	path := ipblock_path(ipblockid)
	return is_delete(path)
}
