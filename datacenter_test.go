package goprofitbricks

import "testing"

func mkdcid() string {
	dc := CreateDatacenter([]byte(`{
					"name":"Original DC",
					"location":"us/las"
					}`))

	dcid := dc.Id
	return dcid
}

func TestListDatacenters(t *testing.T) {
	////t.Parallel()
	shouldbe := "collection"
	want := 200
	resp := ListDatacenters()

	if resp.Type != shouldbe {
		t.Errorf(badType(shouldbe, resp.Type))
	}
	if resp.Resp.StatusCode != want {
		t.Errorf(badStatus(want, resp.Resp.StatusCode))
	}
}
func TestCreateDatacenter(t *testing.T) {
	////t.Parallel()
	want := 202
	var jason = []byte(`{
					"name":"Original DC",
					"location":"us/las"
					}`)
	resp := CreateDatacenter(jason)
	if resp.Resp.StatusCode != want {
		t.Errorf(badStatus(want, resp.Resp.StatusCode))
	}
}
func TestGetDatacenter(t *testing.T) {
	////t.Parallel()
	shouldbe := "datacenter"
	want := 200
	dcid := mkdcid()
	resp := GetDatacenter(dcid)
	if resp.Type != shouldbe {
		t.Errorf(badType(shouldbe, resp.Type))
	}
	if resp.Resp.StatusCode != want {
		t.Errorf(badStatus(want, resp.Resp.StatusCode))
	}
}

func TestPatchDatacenter(t *testing.T) {
	////t.Parallel()
	want := 202
	jason_patch := []byte(`{
					"name":"Renamed DC",
					}`)
	dcid := mkdcid()
	resp := PatchDatacenter(dcid, jason_patch)
	if resp.Resp.StatusCode != want {
		t.Errorf(badStatus(want, resp.Resp.StatusCode))
	}

}

func TestUpdateDatacenter(t *testing.T) {
	////t.Parallel()
	want := 202
	jason_update := []byte(`{
					"name":"Renamed DC",
					"location":"us/las"
					}`)

	dcid := mkdcid()
	resp := UpdateDatacenter(dcid, jason_update)
	if resp.Resp.StatusCode != want {
		t.Errorf(badStatus(want, resp.Resp.StatusCode))
	}
}

func TestDeleteDatacenter(t *testing.T) {
	////t.Parallel()
	want := 202
	dcid := mkdcid()
	resp := DeleteDatacenter(dcid)
	if resp.StatusCode != want {
		t.Errorf(badStatus(want, resp.StatusCode))
	}
}
