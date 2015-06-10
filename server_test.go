// server_test.go
package gonap

import "testing"
import "fmt"

var srv_dcid = mkdcid()

func mksrvid(srv_dcid string) string {
	var jason = []byte(`{
					"name":"Original Server",
					"cores":4,
					"ram": 4096
					}`)
	srv := CreateServer(srv_dcid, jason)

	srvid := srv.Id
	return srvid
}

func ExampleListServers() {
	s := ListServers(srv_dcid)
	fmt.Println(s.Resp.StatusCode)
	// Output: 200

}

func TestListServers(t *testing.T) {
	t.Parallel()
	shouldbe := "collection"
	want := 200
	srvs := ListServers(srv_dcid)
	if srvs.Type != shouldbe {
		t.Errorf("ListServers() type == %v, wanted %v", srvs.Type, shouldbe)
	}
	if srvs.Resp.StatusCode != want {
		t.Errorf("ListServers() StatusCode == %v, wanted %v", srvs.Resp.StatusCode, want)
	}
}

func TestCreateServer(t *testing.T) {
	t.Parallel()
	want := 202
	var jason = []byte(`{
					"name":"Goat",
					"cores":4,
					"ram": 4096
					}`)
	newsrv := CreateServer(srv_dcid, jason)
	if newsrv.Resp.StatusCode != want {
		t.Errorf("CreateServer() StatusCode == %v, wanted %v", newsrv.Resp.StatusCode, want)
	}
}

func TestGetServer(t *testing.T) {
	t.Parallel()
	shouldbe := "server"
	want := 200
	srvid := mksrvid(srv_dcid)
	srv := GetServer(srv_dcid, srvid)

	if srv.Type != shouldbe {
		t.Errorf("GetServer() type == %v, wanted %v", srv.Type, shouldbe)
	}
	if srv.Resp.StatusCode != want {
		t.Errorf("GetServer() StatusCode == %v, wanted %v", srv.Resp.StatusCode, want)
	}
}

func TestPatchServer(t *testing.T) {
	t.Parallel()
	want := 202
	jason_patch := []byte(`{
					"name":"Renamed Server",
					}`)
	srvid := mksrvid(srv_dcid)
	srv := PatchServer(srv_dcid, srvid, jason_patch)
	if srv.Resp.StatusCode != want {
		t.Errorf("PatchServer() StatusCode == %v, wanted %v", srv.Resp.StatusCode, want)
	}
}

func TestUpdateServer(t *testing.T) {
	t.Parallel()
	want := 202
	jason_update := []byte(`{
					"name":"Renamed Server",
					"cores":16,
					"ram": 8192
					}`)

	srvid := mksrvid(srv_dcid)
	srv := UpdateServer(srv_dcid, srvid, jason_update)
	if srv.Resp.StatusCode != want {
		t.Errorf("UpdateServer() StatusCode == %v, wanted %v", srv.Resp.StatusCode, want)
	}
}

func TestDeleteServer(t *testing.T) {
	t.Parallel()
	want := 202
	srvid := mksrvid(srv_dcid)
	srv := DeleteServer(srv_dcid, srvid)
	if srv.StatusCode != want {
		t.Errorf("DeleteServer() StatusCode == %v, wanted %v", srv.StatusCode, want)
	}
}
