// lan_test.go
package gonap

import "testing"
import "fmt"

var lan_dcid = mkdcid()

func mklanid(lan_dcid string) string {
	var jason = []byte(`{
					"name":"A Lan"
					}`)
	lan := CreateLan(lan_dcid, jason)
	lanid := lan.Id
	return lanid
}

func TestListLans(t *testing.T) {
	t.Parallel()
	shouldbe := "collection"
	want := 200
	lans := ListLans(lan_dcid)

	if lans.Type != shouldbe {
		t.Errorf("ListServers() type == %v, wanted %v", lans.Type, shouldbe)
	}
	if lans.Resp.StatusCode != want {
		fmt.Println(lans.Resp.Body)
		t.Errorf("ListServers() StatusCode == %v, wanted %v", lans.Resp.StatusCode, want)
	}
}

func TestCreateLan(t *testing.T) {
	t.Parallel()
	want := 202
	var jason = []byte(`{
					"name":"A lan"
					}`)
	lan := CreateLan(lan_dcid, jason)
	if lan.Resp.StatusCode != want {
		t.Errorf("CreateLan() StatusCode == %v, wanted %v", lan.Resp.StatusCode, want)
	}
}

func TestGetLan(t *testing.T) {
	t.Parallel()
	shouldbe := "lan"
	want := 200
	lanid := mklanid(lan_dcid)
	lan := GetLan(lan_dcid, lanid)
	if lan.Type != shouldbe {
		t.Errorf("GetLan() type == %v, wanted %v", lan.Type, shouldbe)
	}
	if lan.Resp.StatusCode != want {
		t.Errorf("GetLan() StatusCode == %v, wanted %v", lan.Resp.StatusCode, want)
	}
}

func TestPatchLan(t *testing.T) {
	t.Parallel()
	want := 202
	jason_patch := []byte(`{
					"name":"Patched lan",
					}`)
	lanid := mklanid(lan_dcid)
	lan := PatchLan(lan_dcid, lanid, jason_patch)
	if lan.Resp.StatusCode != want {
		t.Errorf("PatchLan() StatusCode == %v, wanted %v", lan.Resp.StatusCode, want)
	}
}

func TestUpdateLan(t *testing.T) {
	t.Parallel()
	want := 202
	jason_update := []byte(`{
					"name":"Updated Lan"
					}`)

	lanid := mklanid(lan_dcid)
	lan := UpdateLan(lan_dcid, lanid, jason_update)
	if lan.Resp.StatusCode != want {
		t.Errorf("UpdateLan() StatusCode == %v, wanted %v", lan.Resp.StatusCode, want)
	}
}

func TestDeleteLan(t *testing.T) {
	t.Parallel()
	want := 202
	lanid := mklanid(lan_dcid)
	resp := DeleteLan(lan_dcid, lanid)
	if resp.StatusCode != want {
		t.Errorf("DeleteLan() StatusCode == %v, wanted %v", resp.StatusCode, want)
	}
}
