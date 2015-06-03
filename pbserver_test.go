// pbserver_test.go
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
	resp := CreateServer(srv_dcid, jason)
	srv := AsServer(resp.Body)
	srvid := srv.Id
	return srvid
}

func ExampleListServers() {
	//	shouldbe := "collection"
	//want := 200
	resp := ListServers(srv_dcid)
	srvs := AsServers(resp.Body)
	fmt.Println(srvs.Type)
	// Output: collection

}

func TestListServers(t *testing.T) {
	t.Parallel()
	shouldbe := "collection"
	want := 200
	resp := ListServers(srv_dcid)
	srvs := AsServers(resp.Body)
	if srvs.Type != shouldbe {
		t.Errorf("ListServers() type == %v, wanted %v", srvs.Type, shouldbe)
	}
	if resp.StatusCode != want {
		t.Errorf("ListServers() StatusCode == %v, wanted %v", resp.StatusCode, want)
	}
}

func TestCreateServer(t *testing.T) {
	t.Parallel()
	want := 202
	var jason = []byte(`{
					"name":"Original Server",
					"cores":4,
					"ram": 4096
					}`)
	resp := CreateServer(srv_dcid, jason)
	if resp.StatusCode != want {
		t.Errorf("CreateServer() StatusCode == %v, wanted %v", resp.StatusCode, want)
	}
}

func TestGetServer(t *testing.T) {
	t.Parallel()
	shouldbe := "server"
	want := 200
	srvid := mksrvid(srv_dcid)
	resp := GetServer(srv_dcid, srvid)
	srv := AsServer(resp.Body)
	if srv.Type != shouldbe {
		t.Errorf("GetServer() type == %v, wanted %v", srv.Type, shouldbe)
	}
	if resp.StatusCode != want {
		t.Errorf("GetServer() StatusCode == %v, wanted %v", resp.StatusCode, want)
	}
}

func TestPatchServer(t *testing.T) {
	t.Parallel()
	want := 202
	jason_patch := []byte(`{
					"name":"Renamed Server",
					}`)
	srvid := mksrvid(srv_dcid)
	resp := PatchServer(srv_dcid, srvid, jason_patch)
	if resp.StatusCode != want {
		t.Errorf("PatchServer() StatusCode == %v, wanted %v", resp.StatusCode, want)
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
	resp := UpdateServer(srv_dcid, srvid, jason_update)
	if resp.StatusCode != want {
		t.Errorf("UpdateServer() StatusCode == %v, wanted %v", resp.StatusCode, want)
	}
}

func TestDeleteServer(t *testing.T) {
	t.Parallel()
	want := 202
	srvid := mksrvid(srv_dcid)
	resp := DeleteServer(srv_dcid, srvid)
	if resp.StatusCode != want {
		t.Errorf("DeleteServer() StatusCode == %v, wanted %v", resp.StatusCode, want)
	}
}
