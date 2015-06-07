package gonap

import "encoding/json"

// Server is a struct for Server data
type Server Instance

// toServer converts a PBResp struct into a Server struct
func toServer(pbresp PBResp) Server {
	var Server Server
	json.Unmarshal(pbresp.Body, &Server)
	Server.Resp = pbresp
	return Server
}

// Servers is a struct for Server struct collections
type Servers Collection

// toServers converts a PBResp struct into a Servers struct
func toServers(pbresp PBResp) Servers {
	var Servers Servers
	json.Unmarshal(pbresp.Body, &Servers)
	Servers.Resp = pbresp
	return Servers
}

// ListServers returns a server struct collection
func ListServers(dcid string) Servers {
	path := server_col_path(dcid)
	return toServers(is_get(path))
}

// CreateServer creates a server from a jason []byte and returns a Server struct
func CreateServer(dcid string, jason []byte) Server {
	path := server_col_path(dcid)
	return toServer(is_post(path, jason))
}

// GetServer pulls data for the server where id = srvid returns a Server struct
func GetServer(dcid, srvid string) Server {
	path := server_path(dcid, srvid)
	return toServer(is_get(path))
}

// UpdateServer is a full update of server properties passed in as jason []byte
// Returns Server struct
func UpdateServer(dcid string, srvid string, jason []byte) Server {
	path := server_path(dcid, srvid)
	return toServer(is_put(path, jason))
}

// PatchServer partial update of server properties passed in as jason []byte
// Returns Server struct
func PatchServer(dcid string, srvid string, jason []byte) Server {
	path := server_path(dcid, srvid)
	return toServer(is_patch(path, jason))
}

// DeleteServer deletes the server where id=srvid and returns Server struct
func DeleteServer(dcid, srvid string) Server {
	path := server_path(dcid, srvid)
	return toServer(is_delete(path))
}

func toCdRom(pbresp PBResp) Image {
	var cdrom Image
	json.Unmarshal(pbresp.Body, &cdrom)
	cdrom.Resp = pbresp
	return cdrom
}

func toCdRoms(pbresp PBResp) Images {
	var cdroms Images
	json.Unmarshal(pbresp.Body, &cdroms)
	cdroms.Resp = pbresp
	return cdroms
}

func ListAttachedCdroms(dcid, srvid string) Images {
	path := server_cdrom_col_path(dcid, srvid)
	return toCdRoms(is_get(path))
}

func AttachCdrom(dcid string, srvid string, cdid string) Image {
	jason := []byte(`{"id":"` + cdid + `"}`)
	path := server_cdrom_col_path(dcid, srvid)
	return toCdRom(is_post(path, jason))
}

func GetAttachedCdrom(dcid, srvid, cdid string) Image {
	path := server_cdrom_path(dcid, srvid, cdid)
	return toCdRom(is_get(path))
}

func DetachCdrom(dcid, srvid, cdid string) Image {
	path := server_cdrom_path(dcid, srvid, cdid)
	return toCdRom(is_delete(path))
}

func ListAttachedVolumes(dcid, srvid string) Volumes {
	path := server_volume_col_path(dcid, srvid)
	return toVolumes(is_get(path))
}

func AttachVolume(dcid string, srvid string, volid string) Volume {
	jason := []byte(`{"id":"` + volid + `"}`)
	path := server_volume_col_path(dcid, srvid)
	return toVolume(is_post(path, jason))
}

func GetAttachedVolume(dcid, srvid, volid string) Volume {
	path := server_volume_path(dcid, srvid, volid)
	return toVolume(is_get(path))
}

func DetachVolume(dcid, srvid, volid string) Volume {
	path := server_volume_path(dcid, srvid, volid)
	return toVolume(is_delete(path))
}

// server_command is a generic function for running server commands
func server_command(dcid, srvid, cmd string) Server {
	jason := `
		{}
		`
	path := server_command_path(dcid, srvid, cmd)
	return toServer(is_command(path, jason))
}

// StartServer starts a server
func StartServer(dcid, srvid string) Server {
	return server_command(dcid, srvid, "start")
}

// StopServer stops a server
func StopServer(dcid, srvid string) Server {
	return server_command(dcid, srvid, "stop")
}

// RebootServer reboots a server
func RebootServer(dcid, srvid string) Server {
	return server_command(dcid, srvid, "reboot")
}
