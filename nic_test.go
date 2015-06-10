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
	t.Parallel()
	shouldbe := "collection"
	want := 200
	nics := ListNics(nic_dcid, nic_srvid)
	if nics.Type != shouldbe {
		t.Errorf("ListNics() type == %v, wanted %v", nics.Type, shouldbe)
	}
	if nics.Resp.StatusCode != want {
		t.Errorf("ListNics() StatusCode == %v, wanted %v", nics.Resp.StatusCode, want)
	}
}

func TestCreateNic(t *testing.T) {
	t.Parallel()
	want := 202
	var jason = []byte(`{
					"name":"Original Nic",
					"lan":1
					}`)

	nic := CreateNic(nic_dcid, nic_srvid, jason)
	if nic.Resp.StatusCode != want {
		t.Errorf("CreateNic() StatusCode == %v, wanted %v", nic.Resp.StatusCode, want)
	}
}

func TestGetNic(t *testing.T) {
	t.Parallel()
	shouldbe := "nic"
	want := 200
	nicid := mknicid(nic_dcid, nic_srvid)
	nic := GetNic(nic_dcid, nic_srvid, nicid)

	if nic.Type != shouldbe {
		t.Errorf("GetNic() type == %v, wanted %v", nic.Type, shouldbe)
	}
	if nic.Resp.StatusCode != want {
		t.Errorf("GetNic() StatusCode == %v, wanted %v", nic.Resp.StatusCode, want)
	}
}

func TestPatchNic(t *testing.T) {
	t.Parallel()
	want := 202
	jason_patch := []byte(`{
					"name":"Patched Nic",
					"lan":1
					}`)

	nicid := mknicid(nic_dcid, nic_srvid)
	nic := PatchNic(nic_dcid, nic_srvid, nicid, jason_patch)
	if nic.Resp.StatusCode != want {
		t.Errorf("PatchNic() StatusCode == %v, wanted %v", nic.Resp.StatusCode, want)
	}
}

func TestUpdateNic(t *testing.T) {
	t.Parallel()
	want := 202
	jason_update := []byte(`{
					"name":"Update Nic",
					"lan":1
					}`)

	nicid := mknicid(nic_dcid, nic_srvid)
	nic := UpdateNic(nic_dcid, nic_srvid, nicid, jason_update)
	if nic.Resp.StatusCode != want {
		t.Errorf("UpdateNic() StatusCode == %v, wanted %v", nic.Resp.StatusCode, want)
	}
}

func TestDeleteNic(t *testing.T) {
	t.Parallel()
	want := 202
	nicid := mknicid(nic_dcid, nic_srvid)
	nic := DeleteNic(nic_dcid, nic_srvid, nicid)
	if nic.StatusCode != want {
		t.Errorf("DeleteNic() StatusCode == %v, wanted %v", nic.StatusCode, want)
	}
}
