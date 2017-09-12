package goprofitbricks

// ListIPBlocks
func ListIPBlocks() Collection {
	path := ipblock_col_path()
	return isList(path)
}

func ReserveIPBlock(jason []byte) Instance {
	path := ipblock_col_path()
	return is_post(path, jason)

}
func GetIPBlock(ipblockid string) Instance {
	path := ipblockPath(ipblockid)
	return isGet(path)
}

func ReleaseIPBlock(ipblockid string) Resp {
	path := ipblockPath(ipblockid)
	return isDelete(path)
}
