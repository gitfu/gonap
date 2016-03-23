package goprofitbricks

// ListIPBlocks
func ListIPBlocks() Collection {
	path := ipblock_col_path()
	return isList(path)
}

func ReserveIpBlock(jason []byte) Instance {
	path := ipblock_col_path()
	return is_post(path, jason)

}
func GetIpBlock(ipblockid string) Instance {
	path := ipblockPath(ipblockid)
	return isGet(path)
}

func ReleaseIpBlock(ipblockid string) Resp {
	path := ipblockPath(ipblockid)
	return isDelete(path)
}
