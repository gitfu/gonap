package gonap

func srvcmd(dcid, srvid, cmd string) PBResp {
	jason := `
		{}
		`
	path := server_cmd_path(dcid, srvid, cmd)
	return is_command(path, jason)
}

/**
Start server command
	Execute start server
Stop server command
	Execute Stop Server
Reboot server command
	Execute reboot server

**/
func StartServer(dcid, srvid string) PBResp {
	return srvcmd(dcid, srvid, "start")
}

func StopServer(dcid, srvid string) PBResp {
	return srvcmd(dcid, srvid, "stop")
}

func RebootServer(dcid, srvid string) PBResp {
	return srvcmd(dcid, srvid, "reboot")
}
