package gonap

import "encoding/json"

// toServer converts a Resp struct into a Instance struct
func toServer(resp Resp) Instance {
	var Server Instance
	json.Unmarshal(resp.Body, &Server)
	Server.Resp = resp
	return Server
}

// toServers converts a Resp struct into a Collection struct
func toServers(resp Resp) Collection {
	var Servers Collection
	json.Unmarshal(resp.Body, &Servers)
	Servers.Resp = resp
	return Servers
}

// ListServers returns a server struct collection
func ListServers(dcid string) Collection {
	path := server_col_path(dcid)
	return toServers(is_get(path))
}

// CreateServer creates a server from a jason []byte and returns a Instance struct
func CreateServer(dcid string, jason []byte) Instance {
	path := server_col_path(dcid)
	return toServer(is_post(path, jason))
}

// GetServer pulls data for the server where id = srvid returns a Instance struct
func GetServer(dcid, srvid string) Instance {
	path := server_path(dcid, srvid)
	return toServer(is_get(path))
}

// UpdateServer is a full update of server properties passed in as jason []byte
// Returns Instance struct
func UpdateServer(dcid string, srvid string, jason []byte) Instance {
	path := server_path(dcid, srvid)
	return toServer(is_put(path, jason))
}

// PatchServer partial update of server properties passed in as jason []byte
// Returns Instance struct
func PatchServer(dcid string, srvid string, jason []byte) Instance {
	path := server_path(dcid, srvid)
	return toServer(is_patch(path, jason))
}

// DeleteServer deletes the server where id=srvid and returns Resp struct
func DeleteServer(dcid, srvid string) Resp {
	path := server_path(dcid, srvid)
	return is_delete(path)
}

func toCdRom(resp Resp) Instance {
	var cdrom Instance
	json.Unmarshal(resp.Body, &cdrom)
	cdrom.Resp = resp
	return cdrom
}

func toCdRoms(resp Resp) Collection {
	var cdroms Collection
	json.Unmarshal(resp.Body, &cdroms)
	cdroms.Resp = resp
	return cdroms
}

func ListAttachedCdroms(dcid, srvid string) Collection {
	path := server_cdrom_col_path(dcid, srvid)
	return toCdRoms(is_get(path))
}

func AttachCdrom(dcid string, srvid string, cdid string) Instance {
	jason := []byte(`{"id":"` + cdid + `"}`)
	path := server_cdrom_col_path(dcid, srvid)
	return toCdRom(is_post(path, jason))
}

func GetAttachedCdrom(dcid, srvid, cdid string) Instance {
	path := server_cdrom_path(dcid, srvid, cdid)
	return toCdRom(is_get(path))
}

func DetachCdrom(dcid, srvid, cdid string) Resp {
	path := server_cdrom_path(dcid, srvid, cdid)
	return is_delete(path)
}

func ListAttachedVolumes(dcid, srvid string) Collection {
	path := server_volume_col_path(dcid, srvid)
	return toVolumes(is_get(path))
}

func AttachVolume(dcid string, srvid string, volid string) Instance {
	jason := []byte(`{"id":"` + volid + `"}`)
	path := server_volume_col_path(dcid, srvid)
	return toVolume(is_post(path, jason))
}

func GetAttachedVolume(dcid, srvid, volid string) Instance {
	path := server_volume_path(dcid, srvid, volid)
	return toVolume(is_get(path))
}

func DetachVolume(dcid, srvid, volid string) Resp {
	path := server_volume_path(dcid, srvid, volid)
	return is_delete(path)
}

// server_command is a generic function for running server commands
func server_command(dcid, srvid, cmd string) Resp {
	jason := `
		{}
		`
	path := server_command_path(dcid, srvid, cmd)
	return is_command(path, jason)
}

// StartServer starts a server
func StartServer(dcid, srvid string) Resp {
	return server_command(dcid, srvid, "start")
}

// StopServer stops a server
func StopServer(dcid, srvid string) Resp {
	return server_command(dcid, srvid, "stop")
}

// RebootServer reboots a server
func RebootServer(dcid, srvid string) Resp {
	return server_command(dcid, srvid, "reboot")
}
