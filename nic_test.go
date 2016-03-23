// nic_test.go

package goprofitbricks

import "testing"

var nicDcid = mkdcid()
var nicSrvid = mksrvid(nicDcid)

func mknicid(nicDcid, nicSrvid string) string {
	var jason = []byte(`{
					"name":"Original Nic",
					"lan":1
					}`)

	nic := CreateNic(nicDcid, nicSrvid, jason)
	nicid := nic.Id
	return nicid
}

func TestListNics(t *testing.T) {
	//t.Parallel()
	shouldbe := "collection"
	want := 200
	resp := ListNics(nicDcid, nicSrvid)
	if resp.Type != shouldbe {
		t.Errorf(badType(shouldbe, resp.Type))
	}
	if resp.Resp.StatusCode != want {
		t.Errorf(badStatus(want, resp.Resp.StatusCode))
	}
}
func TestCreateNic(t *testing.T) {
	//t.Parallel()
	want := 202
	var jason = []byte(`{
					"name":"Original Nic",
					"lan":1
					}`)

	resp := CreateNic(nicDcid, nicSrvid, jason)
	if resp.Resp.StatusCode != want {
		t.Errorf(badStatus(want, resp.Resp.StatusCode))
	}
}
func TestGetNic(t *testing.T) {
	//t.Parallel()
	shouldbe := "nic"
	want := 200
	nicid := mknicid(nicDcid, nicSrvid)
	resp := GetNic(nicDcid, nicSrvid, nicid)
	if resp.Type != shouldbe {
		t.Errorf(badType(shouldbe, resp.Type))
	}
	if resp.Resp.StatusCode != want {
		t.Errorf(badStatus(want, resp.Resp.StatusCode))
	}
}
func TestPatchNic(t *testing.T) {
	//t.Parallel()
	want := 202
	jasonPatch := []byte(`{
					"name":"Patched Nic",
					"lan":1
					}`)

	nicid := mknicid(nicDcid, nicSrvid)
	resp := PatchNic(nicDcid, nicSrvid, nicid, jasonPatch)
	if resp.Resp.StatusCode != want {
		t.Errorf(badStatus(want, resp.Resp.StatusCode))
	}
}
func TestUpdateNic(t *testing.T) {
	//t.Parallel()
	want := 202
	jasonUpdate := []byte(`{
					"name":"Update Nic",
					"lan":1
					}`)

	nicid := mknicid(nicDcid, nicSrvid)
	resp := UpdateNic(nicDcid, nicSrvid, nicid, jasonUpdate)
	if resp.Resp.StatusCode != want {
		t.Errorf(badStatus(want, resp.Resp.StatusCode))
	}
}
func TestDeleteNic(t *testing.T) {
	//t.Parallel()
	want := 202
	nicid := mknicid(nicDcid, nicSrvid)
	resp := DeleteNic(nicDcid, nicSrvid, nicid)
	if resp.StatusCode != want {
		t.Errorf(badStatus(want, resp.StatusCode))
	}
}
