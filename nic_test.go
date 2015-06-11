// nic_test.go

package gonap

import "testing"

var nic_dcid = mkdcid()
var nic_srvid = mksrvid(nic_dcid)

func mknicid(nic_dcid, nic_srvid string) string {
	var jason = []byte(`{
					"name":"Original Nic",
					"lan":1
					}`)

	nic := CreateNic(nic_dcid, nic_srvid, jason)
	nicid := nic.Id
	return nicid
}

func TestListNics(t *testing.T) {
	//t.Parallel()
	shouldbe := "collection"
	want := 200
	resp := ListNics(nic_dcid, nic_srvid)
	if resp.Type != shouldbe {
		t.Errorf("ListNics() type == %v, wanted %v", resp.Type, shouldbe)
	}
	if resp.Resp.StatusCode != want {
		t.Errorf(bad_status(want, resp.Resp.StatusCode))
	}
}
func TestCreateNic(t *testing.T) {
	//t.Parallel()
	want := 202
	var jason = []byte(`{
					"name":"Original Nic",
					"lan":1
					}`)

	resp := CreateNic(nic_dcid, nic_srvid, jason)
	if resp.Resp.StatusCode != want {
		t.Errorf(bad_status(want, resp.Resp.StatusCode))
	}
}
func TestGetNic(t *testing.T) {
	//t.Parallel()
	shouldbe := "nic"
	want := 200
	nicid := mknicid(nic_dcid, nic_srvid)
	resp := GetNic(nic_dcid, nic_srvid, nicid)

	if resp.Type != shouldbe {
		t.Errorf("GetNic() type == %v, wanted %v", resp.Type, shouldbe)
	}
	if resp.Resp.StatusCode != want {
		t.Errorf(bad_status(want, resp.Resp.StatusCode))
	}
}
func TestPatchNic(t *testing.T) {
	//t.Parallel()
	want := 202
	jason_patch := []byte(`{
					"name":"Patched Nic",
					"lan":1
					}`)

	nicid := mknicid(nic_dcid, nic_srvid)
	resp := PatchNic(nic_dcid, nic_srvid, nicid, jason_patch)
	if resp.Resp.StatusCode != want {
		t.Errorf(bad_status(want, resp.Resp.StatusCode))
	}
}
func TestUpdateNic(t *testing.T) {
	//t.Parallel()
	want := 202
	jason_update := []byte(`{
					"name":"Update Nic",
					"lan":1
					}`)

	nicid := mknicid(nic_dcid, nic_srvid)
	resp := UpdateNic(nic_dcid, nic_srvid, nicid, jason_update)
	if resp.Resp.StatusCode != want {
		t.Errorf(bad_status(want, resp.Resp.StatusCode))
	}
}
func TestDeleteNic(t *testing.T) {
	//t.Parallel()
	want := 202
	nicid := mknicid(nic_dcid, nic_srvid)
	resp := DeleteNic(nic_dcid, nic_srvid, nicid)
	if resp.StatusCode != want {
		t.Errorf(bad_status(want, resp.StatusCode))
	}
}
