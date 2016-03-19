// snapshot_test.go

package goprofitbricks

import "testing"

func mksnapid() string {
	resp := ListSnapshots()

	snapid := resp.Items[0].Id
	//fmt.Println(MkJson(resp))
	return snapid
}

func TestListSnapshots(t *testing.T) {
	////t.Parallel()
	shouldbe := "collection"
	want := 200
	resp := ListSnapshots()

	if resp.Type != shouldbe {
		t.Errorf(badType(shouldbe, resp.Type))
	}
	if resp.Resp.StatusCode != want {
		t.Errorf(badStatus(want, resp.Resp.StatusCode))
	}
}

func TestGetSnapshot(t *testing.T) {
	////t.Parallel()
	shouldbe := "snapshot"
	want := 200
	snapid := mksnapid()
	resp := GetSnapshot(snapid)
	if resp.Type != shouldbe {
		t.Errorf(badType(shouldbe, resp.Type))
	}
	if resp.Resp.StatusCode != want {
		t.Errorf(badStatus(want, resp.Resp.StatusCode))
	}
}

func TestPatchSnapshot(t *testing.T) {
	////t.Parallel()
	want := 202
	jason_patch := []byte(`{
					"name":"Renamed snap",
					}`)
	snapid := mksnapid()
	resp := PatchSnapshot(snapid, jason_patch)
	if resp.Resp.StatusCode != want {
		t.Errorf(badStatus(want, resp.Resp.StatusCode))
	}

}

func TestUpdateSnapshot(t *testing.T) {
	////t.Parallel()
	want := 202
	jason_update := []byte(`{
					"name":"Renamed snap",
					"location":"us/las"
					}`)

	snapid := mksnapid()
	resp := UpdateSnapshot(snapid, jason_update)
	if resp.Resp.StatusCode != want {
		t.Errorf(badStatus(want, resp.Resp.StatusCode))
	}
}

func TestDeleteSnapshot(t *testing.T) {
	////t.Parallel()
	want := 202
	snapid := mksnapid()
	resp := DeleteSnapshot(snapid)
	if resp.StatusCode != want {
		t.Errorf(badStatus(want, resp.StatusCode))
	}
}
