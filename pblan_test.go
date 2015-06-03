// pblan_test.go
package gonap

import "testing"

var lan_dcid = mkdcid()

func mklanid(lan_dcid string) string {
	var jason = []byte(`{
					"name":"A Lan"
					}`)
	resp := CreateLan(lan_dcid, jason)
	lan := AsLan(resp.Body)
	lanid := lan.Id
	return lanid
}

func TestListLans(t *testing.T) {
	t.Parallel()
	shouldbe := "collection"
	want := 200
	resp := ListLans(lan_dcid)
	lans := AsLans(resp.Body)
	if lans.Type != shouldbe {
		t.Errorf("ListServers() type == %v, wanted %v", lans.Type, shouldbe)
	}
	if resp.StatusCode != want {
		t.Errorf("ListServers() StatusCode == %v, wanted %v", resp.StatusCode, want)
	}
}

func TestCreateLan(t *testing.T) {
	t.Parallel()
	want := 202
	var jason = []byte(`{
					"name":"A lan"
					}`)
	resp := CreateLan(lan_dcid, jason)
	if resp.StatusCode != want {
		t.Errorf("CreateLan() StatusCode == %v, wanted %v", resp.StatusCode, want)
	}
}

func TestGetLan(t *testing.T) {
	t.Parallel()
	shouldbe := "lan"
	want := 200
	lanid := mklanid(lan_dcid)
	resp := GetLan(lan_dcid, lanid)
	lan := AsLan(resp.Body)
	if lan.Type != shouldbe {
		t.Errorf("GetLan() type == %v, wanted %v", lan.Type, shouldbe)
	}
	if resp.StatusCode != want {
		t.Errorf("GetLan() StatusCode == %v, wanted %v", resp.StatusCode, want)
	}
}

func TestPatchLan(t *testing.T) {
	t.Parallel()
	want := 202
	jason_patch := []byte(`{
					"name":"Patched lan",
					}`)
	lanid := mklanid(lan_dcid)
	resp := PatchLan(lan_dcid, lanid, jason_patch)
	if resp.StatusCode != want {
		t.Errorf("PatchLan() StatusCode == %v, wanted %v", resp.StatusCode, want)
	}
}

func TestUpdateLan(t *testing.T) {
	t.Parallel()
	want := 202
	jason_update := []byte(`{
					"name":"Updated Lan"
					}`)

	lanid := mklanid(lan_dcid)
	resp := UpdateLan(lan_dcid, lanid, jason_update)
	if resp.StatusCode != want {
		t.Errorf("UpdateLan() StatusCode == %v, wanted %v", resp.StatusCode, want)
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
